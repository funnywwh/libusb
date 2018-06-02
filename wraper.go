package libusb

import (
	"unsafe"
)

//#include "libusb.h"
//extern int getlist();
import "C"

func Init(pctx **Context) int {
	var ctx *C.libusb_context
	code := C.libusb_init(&ctx)
	if pctx != nil {
		*pctx = (*Context)(ctx)
	}
	return int(code)
}
func Exit(ctx *Context) {
	C.libusb_exit((*C.libusb_context)(ctx))
}

//libusb_get_device_list
func GetDeviceList(ctx *Context, devs *[]*Device) (cnt int) {
	var cdevs **C.libusb_device
	cnt = int(C.libusb_get_device_list((*C.libusb_context)(ctx), &cdevs))
	if cnt > 0 {
		_devs := (*[1 << 32]*Device)(unsafe.Pointer(cdevs))
		for _, d := range _devs[:cnt] {
			*devs = append(*devs, d)
		}
		C.free(unsafe.Pointer(cdevs))
	}

	return
}
func FreeDeviceList(devs []*Device) {
	for _, dev := range devs {
		C.libusb_unref_device((*C.libusb_device)(dev))
	}
}
func descFromC(desc *DeviceDescriptor, cdesc *C.struct_libusb_device_descriptor) {
	desc.BcdDevice = uint16(cdesc.bcdDevice)
	desc.BcdUSB = uint16(cdesc.bcdUSB)
	desc.DescriptorType = byte(cdesc.bDescriptorType)
	desc.DeviceClass = byte(cdesc.bDeviceClass)
	desc.DeviceProtocol = byte(cdesc.bDeviceProtocol)
	desc.DeviceSubClass = byte(cdesc.bDeviceSubClass)
	desc.IdProduct = uint16(cdesc.idProduct)
	desc.IdVendor = uint16(cdesc.idVendor)
	desc.Length = byte(cdesc.bLength)
	desc.Manufacturer = byte(cdesc.iManufacturer)
	desc.MaxPacketSize0 = byte(cdesc.bMaxPacketSize0)
	desc.NumConfigurations = byte(cdesc.bNumConfigurations)
	desc.Product = byte(cdesc.iProduct)
	desc.SerialNumber = byte(cdesc.iSerialNumber)
}

func descToC(cdesc *C.struct_libusb_device_descriptor, desc *DeviceDescriptor) {
	cdesc.bcdDevice = C.ushort(desc.BcdDevice)
	cdesc.bcdUSB = C.ushort(desc.BcdUSB)
	cdesc.bDescriptorType = C.uchar(desc.DescriptorType)
	cdesc.bDeviceClass = C.uchar(desc.DeviceClass)
	cdesc.bDeviceProtocol = C.uchar(desc.DeviceProtocol)
	cdesc.bDeviceSubClass = C.uchar(desc.DeviceSubClass)
	cdesc.idProduct = C.ushort(desc.IdProduct)
	cdesc.idVendor = C.ushort(desc.IdVendor)
	cdesc.bLength = C.uchar(desc.Length)
	cdesc.iManufacturer = C.uchar(desc.Manufacturer)
	cdesc.bMaxPacketSize0 = C.uchar(desc.MaxPacketSize0)
	cdesc.bNumConfigurations = C.uchar(desc.NumConfigurations)
	cdesc.iProduct = C.uchar(desc.Product)
	cdesc.iSerialNumber = C.uchar(desc.SerialNumber)
}
func GetDeviceDescriptor(dev *Device, desc *DeviceDescriptor) int {
	cdesc := new(C.struct_libusb_device_descriptor)
	code := C.libusb_get_device_descriptor((*C.struct_libusb_device)(dev), cdesc)
	descFromC(desc, cdesc)
	return int(code)
}

func GetBusNumber(dev *Device) (num uint8) {
	num = uint8(C.libusb_get_bus_number((*C.libusb_device)(dev)))
	return
}

func GetDeviceAddress(dev *Device) (addr uint8) {
	addr = uint8(C.libusb_get_device_address((*C.libusb_device)(dev)))
	return
}

func GetPortNumbers(dev *Device, maxPort int) (ports []uint8) {
	ports = make([]uint8, maxPort)
	r := C.libusb_get_port_numbers((*C.libusb_device)(dev), (*C.uchar)(&ports[0]), C.int(len(ports)))
	ports = ports[:r]
	return
}
