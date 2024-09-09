package api

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/sensors/db"
	"example.com/sensors/db/models"
	"example.com/sensors/docs"
	"example.com/sensors/internal/aggregator"
)

type ServerHandlers struct{}

func (s ServerHandlers) SensorData(w http.ResponseWriter, r *http.Request) {
	sd := make(SensorData)
	dbConn := db.GetDB()
	sensors, err := models.ListSensors(dbConn)
	if err != nil {
		log.Println(err)
	}
	si := aggregator.GetSensors(sensors)
	for _, info := range si {
		sd[info.Name] = info.TemperatureC
	}
	data, err := json.Marshal(sd)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (s ServerHandlers) Documentation(w http.ResponseWriter, r *http.Request) {
	ui, _ := docs.UI.ReadFile("ui/index.html")
	w.Header().Set("Content-Type", "text/html")
	w.Write(ui)
}

func (s ServerHandlers) Specs(w http.ResponseWriter, r *http.Request) {
	specs, _ := docs.Specs.ReadFile("sensors-api.yaml")
	w.Header().Set("Content-Type", "text/yaml")
	w.Write(specs)
}
