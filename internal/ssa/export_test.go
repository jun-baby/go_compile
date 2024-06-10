// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"testing"

	"compile/cmd_internal/obj"
	"compile/cmd_internal/obj/arm64"
	"compile/cmd_internal/obj/s390x"
	"compile/cmd_internal/obj/x86"
	"compile/cmd_internal/src"
	"compile/internal/ir"
	"compile/internal/typecheck"
	"compile/internal/types"
)

var CheckFunc = checkFunc
var Opt = opt
var Deadcode = deadcode
var Copyelim = copyelim

var testCtxts = map[string]*obj.Link{
	"amd64": obj.Linknew(&x86.Linkamd64),
	"s390x": obj.Linknew(&s390x.Links390x),
	"arm64": obj.Linknew(&arm64.Linkarm64),
}

func testConfig(tb testing.TB) *Conf      { return testConfigArch(tb, "amd64") }
func testConfigS390X(tb testing.TB) *Conf { return testConfigArch(tb, "s390x") }
func testConfigARM64(tb testing.TB) *Conf { return testConfigArch(tb, "arm64") }

func testConfigArch(tb testing.TB, arch string) *Conf {
	ctxt, ok := testCtxts[arch]
	if !ok {
		tb.Fatalf("unknown arch %s", arch)
	}
	if ctxt.Arch.PtrSize != 8 {
		tb.Fatal("testTypes is 64-bit only")
	}
	c := &Conf{
		config: NewConfig(arch, testTypes, ctxt, true, false),
		tb:     tb,
	}
	return c
}

type Conf struct {
	config *Config
	tb     testing.TB
	fe     Frontend
}

func (c *Conf) Frontend() Frontend {
	if c.fe == nil {
		pkg := types.NewPkg("my/import/path", "path")
		fn := ir.NewFunc(src.NoXPos, src.NoXPos, pkg.Lookup("function"), types.NewSignature(nil, nil, nil))
		fn.DeclareParams(true)
		fn.LSym = &obj.LSym{Name: "my/import/path.function"}

		c.fe = TestFrontend{
			t:    c.tb,
			ctxt: c.config.ctxt,
			f:    fn,
		}
	}
	return c.fe
}

func (c *Conf) Temp(typ *types.Type) *ir.Name {
	n := ir.NewNameAt(src.NoXPos, &types.Sym{Name: "aFakeAuto"}, typ)
	n.Class = ir.PAUTO
	return n
}

// TestFrontend is a test-only frontend.
// It assumes 64 bit integers and pointers.
type TestFrontend struct {
	t    testing.TB
	ctxt *obj.Link
	f    *ir.Func
}

func (TestFrontend) StringData(s string) *obj.LSym {
	return nil
}
func (d TestFrontend) SplitSlot(parent *LocalSlot, suffix string, offset int64, t *types.Type) LocalSlot {
	return LocalSlot{N: parent.N, Type: t, Off: offset}
}
func (d TestFrontend) Syslook(s string) *obj.LSym {
	return d.ctxt.Lookup(s)
}
func (TestFrontend) UseWriteBarrier() bool {
	return true // only writebarrier_test cares
}

func (d TestFrontend) Logf(msg string, args ...interface{}) { d.t.Logf(msg, args...) }
func (d TestFrontend) Log() bool                            { return true }

func (d TestFrontend) Fatalf(_ src.XPos, msg string, args ...interface{}) { d.t.Fatalf(msg, args...) }
func (d TestFrontend) Warnl(_ src.XPos, msg string, args ...interface{})  { d.t.Logf(msg, args...) }
func (d TestFrontend) Debug_checknil() bool                               { return false }

func (d TestFrontend) Func() *ir.Func {
	return d.f
}

var testTypes Types

func init() {
	// TODO(mdempsky): Push into types.InitUniverse or typecheck.InitUniverse.
	types.PtrSize = 8
	types.RegSize = 8
	types.MaxWidth = 1 << 50

	typecheck.InitUniverse()
	testTypes.SetTypPtrs()
}