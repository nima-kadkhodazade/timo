#!/usr/bin/env bash

# ------------------------
# Color
# ------------------------

FMT_RED=$(printf '\033[31m')
FMT_GREEN=$(printf '\033[32m')
FMT_YELLOW=$(printf '\033[33m')
FMT_BLUE=$(printf '\033[34m')
FMT_CYAN=$(printf '\033[0;36m')
FMT_BOLD=$(printf '\033[1m')
FMT_RESET=$(printf '\033[0m')

set -euo pipefail

REPO="nima-kadkhodazade/timo"
INSTALL_DIR="${INSTALL_DIR:-$HOME/.local/bin}"

echo "${FMT_BLUE}Installing Timo...${FMT_RESET}"

# -------------------------
# Detect OS
# -------------------------

OS="$(uname -s)"

case "$OS" in
    Linux)
        OS_NAME="linux"
        ;;

    Darwin)
        OS_NAME="macos"
        ;;

    *)
        echo "${FMT_RED}Error: Unsupported operating system: $OS ${FMT_RESET}"
        exit 1
        ;;
esac

# -------------------------
# Detect architecture
# -------------------------

ARCH="$(uname -m)"

case "$ARCH" in
    x86_64|amd64)
        ARCH_NAME="amd64"
        ;;

    arm64|aarch64)
        ARCH_NAME="arm64"
        ;;
    *)
        echo "Error: Unsupported architecture: $ARCH"
        exit 1
        ;;
esac


# -------------------------
# Get latest release
# -------------------------
LATEST_RELEASE_URL="https://api.github.com/repos/$REPO/releases/latest"

VERSION="$(
    curl -fsSL "$LATEST_RELEASE_URL" |
    grep '"tag_name":' |
    head -n 1 |
    sed -E 's/.*"tag_name": "([^"]+)".*/\1/'
)"

if [ -z "$VERSION" ]; then
    echo "Error: Could not determine latest Timo version."
    exit 1
fi

echo "Latest Version: $VERSION"


# -------------------------
# Build asset names
# -------------------------

ARCHIVE="timo-${OS_NAME}-${ARCH_NAME}.tar.gz"

DOWNLOAD_URL="https://github.com/$REPO/releases/download/$VERSION/$ARCHIVE"

CHECKSUMS_URL="https://github.com/$REPO/releases/download/$VERSION/checksum.txt"

# -------------------------
# Temporary directory
# -------------------------

TMP_DIR="$(mktemp -d)"

cleanup() {
    rm -rf "$TMP_DIR"
}

trap cleanup EXIT

echo "${FMT_BLUE}Downloading $ARCHIVE...${FMT_RESET}"

curl -fL "$DOWNLOAD_URL" -o "$TMP_DIR/$ARCHIVE"
curl -fL "$CHECKSUMS_URL" -o "$TMP_DIR/checksum.txt"

# -------------------------
# Verify checksum
# -------------------------

echo "${FMT_YELLOW}Verifying checksum...${FMT_RESET}"

cd "$TMP_DIR"

EXPECTED_CHECKSUM="$(
    grep " $ARCHIVE$" checksum.txt |
    awk '{print $1}'
)"

if [ -z "$EXPECTED_CHECKSUM" ]; then
    echo "Error: Could not find checksum for $ARCHIVE."
    exit 1
fi

if command -v sha256sum >/dev/null 2>&1; then

    echo "$EXPECTED_CHECKSUM  $ARCHIVE" |
        sha256sum -c -

elif command -v shasum >/dev/null 2>&1; then

    ACTUAL_CHECKSUM="$(shasum -a 256 "$ARCHIVE" | awk '{print $1}')"

    if [ "$EXPECTED_CHECKSUM" != "$ACTUAL_CHECKSUM" ]; then
        echo "Error: Checksum verification failed."
        exit 1
    fi

else
    echo "Error: No SHA-256 utility found."
    exit 1
fi

# -------------------------
# Extract
# -------------------------

echo "${FMT_BLUE}Extracting...${FMT_RESET}"

tar -xzf "$ARCHIVE"

# -------------------------
# Install
# -------------------------

mkdir -p "$INSTALL_DIR"

mv timo "$INSTALL_DIR/timo"

chmod +x "$INSTALL_DIR/timo"

echo
echo "${FMT_GREEN}Timo $VERSION installed successfully!${FMT_RESET}"
echo
echo "Installed to:"
echo "$INSTALL_DIR/timo"
echo

# -------------------------
# PATH check
# -------------------------

case ":$PATH:" in
    *":$INSTALL_DIR:"*)
        echo "You can now run:"
        echo "  timo --v"
        ;;

    *)
        echo "WARNING: $INSTALL_DIR is not in your PATH."
        echo
        echo "Add this line to your shell configuration:"
        echo
        echo "export PATH=\"\$HOME/.local/bin:\$PATH\""
        echo
        echo "Then restart your shell."
        ;;
esac
