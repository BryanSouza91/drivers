package ist83xx

// Device types
const (
	IST8308 = 0x08
	IST8310 = 0x10
)

// IST8308 I2C Address and Device ID
const (
	IST8308_I2C_ADDRESS_DEFAULT = 0x0C
	IST8308_DeviceID            = 0x08
)

// IST8308 Register Addresses
const (
	IST8308_RegisterWAI     = 0x00
	IST8308_RegisterSTAT    = 0x10
	IST8308_RegisterDATAXL  = 0x11
	IST8308_RegisterDATAXH  = 0x12
	IST8308_RegisterDATAYL  = 0x13
	IST8308_RegisterDATAYH  = 0x14
	IST8308_RegisterDATAZL  = 0x15
	IST8308_RegisterDATAZH  = 0x16
	IST8308_RegisterCNTL1   = 0x30
	IST8308_RegisterCNTL2   = 0x31
	IST8308_RegisterCNTL3   = 0x32
	IST8308_RegisterCNTL4   = 0x34
	IST8308_RegisterOSRCNTL = 0x41
)

// IST8308 Bit Masks
const (
	IST8308_STAT_BIT_DRDY  = 0x01
	IST8308_CNTL3_BIT_SRST = 0x01
)

// IST8310 I2C Address and Device ID
const (
	IST8310_I2C_ADDRESS_DEFAULT = 0x0E
	IST8310_DeviceID            = 0x10
)

// IST8310 Register Addresses
const (
	IST8310_RegisterWAI     = 0x00
	IST8310_RegisterSTAT1   = 0x02
	IST8310_RegisterDATAXL  = 0x03
	IST8310_RegisterDATAXH  = 0x04
	IST8310_RegisterDATAYL  = 0x05
	IST8310_RegisterDATAYH  = 0x06
	IST8310_RegisterDATAZL  = 0x07
	IST8310_RegisterDATAZH  = 0x08
	IST8310_RegisterCNTL1   = 0x0A
	IST8310_RegisterCNTL2   = 0x0B
	IST8310_RegisterCNTL3   = 0x0D
	IST8310_RegisterAVGCNTL = 0x41
	IST8310_RegisterPDCNTL  = 0x42
)

// IST8310 Bit Masks
const (
	IST8310_STAT1_BIT_DRDY                    = 0x01
	IST8310_CNTL1_BIT_MODE_SINGLE_MEASUREMENT = 0x01
	IST8310_CNTL2_BIT_SRST                    = 0x01
	IST8310_CNTL3_BIT_X_16BIT                 = 0x10
	IST8310_CNTL3_BIT_Y_16BIT                 = 0x20
	IST8310_CNTL3_BIT_Z_16BIT                 = 0x40
)
