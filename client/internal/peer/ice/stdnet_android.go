package ice

import "github.com/Bee-Bros-Software/r-vpn/client/internal/stdnet"

func newStdNet(iFaceDiscover stdnet.ExternalIFaceDiscover, ifaceBlacklist []string) (*stdnet.Net, error) {
	return stdnet.NewNetWithDiscover(iFaceDiscover, ifaceBlacklist)
}
