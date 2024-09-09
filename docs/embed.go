package docs

import "embed"

//go:embed ui/*
var UI embed.FS

//go:embed sensors-api.yaml
var Specs embed.FS
