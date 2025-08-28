# Gover

![Gover Logo](assets/logo.png)

Gover is a lightweight version manager for Go. It helps you install, switch, and manage multiple versions of Go with ease.

## Installation

### macOS / Linux

```bash
curl -fsSL https://raw.githubusercontent.com/goverx/gover/main/install.sh | sudo bash
```

### Windows (PowerShell)

```powershell
iwr -useb https://raw.githubusercontent.com/goverx/gover/main/install.ps1 | iex
```

## Usage

Check installed Go versions:

```bash
gover list
```

Install a new version of Go:

```bash
gover install 1.22.0
```

Switch to a specific version:

```bash
gover use 1.22.0
```

Check the currently active Go version:

```bash
gover current
```

## License

MIT
