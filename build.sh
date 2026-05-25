#!/usr/bin/env bash
set -euo pipefail

root="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$root"

mkdir -p .bin
go build -o .bin/pm .
echo "built $root/.bin/pm"
