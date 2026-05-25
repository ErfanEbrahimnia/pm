#!/usr/bin/env bash
set -euo pipefail

REPO="${PM_REPO:-ErfanEbrahimnia/pm}"
INSTALL_DIR="${INSTALL_DIR:-${HOME}/.local/bin}"
VERSION="${PM_VERSION:-}"

usage() {
  cat <<EOF
Usage: install.sh

Install pm from GitHub releases into INSTALL_DIR (default: ~/.local/bin)
and add it to your PATH.

Environment:
  PM_REPO       GitHub repository (default: ErfanEbrahimnia/pm)
  PM_VERSION    Release tag, e.g. v1.0.0 (default: latest)
  INSTALL_DIR   Install location (default: ~/.local/bin)
EOF
}

if [[ "${1:-}" == "-h" || "${1:-}" == "--help" ]]; then
  usage
  exit 0
fi

detect_platform() {
  local os arch
  os="$(uname -s)"
  arch="$(uname -m)"

  case "$os" in
    Darwin) os="darwin" ;;
    Linux) os="linux" ;;
    *)
      echo "unsupported operating system: $os" >&2
      exit 1
      ;;
  esac

  case "$arch" in
    x86_64 | amd64) arch="amd64" ;;
    arm64 | aarch64) arch="arm64" ;;
    *)
      echo "unsupported architecture: $arch" >&2
      exit 1
      ;;
  esac

  echo "${os}-${arch}"
}

resolve_version() {
  if [[ -n "$VERSION" ]]; then
    echo "$VERSION"
    return
  fi
  curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" |
    grep -E '"tag_name"' | head -1 | sed -E 's/.*"([^"]+)".*/\1/'
}

download_release() {
  local platform="$1"
  local version="$2"
  local base="https://github.com/${REPO}/releases/download/${version}"
  local archive="pm-${platform}.tar.gz"
  local url="${base}/${archive}"
  local tmp
  tmp="$(mktemp -d)"
  trap 'rm -rf "$tmp"' EXIT

  echo "Downloading ${url}..."
  curl -fsSL "$url" -o "${tmp}/${archive}"
  tar -xzf "${tmp}/${archive}" -C "$tmp"
  mkdir -p "$INSTALL_DIR"
  install -m 0755 "${tmp}/pm-${platform}" "${INSTALL_DIR}/pm"
  echo "Installed pm to ${INSTALL_DIR}/pm"
}

ensure_path() {
  case ":$PATH:" in
    *":${INSTALL_DIR}:"*) return ;;
  esac

  local line="export PATH=\"${INSTALL_DIR}:\$PATH\""
  local updated=0

  for profile in "${HOME}/.zshrc" "${HOME}/.bashrc"; do
    if [[ -f "$profile" ]] && ! grep -qF "$INSTALL_DIR" "$profile"; then
      printf '\n# pm\n%s\n' "$line" >>"$profile"
      echo "Added ${INSTALL_DIR} to PATH in ${profile}"
      updated=1
    fi
  done

  if [[ "$updated" -eq 0 ]]; then
    echo "Add ${INSTALL_DIR} to your PATH:"
    echo "  ${line}"
  fi
}

main() {
  local platform version
  platform="$(detect_platform)"
  version="$(resolve_version)"
  echo "Installing pm ${version} for ${platform}..."
  download_release "$platform" "$version"
  ensure_path
  echo "Done. Restart your shell or run: export PATH=\"${INSTALL_DIR}:\$PATH\""
}

main "$@"
