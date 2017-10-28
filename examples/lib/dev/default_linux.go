package dev

import (
	"github.com/hybridgroup/ble"
	"github.com/hybridgroup/ble/linux"
)

// DefaultDevice ...
func DefaultDevice() (d ble.Device, err error) {
	return linux.NewDevice()
}
