package api

import "flag"

type ServerConfig struct {
	Host string
	Port string
}

func NewServerConfig() ServerConfig {
	var config ServerConfig
	flag.StringVar(&config.Host, "h", "localhost", "Host to bind to")
	flag.StringVar(&config.Port, "p", "8080", "Port to listen to")
	flag.Parse()
	return config
}
