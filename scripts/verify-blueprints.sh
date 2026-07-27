#!/usr/bin/env bash
set -euo pipefail

echo "========================================="
echo " Verifying OpenTofu HCL Blueprints..."
echo "========================================="

for dir in blueprints/*/; do
    if [ -d "$dir" ]; then
        echo "Validating blueprint: $dir"
        (cd "$dir" && tofu fmt -check && tofu init -backend=false && tofu validate)
    fi
done

echo "All OpenTofu Blueprints Verified!"
