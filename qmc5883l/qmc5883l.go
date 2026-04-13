package qmc5883l

import (
	"machine"
)

type Device struct {
	bus         machine.I2C
	Address     uint16
	sensitivity int32
}

// Configuration for QMC5883L device.
type Configuration struct {
	Mode uint8 // Standby or Continuous
	ODR  uint8 // Output Data Rate (10Hz, 50Hz, 100Hz, 200Hz)
	RNG  uint8 // Range (2G or 8G)
	OSR  uint8 // Oversampling Ratio (512, 256, 128, 64)
}

func New(bus machine.I2C) Device {
	return Device{
		bus:     bus,
		Address: Address,
	}
}

// Configure sets up the sensor with default or custom settings.
func (d *Device) Configure(cfg Configuration) error {
	// Standard recommendation from datasheet:
	// 0x0B = 0x01 (Set/Reset Period)
	err := d.writeByte(REG_SET_RESET, 0x01)
	if err != nil {
		return err
	}

	// Store the sensitivity to use in ReadMagneticField()
	switch cfg.RNG {
	case RNG_2G:
		d.sensitivity = 12000
	case RNG_8G:
		d.sensitivity = 3000
	default:
		d.sensitivity = 3000 // Default to 8G
	}

	// Combine Mode, Output Data Rate, Range, and Over Sampling Ratio
	configVal := cfg.Mode | cfg.ODR | cfg.RNG | cfg.OSR
	return d.writeByte(REG_CONTROL_1, configVal)
}

// ReadRaw returns the raw X, Y, Z magnetic field components from the sensor.
// These are the direct 16-bit signed values from the registers.
func (d *Device) ReadRaw() (x, y, z int16, err error) {
	data := make([]byte, 6)
	err = d.bus.ReadRegister(uint8(d.Address), REG_DATA_X_LSB, data)
	if err != nil {
		return 0, 0, 0, err
	}

	// QMC5883L is Little-Endian
	x = int16(uint16(data[0]) | (uint16(data[1]) << 8))
	y = int16(uint16(data[2]) | (uint16(data[3]) << 8))
	z = int16(uint16(data[4]) | (uint16(data[5]) << 8))
	return
}

// ReadMagneticField returns the magnetic field components in microteslas (uT).
// This satisfies the TinyGo Magnetometer interface.
func (d *Device) ReadMagneticField() (x, y, z int32, err error) {
	rawX, rawY, rawZ, err := d.ReadRaw()
	if err != nil {
		return 0, 0, 0, err
	}

	// Conversion: (Raw_LSB * 100) / (LSB_per_Gauss)
	x = (int32(rawX) * 100) / d.sensitivity
	y = (int32(rawY) * 100) / d.sensitivity
	z = (int32(rawZ) * 100) / d.sensitivity

	return x, y, z, nil
}

// Helper to write a single byte to a register
func (d *Device) writeByte(reg, val uint8) error {
	return d.bus.WriteRegister(uint8(d.Address), reg, []byte{val})
}
