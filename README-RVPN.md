# R-VPN

**R-VPN** is a branded VPN client for R-Software infrastructure, based on [NetBird](https://github.com/netbirdio/netbird).

## About

R-VPN provides secure, encrypted access to R-Software's private infrastructure using WireGuard® technology. It features:

- **Zero-configuration VPN**: Automatic peer discovery and connection
- **OIDC Authentication**: Seamless login via Authentik (auth.rsoftware.net)
- **Cross-platform**: macOS, iOS, Windows, Linux, Android
- **Private networking**: Connect to infrastructure without exposing public IPs
- **End-to-end encrypted**: Direct peer-to-peer connections when possible

## Pre-configured Settings

R-VPN comes pre-configured for R-Software infrastructure:
- **Management Server**: `https://vpn.rsoftware.net:443`
- **Authentication**: Authentik OIDC at `https://auth.rsoftware.net`

## Installation

### macOS
Download the latest `.dmg` from [Releases](https://github.com/Bee-Bros-Software/r-vpn/releases)

### iOS
Coming soon to TestFlight

### Building from Source

Requirements:
- Go 1.21+
- Platform-specific build tools (Xcode for macOS/iOS)

```bash
# Clone the repository
git clone https://github.com/Bee-Bros-Software/r-vpn.git
cd r-vpn

# Build for macOS
cd client
go build -o rvpn

# Build UI client for macOS
cd ui
go build -o rvpn-ui
```

## Usage

1. Launch R-VPN
2. Click "Connect"
3. Authenticate via Authentik when prompted
4. You're connected!

R-VPN runs in your system tray and automatically reconnects when network conditions change.

## Architecture

R-VPN uses a hybrid mesh architecture:
- **Direct connections** between peers when possible (NAT traversal)
- **Relay connections** via TURN server when direct connections fail
- **Central management** server for configuration and discovery

## Branding

R-VPN is customized with R-Software branding:
- App name: **R-VPN**
- Bundle ID: `net.rsoftware.rvpn` (macOS), `net.rsoftware.rvpn.ios` (iOS)
- Colors: R-Software blue gradient (#0ea5e9 to #0369a1)
- Logo: R-Software "R" logo with sparkles

## Upstream

R-VPN is based on NetBird v0.59.11. We maintain this fork to provide a branded experience for R-Software users.

To sync with upstream NetBird:
```bash
git remote add netbird https://github.com/netbirdio/netbird.git
git fetch netbird
git merge netbird/main
```

## Development

See [CONTRIBUTING.md](CONTRIBUTING.md) for development guidelines.

## License

R-VPN inherits the BSD 3-Clause License from NetBird. See [LICENSE](LICENSE) for details.

## Support

For R-Software infrastructure support:
- Email: support@rsoftware.net
- Documentation: https://docs.rsoftware.net

For R-VPN client issues:
- GitHub Issues: https://github.com/Bee-Bros-Software/r-vpn/issues

---

Built with ❤️ by [R-Software](https://rsoftware.net)
