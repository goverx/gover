#!/usr/bin/env bash
set -e

REPO="goverx/gover"
BINARY="gover"

echo "👉 Detecting latest release of $REPO..."
LATEST=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" | jq -r .tag_name)
echo "✅ Latest version: $LATEST"

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64) ARCH="amd64" ;;
    arm64)  ARCH="arm64" ;;
    aarch64) ARCH="arm64" ;;
esac

URL="https://github.com/$REPO/releases/download/$LATEST/${BINARY}-${OS}-${ARCH}"
echo "📥 Downloading $URL..."

TMP=$(mktemp)
curl -fsSL "$URL" -o "$TMP"
chmod +x "$TMP"

INSTALL_PATH="/usr/local/bin/$BINARY"

if [ -w "$(dirname "$INSTALL_PATH")" ]; then
    mv "$TMP" "$INSTALL_PATH"
else
    echo "⚠️  No write permission, using sudo..."
    sudo mv "$TMP" "$INSTALL_PATH"
fi

echo "✅ Installed at $INSTALL_PATH"
"$INSTALL_PATH" --version

SHELL_RC="$HOME/.zshrc"
if ! grep -q "alias gover=" "$SHELL_RC" 2>/dev/null; then
    echo "alias gover='$INSTALL_PATH'" >> "$SHELL_RC"
    echo "✅ Alias 'gover' added to $SHELL_RC (reload shell to use it)"
fi
