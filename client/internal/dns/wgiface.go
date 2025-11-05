//go:build !windows

package dns

import (
	"net"

	"github.com/Bee-Bros-Software/r-vpn/client/iface/device"
	"github.com/Bee-Bros-Software/r-vpn/client/iface/wgaddr"
)

// WGIface defines subset methods of interface required for manager
type WGIface interface {
	Name() string
	Address() wgaddr.Address
	ToInterface() *net.Interface
	IsUserspaceBind() bool
	GetFilter() device.PacketFilter
	GetDevice() *device.FilteredDevice
}
