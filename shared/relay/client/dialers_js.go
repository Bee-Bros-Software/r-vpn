//go:build js

package client

import (
	"github.com/Bee-Bros-Software/r-vpn/shared/relay/client/dialer"
	"github.com/Bee-Bros-Software/r-vpn/shared/relay/client/dialer/ws"
)

func (c *Client) getDialers() []dialer.DialeFn {
	// JS/WASM build only uses WebSocket transport
	return []dialer.DialeFn{ws.Dialer{}}
}
