//go:build !(linux && 386) && !windows

package main

import (
	_ "embed"
)

//go:embed assets/rvpn.png
var iconAbout []byte

//go:embed assets/rvpn-systemtray-connected.png
var iconConnected []byte

//go:embed assets/rvpn-systemtray-connected-dark.png
var iconConnectedDark []byte

//go:embed assets/rvpn-systemtray-disconnected.png
var iconDisconnected []byte

//go:embed assets/rvpn-systemtray-update-disconnected.png
var iconUpdateDisconnected []byte

//go:embed assets/rvpn-systemtray-update-disconnected-dark.png
var iconUpdateDisconnectedDark []byte

//go:embed assets/rvpn-systemtray-update-connected.png
var iconUpdateConnected []byte

//go:embed assets/rvpn-systemtray-update-connected-dark.png
var iconUpdateConnectedDark []byte

//go:embed assets/rvpn-systemtray-connecting.png
var iconConnecting []byte

//go:embed assets/rvpn-systemtray-connecting-dark.png
var iconConnectingDark []byte

//go:embed assets/rvpn-systemtray-error.png
var iconError []byte

//go:embed assets/rvpn-systemtray-error-dark.png
var iconErrorDark []byte
