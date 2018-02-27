bme280-exporter: main.go
	GOOS=linux GOARCH=arm GOARM=7 go build -o bme280-exporter

bme280-exporter-arm32v6: main.go
	GOOS=linux GOARCH=arm GOARM=6 go build -o bme280-exporter-arm32v6
