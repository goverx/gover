#!/usr/bin/env bash
set -e

REPO="goverx/gover" # поменяй на настоящий репозиторий
INSTALL_DIR="/usr/local/bin"

# Определяем OS и архитектуру
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
  x86_64) ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *) echo "❌ Unsupported architecture: $ARCH"; exit 1 ;;
esac

echo "👉 Detecting latest release of $REPO..."
TAG=$(curl -s https://api.github.com/repos/$REPO/releases/latest \
  | grep '"tag_name":' \
  | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$TAG" ]; then
  echo "❌ Could not fetch latest release tag"
  exit 1
fi

echo "✅ Latest version: $TAG"

URL="https://github.com/$REPO/releases/download/$TAG/gover-${OS}-${ARCH}.tar.gz"
TMP_DIR=$(mktemp -d)

echo "📥 Downloading $URL..."
curl -L "$URL" -o "$TMP_DIR/gover.tar.gz"

echo "📦 Extracting..."
tar -xzf "$TMP_DIR/gover.tar.gz" -C "$TMP_DIR"

echo "🚀 Installing to $INSTALL_DIR"
sudo mv "$TMP_DIR/gover" "$INSTALL_DIR/gover"
sudo chmod +x "$INSTALL_DIR/gover"

rm -rf "$TMP_DIR"

echo "🎉 Installed gover $TAG"
echo "Run: gover --help"
