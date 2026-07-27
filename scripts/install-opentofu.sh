#!/usr/bin/env bash
set -euo pipefail

echo "========================================="
echo " Installing OpenTofu Binary CLI..."
echo "========================================="

# Download and install OpenTofu repository installer
curl --proto '=https' --tlsv1.2 -fsSL https://get.opentofu.org/install-opentofu.sh -o install-opentofu.sh
chmod +x install-opentofu.sh
./install-opentofu.sh --tofu-version 1.6.2
rm -f install-opentofu.sh

tofu --version
echo "OpenTofu Binary Installed Successfully!"
