package iface

import (
	"net"
	"net/netip"

	"github.com/Bee-Bros-Software/r-vpn/client/iface/device"
	"github.com/Bee-Bros-Software/r-vpn/client/iface/wgaddr"
)

type wgIfaceBase interface {
	AddAllowedIP(peerKey string, allowedIP netip.Prefix) error
	RemoveAllowedIP(peerKey string, allowedIP netip.Prefix) error

	Name() string
	Address() wgaddr.Address
	ToInterface() *net.Interface
	IsUserspaceBind() bool
	GetFilter() device.PacketFilter
	GetDevice() *device.FilteredDevice
}
