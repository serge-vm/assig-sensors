package aggregator

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"example.com/sensors/db/models"
)

type sensorData struct {
	SenorID      string
	TemperatureC float32
}

func GetSensors(sensors []models.Sensor) []models.Sensor {
	var wgs, wga sync.WaitGroup
	valuesC := make(chan models.Sensor)
	var si []models.Sensor
	httpClient := http.Client{Timeout: 2 * time.Second}

	wga.Add(1)
	go func() {
		for info := range valuesC {
			si = append(si, info)
		}
		wga.Done()
	}()

	wgs.Add(len(sensors))
	for _, s := range sensors {
		go scanSensor(&httpClient, s, &wgs, valuesC)
	}
	wgs.Wait()
	close(valuesC)

	wga.Wait()
	return si
}

func scanSensor(c *http.Client, s models.Sensor, wgs *sync.WaitGroup, result chan<- models.Sensor) {
	var sd sensorData

	defer wgs.Done()

	resp, err := c.Get(fmt.Sprintf("http://%v:%d", s.Address, s.Port))
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	err = json.Unmarshal(body, &sd)
	if err != nil {
		log.Println(err)
		return
	}
	s.TemperatureC = sd.TemperatureC
	result <- s
}
