package desktop

import "github.com/Bee-Bros-Software/r-vpn/version"

// GetUIUserAgent returns the Desktop ui user agent
func GetUIUserAgent() string {
	return "rvpn-desktop-ui/" + version.NetbirdVersion()
}
