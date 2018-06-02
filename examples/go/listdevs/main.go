package main

import (
	"fmt"

	"github.com/funnywwh/libusb"
)

func main() {
	var desc libusb.DeviceDescriptor
	var ctx *libusb.Context
	if code := libusb.Init(&ctx); code != 0 {
		panic(code)
	}
	defer libusb.Exit(ctx)
	var devs []*libusb.Device
	cnt := libusb.GetDeviceList(nil, &devs)
	if cnt > 0 {
		defer libusb.FreeDeviceList(devs)
		for _, dev := range devs {
			code := libusb.GetDeviceDescriptor(dev, &desc)
			if code != 0 {
				fmt.Printf("code:%d\n", code)
				continue
			}
			fmt.Printf("%04x:%04x (bus %d, device %d)", desc.IdVendor, desc.IdProduct,
				libusb.GetBusNumber(dev), libusb.GetDeviceAddress(dev))
			ports := libusb.GetPortNumbers(dev, 256)
			for i, port := range ports {
				if i == 0 {
					fmt.Printf("path:%d", port)
				} else {
					fmt.Printf(".%d", port)
				}
			}
			fmt.Println()
		}
	}
}
