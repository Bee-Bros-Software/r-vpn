package common

import (
	"sync/atomic"
	"time"

	"github.com/Bee-Bros-Software/r-vpn/client/firewall/manager"
	"github.com/Bee-Bros-Software/r-vpn/client/internal/dns"
	"github.com/Bee-Bros-Software/r-vpn/client/internal/peer"
	"github.com/Bee-Bros-Software/r-vpn/client/internal/peerstore"
	"github.com/Bee-Bros-Software/r-vpn/client/internal/routemanager/fakeip"
	"github.com/Bee-Bros-Software/r-vpn/client/internal/routemanager/iface"
	"github.com/Bee-Bros-Software/r-vpn/client/internal/routemanager/refcounter"
	"github.com/Bee-Bros-Software/r-vpn/route"
)

type HandlerParams struct {
	Route                *route.Route
	RouteRefCounter      *refcounter.RouteRefCounter
	AllowedIPsRefCounter *refcounter.AllowedIPsRefCounter
	DnsRouterInterval    time.Duration
	StatusRecorder       *peer.Status
	WgInterface          iface.WGIface
	DnsServer            dns.Server
	PeerStore            *peerstore.Store
	UseNewDNSRoute       bool
	Firewall             manager.Manager
	FakeIPManager        *fakeip.Manager
	ForwarderPort        *atomic.Uint32
}
