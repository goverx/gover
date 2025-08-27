#!/usr/bin/env bash
set -e

REPO="goverx/gover"
BINARY_NAME="gover"

echo "👉 Detecting latest release of $REPO..."
LATEST=$(curl -s https://api.github.com/repos/$REPO/releases/latest | grep tag_name | cut -d '"' -f4)
echo "✅ Latest version: $LATEST"

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64) ARCH=amd64 ;;
    arm64) ARCH=arm64 ;;
esac

FILE="${BINARY_NAME}-${OS}-${ARCH}"

URL="https://github.com/$REPO/releases/download/$LATEST/$FILE"

echo "📥 Downloading $URL..."
curl -L "$URL" -o "/usr/local/bin/$BINARY_NAME"

chmod +x "/usr/local/bin/$BINARY_NAME"
echo "✅ Installed $BINARY_NAME to /usr/local/bin/$BINARY_NAME"
echo "👉 Run '$BINARY_NAME --help'"
