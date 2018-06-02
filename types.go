package libusb

//#cgo CFLAGS:-Iinclude/libusb-1.0
//#cgo LDFLAGS:-LMinGW64/static -lusb-1.0
//#include "libusb.h"
import "C"

type Context C.libusb_context
type Device C.libusb_device

type DeviceDescriptor struct {
	/** Size of this descriptor (in bytes) */
	//uint8_t  bLength;
	Length byte
	/** Descriptor type. Will have value
	 * \ref libusb_descriptor_type::LIBUSB_DT_DEVICE LIBUSB_DT_DEVICE in this
	 * context. */
	//	uint8_t bDescriptorType
	DescriptorType byte
	/** USB specification release number in binary-coded decimal. A value of
	 * 0x0200 indicates USB 2.0, 0x0110 indicates USB 1.1, etc. */
	//	uint16_t bcdUSB
	BcdUSB uint16
	/** USB-IF class code for the device. See \ref libusb_class_code. */
	//	uint8_t bDeviceClass
	DeviceClass byte
	/** USB-IF subclass code for the device, qualified by the bDeviceClass
	 * value */
	//	uint8_t bDeviceSubClass
	DeviceSubClass byte
	/** USB-IF protocol code for the device, qualified by the bDeviceClass and
	 * bDeviceSubClass values */
	//	uint8_t bDeviceProtocol
	DeviceProtocol byte

	/** Maximum packet size for endpoint 0 */
	//	uint8_t bMaxPacketSize0
	MaxPacketSize0 byte
	/** USB-IF vendor ID */
	//	uint16_t idVendor
	IdVendor uint16
	/** USB-IF product ID */
	//	uint16_t idProduct
	IdProduct uint16
	/** Device release number in binary-coded decimal */
	//	uint16_t bcdDevice
	BcdDevice uint16

	/** Index of string descriptor describing manufacturer */
	//	uint8_t iManufacturer
	Manufacturer byte
	/** Index of string descriptor describing product */
	//	uint8_t iProduct
	Product byte
	/** Index of string descriptor containing device serial number */
	//	uint8_t iSerialNumber
	SerialNumber byte
	/** Number of possible configurations */
	//	uint8_t bNumConfigurations
	NumConfigurations byte
}
