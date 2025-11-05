//go:build linux && !android

package iface

import (
	"fmt"

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

	wgIFace := &WGIface{}

	if netstack.IsEnabled() {
		iceBind := bind.NewICEBind(opts.TransportNet, opts.FilterFn, wgAddress, opts.MTU)
		wgIFace.tun = device.NewNetstackDevice(opts.IFaceName, wgAddress, opts.WGPort, opts.WGPrivKey, opts.MTU, iceBind, netstack.ListenAddr())
		wgIFace.userspaceBind = true
		wgIFace.wgProxyFactory = wgproxy.NewUSPFactory(iceBind, opts.MTU)
		return wgIFace, nil
	}

	if device.WireGuardModuleIsLoaded() {
		wgIFace.tun = device.NewKernelDevice(opts.IFaceName, wgAddress, opts.WGPort, opts.WGPrivKey, opts.MTU, opts.TransportNet)
		wgIFace.wgProxyFactory = wgproxy.NewKernelFactory(opts.WGPort, opts.MTU)
		return wgIFace, nil
	}
	if device.ModuleTunIsLoaded() {
		iceBind := bind.NewICEBind(opts.TransportNet, opts.FilterFn, wgAddress, opts.MTU)
		wgIFace.tun = device.NewUSPDevice(opts.IFaceName, wgAddress, opts.WGPort, opts.WGPrivKey, opts.MTU, iceBind)
		wgIFace.userspaceBind = true
		wgIFace.wgProxyFactory = wgproxy.NewUSPFactory(iceBind, opts.MTU)
		return wgIFace, nil
	}

	return nil, fmt.Errorf("couldn't check or load tun module")
}
