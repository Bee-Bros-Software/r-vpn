//go:build !linux || android

package server

import (
	"github.com/Bee-Bros-Software/r-vpn/client/internal/dns"
	"github.com/Bee-Bros-Software/r-vpn/client/internal/routemanager/systemops"
	"github.com/Bee-Bros-Software/r-vpn/client/internal/statemanager"
)

func registerStates(mgr *statemanager.Manager) {
	mgr.RegisterState(&dns.ShutdownState{})
	mgr.RegisterState(&systemops.ShutdownState{})
}
