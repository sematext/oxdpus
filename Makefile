SHELL=bash -o pipefail -e
UNAME=$(shell uname -r)
LINUX_HEADERS ?= /lib/modules/$(UNAME)
CLANG ?= clang
LLC ?= llc
GO ?= go
GOFMT ?= gofmt

ALL_SRC := $(shell find . -name "*.go" | grep -v -e ".*/\..*" -e ".*/_.*")


CFLAGS = -I $(LINUX_HEADERS)/build/arch/x86/include \
	-I $(LINUX_HEADERS)/build/arch/x86/include/generated/uapi \
	-I $(LINUX_HEADERS)/build/arch/x86/include/generated \
	-I $(LINUX_HEADERS)/build/include \
	-I $(LINUX_HEADERS)/build/arch/x86/include/uapi \
	-I $(LINUX_HEADERS)/build/include/uapi \
	-include $(LINUX_HEADERS)/build/include/linux/kconfig.h \
	-I $(LINUX_HEADERS)/build/include/generated/uapi \
	-D__KERNEL__ -D__ASM_SYSREG_H \
	-Wunused \
	-Wall \
	-Wno-compare-distinct-pointer-types \
	-fno-stack-protector \
	-Wno-pointer-sign \
	-O2 -S -emit-llvm

XDP_PROG := pkg/xdp/prog/obj/xdp.o
pkg/xdp/prog/obj/xdp.o: pkg/xdp/prog/xdp.c
	$(CLANG) $(CFLAGS) -c $< -o - | $(LLC) -march=bpf -mcpu=$(CPU) -filetype=obj -o $@

.PHONY: go
go:
	$(GO) build -ldflags "-s -w" -o cmd/oxdpus/oxdpus cmd/oxdpus/main.go

.PHONY: xdp
xdp: $(XDP_PROG)
	go-bindata -pkg gen -prefix "pkg/xdp/prog/obj" -o "pkg/xdp/prog/gen/xdp.go" "pkg/xdp/prog/obj/"

.PHONY: clean
clean:
	rm -f pkg/xdp/prog/obj/xdp.o

.PHONY: fmt
fmt:
	$(GOFMT) -e -s -l -w $(ALL_SRC)
