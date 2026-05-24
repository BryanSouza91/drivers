package ist83xx

import (
	"errors"
	"machine"
	"time"
)

// Device wraps an I2C connection to an IST8308 or IST8310 device.
type Device struct {
	bus        *machine.I2C
	deviceType byte
	Address    uint16
}

// New creates a new IST83xx connection. The I2C bus must already be configured.
//
// This function only creates the Device object, it does not touch the device.
func New(bus *machine.I2C) Device {
	return Device{
		bus: bus,
	}
}

// Configure sets up the device, attempting to autodetect either an IST8308 or IST8310.
func (d *Device) Configure() error {
	addresses := []uint16{IST8308_I2C_ADDRESS_DEFAULT, IST8310_I2C_ADDRESS_DEFAULT}

	for _, addr := range addresses {
		d.setAddress(addr)
		if deviceType, found := d.probe(); found {
			d.deviceType = deviceType
			return d.configure()
		}
	}

	return errors.New("IST83xx not detected")
}

// Connected returns whether an IST83xx device has been found.
func (d *Device) Connected() bool {
	return d.deviceType != 0
}

// probe reads the WAI register and identifies which device is present.
func (d *Device) probe() (byte, bool) {
	val, err := d.readRegister(IST8308_RegisterWAI)
	if err != nil {
		return 0, false
	}
	if val == IST8308_DeviceID {
		return IST8308, true
	}
	if val == IST8310_DeviceID {
		return IST8310, true
	}
	return 0, false
}

// configure calls the appropriate setup method for the detected device.
func (d *Device) configure() error {
	switch d.deviceType {
	case IST8308:
		return d.configureIST8308()
	case IST8310:
		return d.configureIST8310()
	default:
		return errors.New("unknown device type")
	}
}

// setAddress sets the I2C address for the device.
func (d *Device) setAddress(addr uint16) {
	d.Address = addr
}

// readRegister reads a 1-byte register value.
func (d *Device) readRegister(reg uint8) (uint8, error) {
	data := []byte{0}
	err := d.bus.Tx(uint16(d.Address), []byte{reg}, data)
	if err != nil {
		return 0, err
	}
	return data[0], nil
}

// writeRegister writes a 1-byte value to a register.
func (d *Device) writeRegister(reg uint8, val uint8) error {
	return d.bus.Tx(uint16(d.Address), append([]byte{reg}, val), nil)
}

// readRegisters reads a block of registers starting from the specified address.
func (d *Device) readRegisters(reg uint8, len int) ([]byte, error) {
	data := make([]byte, len)
	err := d.bus.Tx(uint16(d.Address), []byte{reg}, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// configureIST8308 sets up the IST8308 device.
func (d *Device) configureIST8308() error {
	d.writeRegister(IST8308_RegisterCNTL3, IST8308_CNTL3_BIT_SRST)
	time.Sleep(50 * time.Millisecond)
	d.writeRegister(IST8308_RegisterCNTL2, 0x06)
	d.writeRegister(IST8308_RegisterCNTL4, 0x01)
	d.writeRegister(IST8308_RegisterOSRCNTL, 0x55)
	return nil
}

// configureIST8310 sets up the IST8310 device.
func (d *Device) configureIST8310() error {
	d.writeRegister(IST8310_RegisterCNTL2, IST8310_CNTL2_BIT_SRST)
	time.Sleep(50 * time.Millisecond)
	d.writeRegister(IST8310_RegisterCNTL3, IST8310_CNTL3_BIT_X_16BIT|IST8310_CNTL3_BIT_Y_16BIT|IST8310_CNTL3_BIT_Z_16BIT)
	d.writeRegister(IST8310_RegisterAVGCNTL, 0x44)
	d.writeRegister(IST8310_RegisterPDCNTL, 0xC0)
	return nil
}

// ReadMagnetometer reads the magnetometer values for X, Y, and Z axes.
func (d *Device) ReadMagnetometer() (x, y, z int16, err error) {
	switch d.deviceType {
	case IST8308:
		return d.readMagnetometerIST8308()
	case IST8310:
		return d.readMagnetometerIST8310()
	default:
		return 0, 0, 0, errors.New("device not configured")
	}
}

// readMagnetometerIST8308 reads and processes data from the IST8308.
func (d *Device) readMagnetometerIST8308() (x, y, z int16, err error) {
	buf, err := d.readRegisters(IST8308_RegisterSTAT, 7)
	if err != nil {
		return 0, 0, 0, err
	}

	if (buf[0] & IST8308_STAT_BIT_DRDY) == 0 {
		return 0, 0, 0, errors.New("data not ready")
	}

	x = int16(buf[2])<<8 | int16(buf[1])
	y = int16(buf[4])<<8 | int16(buf[3])
	z = int16(buf[6])<<8 | int16(buf[5])
	if z == -32768 {
		z = 32767
	} else {
		z = -z
	}

	return
}

// readMagnetometerIST8310 reads and processes data from the IST8310.
func (d *Device) readMagnetometerIST8310() (x, y, z int16, err error) {
	d.writeRegister(IST8310_RegisterCNTL1, IST8310_CNTL1_BIT_MODE_SINGLE_MEASUREMENT)
	time.Sleep(20 * time.Millisecond)

	buf, err := d.readRegisters(IST8310_RegisterSTAT1, 7)
	if err != nil {
		return 0, 0, 0, err
	}

	if (buf[0] & IST8310_STAT1_BIT_DRDY) == 0 {
		return 0, 0, 0, errors.New("data not ready")
	}

	x = int16(buf[2])<<8 | int16(buf[1])
	y = int16(buf[4])<<8 | int16(buf[3])
	z = int16(buf[6])<<8 | int16(buf[5])
	if z == -32768 {
		z = 32767
	} else {
		z = -z
	}

	return
}
