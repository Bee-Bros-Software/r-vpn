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

	if netstack.IsEnabled() {
		wgIFace := &WGIface{
			userspaceBind:  true,
			tun:            device.NewNetstackDevice(opts.IFaceName, wgAddress, opts.WGPort, opts.WGPrivKey, opts.MTU, iceBind, netstack.ListenAddr()),
			wgProxyFactory: wgproxy.NewUSPFactory(iceBind, opts.MTU),
		}
		return wgIFace, nil
	}

	wgIFace := &WGIface{
		userspaceBind:  true,
		tun:            device.NewTunDevice(wgAddress, opts.WGPort, opts.WGPrivKey, opts.MTU, iceBind, opts.MobileArgs.TunAdapter, opts.DisableDNS),
		wgProxyFactory: wgproxy.NewUSPFactory(iceBind, opts.MTU),
	}
	return wgIFace, nil
}
