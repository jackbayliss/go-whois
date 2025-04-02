
# go-whois

A simple tool to look up the IP address for any domain using Go.

## Installation

### Getting Started

To easily get the latest release, use GoReleaser's pre-built binaries. You can download the appropriate release for your operating system from the [Releases](https://github.com/yourusername/go-whois/releases) page.

Alternatively, if you have Go installed, you can build it from source:

```bash
go install github.com/jackbayliss/go-whois@latest
```

This will download and install the latest version of `go-whois` and place it in your `$GOPATH/bin` (or `$HOME/go/bin` for Go 1.17+).

### Windows Setup

1. Download the latest Windows release from the [Releases](https://github.com/jackbayliss/go-whois/releases) page.
2. Extract the `.exe` file and add its location to your system's PATH.

After this, you can call the tool from anywhere in your terminal.

---

## Usage

Once installed, you can simply run the command to look up IP addresses for any domain:

```bash
go-whois google.com
```

This will return the following output:

```
IP addresses for google.com:
142.250.178.14
2a00:1450:4009:81d::200e
```

---
