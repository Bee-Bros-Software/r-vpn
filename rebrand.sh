#!/bin/bash
# R-VPN Rebranding Script

echo "Starting rebranding from NetBird to R-VPN..."

# Text replacements
find . -type f \( -name "*.go" -o -name "*.swift" -o -name "*.java" -o -name "*.xml" -o -name "*.yaml" -o -name "*.md" -o -name "*.rb" -o -name "*.json" -o -name "*.toml" \) ! -path "./.git/*" ! -path "./vendor/*" -exec sed -i '' 's/NetBird/R-VPN/g' {} +

# URL replacements
find . -type f \( -name "*.go" -o -name "*.swift" -o -name "*.java" -o -name "*.xml" -o -name "*.yaml" -o -name "*.md" \) ! -path "./.git/*" ! -path "./vendor/*" -exec sed -i '' 's/netbirdio\/netbird/Bee-Bros-Software\/r-vpn/g' {} +

find . -type f \( -name "*.go" -o -name "*.swift" -o -name "*.java" -o -name "*.xml" -o -name "*.yaml" -o -name "*.md" \) ! -path "./.git/*" ! -path "./vendor/*" -exec sed -i '' 's/netbird\.io/rsoftware.net/g' {} +

# Variable name replacements (lowercase)
find . -type f -name "*.go" ! -path "./.git/*" ! -path "./vendor/*" -exec sed -i '' 's/netbird/rvpn/g' {} +

echo "Text rebranding complete!"
