package iface

import (
	"github.com/Bee-Bros-Software/r-vpn/client/iface/device"
)

// GetInterfaceGUIDString returns an interface GUID. This is useful on Windows only
func (w *WGIface) GetInterfaceGUIDString() (string, error) {
	return w.tun.(*device.TunDevice).GetInterfaceGUIDString()
}
