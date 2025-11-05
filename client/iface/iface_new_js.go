package iface

import (
	"github.com/Bee-Bros-Software/r-vpn/client/iface/bind"
	"github.com/Bee-Bros-Software/r-vpn/client/iface/device"
	"github.com/Bee-Bros-Software/r-vpn/client/iface/netstack"
	"github.com/Bee-Bros-Software/r-vpn/client/iface/wgaddr"
	"github.com/Bee-Bros-Software/r-vpn/client/iface/wgproxy"
)

// NewWGIFace creates a new WireGuard interface for WASM (always uses netstack mode)
func NewWGIFace(opts WGIFaceOpts) (*WGIface, error) {
	wgAddress, err := wgaddr.ParseWGAddress(opts.Address)
	if err != nil {
		return nil, err
	}

	relayBind := bind.NewRelayBindJS()

	wgIface := &WGIface{
		tun:            device.NewNetstackDevice(opts.IFaceName, wgAddress, opts.WGPort, opts.WGPrivKey, opts.MTU, relayBind, netstack.ListenAddr()),
		userspaceBind:  true,
		wgProxyFactory: wgproxy.NewUSPFactory(relayBind, opts.MTU),
	}

	return wgIface, nil
}
