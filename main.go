// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"compile/internal/amd64"
	"compile/internal/arm"
	"compile/internal/arm64"
	"compile/internal/base"
	"compile/internal/gc"
	"compile/internal/loong64"
	"compile/internal/mips"
	"compile/internal/mips64"
	"compile/internal/ppc64"
	"compile/internal/riscv64"
	"compile/internal/s390x"
	"compile/internal/ssagen"
	"compile/internal/wasm"
	"compile/internal/x86"
	"compile/src_internal/buildcfg"
	"fmt"
	"log"
	"os"
)

var archInits = map[string]func(*ssagen.ArchInfo){
	"386":      x86.Init,
	"amd64":    amd64.Init,
	"arm":      arm.Init,
	"arm64":    arm64.Init,
	"loong64":  loong64.Init,
	"mips":     mips.Init,
	"mipsle":   mips.Init,
	"mips64":   mips64.Init,
	"mips64le": mips64.Init,
	"ppc64":    ppc64.Init,
	"ppc64le":  ppc64.Init,
	"riscv64":  riscv64.Init,
	"s390x":    s390x.Init,
	"wasm":     wasm.Init,
}

func main() {
	// disable timestamps for reproducible output
	log.SetFlags(0)
	log.SetPrefix("compile: ")

	buildcfg.Check()
	archInit, ok := archInits[buildcfg.GOARCH]
	if !ok {
		fmt.Fprintf(os.Stderr, "compile: unknown architecture %q\n", buildcfg.GOARCH)
		os.Exit(2)
	}

	gc.Main(archInit)
	base.Exit(0)
}
