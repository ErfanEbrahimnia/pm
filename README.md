# pm

A small CLI that forwards commands to the package manager your project already uses. Run `pm install` instead of guessing whether the repo expects `pnpm`, `npm`, `yarn`, or `bun`.

## Detection

`pm` walks up from the current directory until it finds a project root:

1. **`package.json` → `packageManager`** — set explicitly via `pm use`
2. **Lockfiles** — if several exist in the same directory, priority is: `pnpm-lock.yaml`, `bun.lockb` / `bun.lock`, `yarn.lock`, `package-lock.json`

Supported managers: **pnpm**, **bun**, **yarn**, **npm**.

## Usage

```bash
pm                  # show detected manager
pm install          # e.g. pnpm install
pm run dev
pm use              # pick a manager interactively
pm --version        # show release version

pmx                 # show detected exec tool (npx, pnpx, bunx, yarn dlx)
pmx vitest          # e.g. pnpx vitest
pmx --version       # show pmx and exec tool version
```

## Install

**Install script** (macOS and Linux):

```bash
curl -fsSL https://raw.githubusercontent.com/ErfanEbrahimnia/pm/main/install.sh | bash
```

Downloads a release binary and adds the install directory to your `PATH` in `~/.zshrc` or `~/.bashrc`.

Options:

```bash
PM_VERSION=1.0.0-beta.0 curl -fsSL .../install.sh | bash
INSTALL_DIR=/usr/local/bin curl -fsSL .../install.sh | bash
```

**From source:**

```bash
git clone https://github.com/ErfanEbrahimnia/pm.git
cd pm
./build.sh
export PATH="$(pwd)/.bin:$PATH"
```

**With Go:**

```bash
go install github.com/ErfanEbrahimnia/pm@latest
go install github.com/ErfanEbrahimnia/pm/cmd/pmx@latest
```

Prebuilt binaries for Linux and macOS (amd64 and arm64) are attached to each [GitHub release](https://github.com/ErfanEbrahimnia/pm/releases).

## Requirements

- The detected package manager installed and on your `PATH`
- For source builds: Go 1.22+

## Releasing

Create a GitHub release (tag + publish). The release workflow builds platform binaries and uploads them as release assets.
