package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/sensors/api"
)

func main() {
	var sh api.ServerHandlers
	config := api.NewServerConfig()
	fmt.Println("Sensor aggregator is running at http://" + config.Host + ":" + config.Port)
	h := api.Handler(sh)
	log.Fatal(http.ListenAndServe(config.Host+":"+config.Port, h))
}
