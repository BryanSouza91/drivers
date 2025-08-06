package is31fl3731

import (
	"fmt"

	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/internal/legacy"
)

// DeviceAdafruitCharliePlex16x9 implements TinyGo driver for Lumissil
// IS31FL3731 matrix LED driver on Adafruit 16x9 CharliePlex PWM LED Matrix
// Driver board: https://www.adafruit.com/product/2946
type DeviceAdafruitCharliePlex16x9 struct {
	Device
}

// enableLEDs enables only LEDs that are soldered on the Adafruit CharlieWing
// board. The board has following LEDs matrix layout:
//
//	"o" - connected (soldered) LEDs
//	"x" - not connected LEDs
//
//	  + - - - - - - - - - - - - - - - +
//	  | + - - - - - - - - - - - - - + |
//	  | |                           | |
//	  | |                           v v
//	+-----------------------------------+
//	| o o o o o o o o o o o o o o o o x |
//	| o o o o o o o o o o o o o o o o x |
//	| o o o o o o o o o o o o o o o o x |
//	| o o o o o o o o o o o o o o o o x |
//	| o o o o o o o o o o o o o o o o x |
//	| o o o o o o o o o o o o o o o o x |
//	| o o o o o o o o o o o o o o o o x |
//	| o o o o o o o o o o o o o o o o x |
//	| o o o o o o o o o o o o o o o o x |
//	| x x x x x x x x x x x x x x x x x |
//	+-----------------------------------+
//	  ^ ^                           | |
//	  | |                   ... - - + |
//	  | + - - - - - - - - - - - - - - +
//	  |
//	  start (address 0x00)
func (d *DeviceAdafruitCharliePlex16x9) enableLEDs() (err error) {
	for frame := FRAME_0; frame <= FRAME_7; frame++ {
		err = d.selectCommand(frame)
		if err != nil {
			return err
		}

		// Enable left half
		for i := uint8(0); i < 17; i += 2 {
			err = legacy.WriteRegister(d.bus, d.Address, i, []byte{0b11111110})
			if err != nil {
				return err
			}
		}
		// Enable right half
		for i := uint8(3); i < 17; i += 2 {
			err = legacy.WriteRegister(d.bus, d.Address, i, []byte{0b01111111})
			if err != nil {
				return err
			}
		}
		// Disable invisible column on the right side
		err = legacy.WriteRegister(d.bus, d.Address, 1, []byte{0b00000000})
		if err != nil {
			return err
		}
	}

	return nil
}

// DrawPixelXY draws a single pixel on the selected frame by its XY coordinates
// with provided PWM value [0-255]
func (d *DeviceAdafruitCharliePlex16x9) DrawPixelXY(frame, x, y, value uint8) (err error) {
	var index uint8

	if x >= 17 {
		return fmt.Errorf("invalid value: X is out of range [0, 16]")
	} else if y >= 9 {
		return fmt.Errorf("invalid value: Y is out of range [0, 9]")
	}

	// Board is one pixel shorter (9 vs 10 supported pixels)
	if x < 10 {
		index = 17*x + y + 1
	} else {
		index = 17*(17-x) - y - 1 - 1
	}

	return d.setPixelPWD(frame, index, value)
}

// NewAdafruitCharliePlex16x9 creates a new driver with Adafruit 16x9
// CharliePlex PWM LED Matrix layout.
// Available addresses:
// - 0x74 (default)
// - 0x75 (when the address jumper soldered to SCL)
// - 0x76 (when the address jumper soldered to SDA)
// - 0x77 (when the address jumper soldered to VCC)
func NewAdafruitCharliePlex16x9(bus drivers.I2C, address uint8) DeviceAdafruitCharliePlex16x9 {
	return DeviceAdafruitCharliePlex16x9{
		Device: Device{
			Address: address,
			bus:     bus,
		},
	}
}
