#!/usr/bin/env bash
set -e

REPO="goverx/gover"
INSTALL_DIR="/usr/local/bin"

echo "👉 Detecting latest release of $REPO..."
VERSION=$(curl -s https://api.github.com/repos/$REPO/releases/latest | grep -Po '"tag_name": "\K.*?(?=")')
echo "✅ Latest version: $VERSION"

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64) ARCH=amd64 ;;
    arm64)  ARCH=arm64 ;;
    aarch64) ARCH=arm64 ;;
esac

BINARY="gover-${OS}-${ARCH}"
URL="https://github.com/$REPO/releases/download/$VERSION/$BINARY"

TMP=$(mktemp)
echo "📥 Downloading $URL..."
curl -fsSL "$URL" -o "$TMP"

chmod +x "$TMP"
echo "🔑 Moving binary to $INSTALL_DIR (may require sudo)..."
sudo mv "$TMP" "$INSTALL_DIR/gover"

echo "🎉 Installed gover $VERSION to $INSTALL_DIR/gover"
