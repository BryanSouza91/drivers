package main

import (
	"fmt"
	"machine"
	"time"
	"tinygo.org/x/drivers/qmc5883l"
)

func main() {
	machine.I2C0.Configure(machine.I2CConfig{})
	sensor := qmc5883l.New(machine.I2C0)

	cfg := qmc5883l.Configuration{
		Mode: qmc5883l.ModeContinuous,
		ODR:  qmc5883l.ODR_200Hz,
		RNG:  qmc5883l.RNG_8G,
		OSR:  qmc5883l.OSR_512,
	}
	sensor.Configure(cfg)

	for {
		x, y, z, _ := sensor.ReadMagneticField()
		fmt.Printf("X: %d, Y: %d, Z: %d\n", x, y, z)
		time.Sleep(time.Millisecond * 100)
	}
}
