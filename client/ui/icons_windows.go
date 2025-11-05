package main

import (
	_ "embed"
)

//go:embed assets/rvpn.ico
var iconAbout []byte

//go:embed assets/rvpn-systemtray-connected.ico
var iconConnected []byte

//go:embed assets/rvpn-systemtray-connected-dark.ico
var iconConnectedDark []byte

//go:embed assets/rvpn-systemtray-disconnected.ico
var iconDisconnected []byte

//go:embed assets/rvpn-systemtray-update-disconnected.ico
var iconUpdateDisconnected []byte

//go:embed assets/rvpn-systemtray-update-disconnected-dark.ico
var iconUpdateDisconnectedDark []byte

//go:embed assets/rvpn-systemtray-update-connected.ico
var iconUpdateConnected []byte

//go:embed assets/rvpn-systemtray-update-connected-dark.ico
var iconUpdateConnectedDark []byte

//go:embed assets/rvpn-systemtray-connecting.ico
var iconConnecting []byte

//go:embed assets/rvpn-systemtray-connecting-dark.ico
var iconConnectingDark []byte

//go:embed assets/rvpn-systemtray-error.ico
var iconError []byte

//go:embed assets/rvpn-systemtray-error-dark.ico
var iconErrorDark []byte
