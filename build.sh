#!/usr/bin/env bash
set -euo pipefail

root="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$root"

version="dev"
if git rev-parse --git-dir >/dev/null 2>&1; then
  version="$(git describe --tags --always --dirty 2>/dev/null || echo dev)"
fi

mkdir -p bin
ldflags="-X github.com/ErfanEbrahimnia/pm/internal/version.Version=${version}"
go build -ldflags="${ldflags}" -o bin/pm .
go build -ldflags="${ldflags}" -o bin/pmx ./cmd/pmx
echo "built $root/bin/pm and $root/bin/pmx"
