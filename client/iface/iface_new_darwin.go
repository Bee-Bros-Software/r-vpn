//go:build !ios

package iface

import (
	"github.com/Bee-Bros-Software/r-vpn/client/iface/bind"
	"github.com/Bee-Bros-Software/r-vpn/client/iface/device"
	"github.com/Bee-Bros-Software/r-vpn/client/iface/netstack"
	"github.com/Bee-Bros-Software/r-vpn/client/iface/wgaddr"
	"github.com/Bee-Bros-Software/r-vpn/client/iface/wgproxy"
)

// NewWGIFace Creates a new WireGuard interface instance
func NewWGIFace(opts WGIFaceOpts) (*WGIface, error) {
	wgAddress, err := wgaddr.ParseWGAddress(opts.Address)
	if err != nil {
		return nil, err
	}

	iceBind := bind.NewICEBind(opts.TransportNet, opts.FilterFn, wgAddress, opts.MTU)

	var tun WGTunDevice
	if netstack.IsEnabled() {
		tun = device.NewNetstackDevice(opts.IFaceName, wgAddress, opts.WGPort, opts.WGPrivKey, opts.MTU, iceBind, netstack.ListenAddr())
	} else {
		tun = device.NewTunDevice(opts.IFaceName, wgAddress, opts.WGPort, opts.WGPrivKey, opts.MTU, iceBind)
	}

	wgIFace := &WGIface{
		userspaceBind:  true,
		tun:            tun,
		wgProxyFactory: wgproxy.NewUSPFactory(iceBind, opts.MTU),
	}
	return wgIFace, nil
}
