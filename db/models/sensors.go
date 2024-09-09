package models

import (
	"database/sql"
)

type Sensor struct {
	ID           string
	Name         string
	Address      string
	Port         int
	TemperatureC float32
}

func ListSensors(db *sql.DB) ([]Sensor, error) {
	var sensors []Sensor
	rows, err := db.Query("SELECT * FROM sensors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var sensor Sensor
		if err := rows.Scan(&sensor.ID, &sensor.Name, &sensor.Address, &sensor.Port); err != nil {
			return nil, err
		}
		sensors = append(sensors, sensor)
	}
	return sensors, nil
}
