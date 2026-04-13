package qmc5883l

// I2C address
const Address = 0x0D

// Registers
const (
	REG_DATA_X_LSB = 0x00
	REG_DATA_X_MSB = 0x01
	REG_DATA_Y_LSB = 0x02
	REG_DATA_Y_MSB = 0x03
	REG_DATA_Z_LSB = 0x04
	REG_DATA_Z_MSB = 0x05
	REG_STATUS     = 0x06
	REG_TEMP_LSB   = 0x07
	REG_TEMP_MSB   = 0x08
	REG_CONTROL_1  = 0x09
	REG_CONTROL_2  = 0x0A
	REG_SET_RESET  = 0x0B
	REG_CHIP_ID    = 0x0D
)

// Mode and Config constants
const (
	ModeStandby    = 0x00
	ModeContinuous = 0x01

	// ODR (Output Data Rate) options
	ODR_10Hz  = 0x00 << 2
	ODR_50Hz  = 0x01 << 2
	ODR_100Hz = 0x02 << 2
	ODR_200Hz = 0x03 << 2

	// RNG (Range) options
	RNG_2G = 0x00 << 4
	RNG_8G = 0x01 << 4

	// OSR (Oversampling Ratio) options
	OSR_512 = 0x00 << 6
	OSR_256 = 0x01 << 6
	OSR_128 = 0x02 << 6
	OSR_64  = 0x03 << 6
)
