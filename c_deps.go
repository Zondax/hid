// +build dummy

// This file is part of a workaround for `go mod vendor` which won't vendor
// C files if there's no Go file in the same directory.
// This would prevent the bundled c-libs from being removed when vendored.
//
// This Go file imports the c directory and subdirectories where there is
// another go (named after the directory) file which is the second part
// of this workaround.
//
// These files combined make it so `go mod vendor` behaves correctly.
// See this issue for reference: https://github.com/golang/go/issues/26366

package hid

import (
	_ "github.com/zondax/hid/hidapi"
	_ "github.com/zondax/hid/hidapi/hidapi"
	_ "github.com/zondax/hid/hidapi/libusb"
	_ "github.com/zondax/hid/hidapi/linux"
	_ "github.com/zondax/hid/hidapi/mac"
	_ "github.com/zondax/hid/hidapi/windows"

	_ "github.com/zondax/hid/libusb"
	_ "github.com/zondax/hid/libusb/libusb"
	_ "github.com/zondax/hid/libusb/libusb/os"
)
