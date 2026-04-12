package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/ist83xx"
)

func main() {

	// IST83xx is connected to the I2C0 bus on xiao-ble 
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: 400 * machine.KHz,
	})

	sensor := ist83xx.New(machine.I2C0)
	err := sensor.Configure(ist83xx.Configuration{}) //default settings
	if err != nil {
		for {
			println("Failed to configure", err.Error())
			time.Sleep(time.Second)
		}
	}

	// ReadMagnetometer returns the magnetometer readings
	m, err := sensor.ReadMagnetometer()
	if err != nil {
		println(err)
	}

	println(m)
}