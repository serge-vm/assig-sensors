package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
)

var config SensoremuConfig

type SensorData struct {
	SensorId     string  `json:"sensorId"`
	TemperatureC float32 `json:"temperatureC"`
}

type SensoremuConfig struct {
	Host     string
	Port     string
	SensorId string
}

func parseParams() {
	flag.StringVar(&config.Host, "h", "localhost", "Host to bind to")
	flag.StringVar(&config.Port, "p", "8081", "Port to listen to")
	flag.StringVar(&config.SensorId, "i", "5d316ee8-a785-4e87-91d8-06f901c98a88", "Sensor ID to return")
	flag.Parse()
}

func sensorHandler(w http.ResponseWriter, r *http.Request) {
	// Generate random temperature between -30 and +50 and round value to precision 1
	temperature := float32(math.Round((-30+rand.Float64()*80)*10) / 10)

	pData := SensorData{config.SensorId, temperature}
	data, err := json.Marshal(pData)
	if err != nil {
		log.Fatal("Wrong json data")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	parseParams()
	http.HandleFunc("/", sensorHandler)

	fmt.Println("Sensor emulator is running at http://" + config.Host + ":" + config.Port)
	log.Fatal(http.ListenAndServe(config.Host+":"+config.Port, nil))
}
