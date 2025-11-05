//go:build !linux || android

package conntrack

import nftypes "github.com/Bee-Bros-Software/r-vpn/client/internal/netflow/types"

func New(flowLogger nftypes.FlowLogger, iface nftypes.IFaceMapper) nftypes.ConnTracker {
	return nil
}
