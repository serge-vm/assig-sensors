all: clean sensoremu sensors

.PHONY: sensoremu
sensoremu:
	go build -o bin/sensoremu cmd/sensoremu/main.go

.PHONY: sensors
sensors:
	go build -o bin/sensors cmd/sensors/main.go

clean:
	rm -f bin/*