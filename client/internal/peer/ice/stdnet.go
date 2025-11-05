//go:build !android

package ice

import (
	"github.com/Bee-Bros-Software/r-vpn/client/internal/stdnet"
)

func newStdNet(_ stdnet.ExternalIFaceDiscover, ifaceBlacklist []string) (*stdnet.Net, error) {
	return stdnet.NewNet(ifaceBlacklist)
}
