# Warframe Paster

A utility program that enables clipboard pasting in Warframe using a custom keyboard shortcut (Alt + V), solving the issue where Ctrl + V doesn't work in the game.

## Why?
Warframe "sometimes" doesn't support the standard Ctrl + V clipboard paste, so it makes copying "/w" messages from warframe.market annoying

## IMPORTNANT
**⚠️ WARNING: Use at your own risk. Digital Extremes (Warframe developers) may consider automated input programs as third-party tools that could result in account penalties or bans. This tool simulates keyboard input and may be detected by anti-cheat systems. Use responsibly and be aware of the potential consequences.**

## Installation

### Prerequisites

- Go 1.24.0 or higher

### Building from Source

1. Clone or download the repository
2. Install dependencies:
   ```bash
   go mod download
   ```

### Build Options

#### Build for current platform:
```bash
go build -o warframe-paster .
```

#### Build for all platforms:
```bash
make all
```

#### Build for specific platforms:
```bash
# Linux
make build-linux

# Windows
make build-windows
```

#### Linux Installation (system-wide):
```bash
make install-linux
```

#### Linux Uninstallation:
```bash
make uninstall-linux
```

## Usage

1. Run the program
4. In Warframe (or any application), press **Alt + V** to paste the clipboard content
5. The program will automatically type the clipboard content for you

## Configuration

You can modify the following constants in [warframa-paster.go](warframa-paster.go):

- `PASTE_SHORTCUT`: Change the keyboard shortcut (default: `["alt", "v"]`)
- `PASTE_CONTENT_LIMIT`: Maximum characters to paste (default: `300`)
- `PASTE_TIMEOUT`: Delay between characters in milliseconds (default: `50`)

## Dependencies

- [github.com/go-vgo/robotgo](https://github.com/go-vgo/robotgo) - For simulating keyboard input
- [github.com/robotn/gohook](https://github.com/robotn/gohook) - For global hotkey detection
- [golang.design/x/clipboard](https://golang.design/x/clipboard) - For clipboard access

