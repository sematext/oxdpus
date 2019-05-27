# oxdpus
o**xdp**us is a toy tool that demonstrates some of the super powers of [XDP](https://www.iovisor.org/technology/xdp) - a high performance packet processing path built into the kernel.


## Requirements

To build oxdpus you have to satisify the following requirements:
- have a modern Linux kernel (4.12>) that supports XDP
- linux headers
- clang
- LLVM
- Go 1.12>
- gobindata (to embed XDP bytecode inside Go binary)

This repository ships with `Makefile` to facilitate the build process. `make xdp` compiles the XDP program and generates the Go source to reference the resulting bytecode. Once XDP ELF object is produced, you can build the Go binary with `make go`. After build is done, the binary will be availalbe in `cmd/oxdpus/oxdpus`.

## Usage

To see available CLI options, run `oxdpus --help`:

```
oxdpus --help
A toy tool that leverages the super powers of XDP to bring in-kernel IP filtering

Usage:
  oxdpus [command]

Available Commands:
  add         Appends a new IP address to the blacklist
  attach      Attaches the XDP program on the specified device
  detach      Removes the XDP program from the specified device
  help        Help about any command
  list        Shows all IP addresses registered in the blacklist
  remove      Removes an IP address from the blacklist

Flags:
  -h, --help   help for oxdpus

Use "oxdpus [command] --help" for more information about a command.
```
