package firewall

import (
	wgdevice "golang.zx2c4.com/wireguard/device"

	"github.com/Bee-Bros-Software/r-vpn/client/iface/device"
	"github.com/Bee-Bros-Software/r-vpn/client/iface/wgaddr"
)

// IFaceMapper defines subset methods of interface required for manager
type IFaceMapper interface {
	Name() string
	Address() wgaddr.Address
	IsUserspaceBind() bool
	SetFilter(device.PacketFilter) error
	GetDevice() *device.FilteredDevice
	GetWGDevice() *wgdevice.Device
}
