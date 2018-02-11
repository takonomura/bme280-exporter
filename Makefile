bme280-exporter: main.go
	GOOS=linux GOARCH=arm GOARM=7 go build -o bme280-exporter
