package main

import (
	"fmt"
	"log"

	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	d := i2c.NewBME280Driver(r, i2c.WithBus(1), i2c.WithAddress(0x76))

	if err := d.Start(); err != nil {
		log.Fatalf("starting driver: %s", err)
	}

	temp, err := d.Temperature()
	if err != nil {
		log.Fatalf("getting temp: %s", err)
	}
	fmt.Printf("Temperature: %.2f 'C\n", temp)

	press, err := d.Pressure()
	if err != nil {
		log.Fatalf("getting press: %s", err)
	}
	fmt.Printf("Pressure: %.2f hPa\n", press/100)

	humidity, err := d.Humidity()
	if err != nil {
		log.Fatalf("getting humidity: %s", err)
	}
	fmt.Printf("Humidity: %.2f %%\n", humidity)
}
