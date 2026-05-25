#!/usr/bin/env bash
set -euo pipefail

root="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$root"

version="dev"
if git rev-parse --git-dir >/dev/null 2>&1; then
  version="$(git describe --tags --always --dirty 2>/dev/null || echo dev)"
fi

mkdir -p .bin
go build -ldflags="-X github.com/ErfanEbrahimnia/pm/internal/version.Version=${version}" -o .bin/pm .
echo "built $root/.bin/pm"
