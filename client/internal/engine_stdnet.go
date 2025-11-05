//go:build !android

package internal

import (
	"github.com/Bee-Bros-Software/r-vpn/client/internal/stdnet"
)

func (e *Engine) newStdNet() (*stdnet.Net, error) {
	return stdnet.NewNet(e.config.IFaceBlackList)
}
