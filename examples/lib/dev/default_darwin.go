package dev

import (
	"github.com/hybridgroup/ble"
	"github.com/hybridgroup/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice() (d ble.Device, err error) {
	return darwin.NewDevice()
}
