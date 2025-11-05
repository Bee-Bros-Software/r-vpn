//go:build !android

package server

import (
	"github.com/Bee-Bros-Software/r-vpn/client/firewall/iptables"
	"github.com/Bee-Bros-Software/r-vpn/client/firewall/nftables"
	"github.com/Bee-Bros-Software/r-vpn/client/internal/dns"
	"github.com/Bee-Bros-Software/r-vpn/client/internal/routemanager/systemops"
	"github.com/Bee-Bros-Software/r-vpn/client/internal/statemanager"
)

func registerStates(mgr *statemanager.Manager) {
	mgr.RegisterState(&dns.ShutdownState{})
	mgr.RegisterState(&systemops.ShutdownState{})
	mgr.RegisterState(&nftables.ShutdownState{})
	mgr.RegisterState(&iptables.ShutdownState{})
}
