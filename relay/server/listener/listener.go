package listener

import (
	"context"
	"net"

	"github.com/Bee-Bros-Software/r-vpn/relay/protocol"
)

type Listener interface {
	Listen(func(conn net.Conn)) error
	Shutdown(ctx context.Context) error
	Protocol() protocol.Protocol
}
