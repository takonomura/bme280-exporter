package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

var (
	tempDesc     = prometheus.NewDesc("bme280_temperature_celsius", "Temperature in celsius degree", nil, nil)
	pressDesc    = prometheus.NewDesc("bme280_pressure_hpa", "Barometric pressure in hPa", nil, nil)
	humidityDesc = prometheus.NewDesc("bme280_humidity", "Humidity in percentage of relative humidity", nil, nil)
)

type collector struct {
	*i2c.BME280Driver
}

func (c collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- tempDesc
	ch <- pressDesc
	ch <- humidityDesc
}

func (c collector) Collect(ch chan<- prometheus.Metric) {
	temp, err := c.Temperature()
	if err != nil {
		log.Printf("getting temp: %s", err)
	}
	ch <- prometheus.MustNewConstMetric(tempDesc, prometheus.GaugeValue, float64(temp))

	press, err := c.Pressure()
	if err != nil {
		log.Printf("getting press: %s", err)
	}
	ch <- prometheus.MustNewConstMetric(pressDesc, prometheus.GaugeValue, float64(press)/100)

	humidity, err := c.Humidity()
	if err != nil {
		log.Printf("getting humidity: %s", err)
	}
	ch <- prometheus.MustNewConstMetric(humidityDesc, prometheus.GaugeValue, float64(humidity))
}

func main() {
	adapter := raspi.NewAdaptor()
	d := i2c.NewBME280Driver(adapter, i2c.WithBus(1), i2c.WithAddress(0x76))

	if err := d.Start(); err != nil {
		log.Fatalf("starting driver: %s", err)
	}
	log.Print("Connected to BME280")

	r := prometheus.NewRegistry()
	c := &collector{d}
	r.MustRegister(c)
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	s := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	log.Print("Starting server...")
	log.Fatal(s.ListenAndServe())
}
