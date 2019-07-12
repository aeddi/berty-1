package ble

import (
	bledrv "berty.tech/core/network/protocol/ble/driver"
)

func init() {
	// Bind native to golang bridge functions
	bledrv.BindNativeToGoFunctions(
		HandlePeerFound,
		ReceiveFromDevice,
	)
}
