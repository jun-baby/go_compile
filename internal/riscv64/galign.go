// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package riscv64

import (
	"compile/cmd_internal/obj/riscv"
	"compile/internal/ssagen"
)

func Init(arch *ssagen.ArchInfo) {
	arch.LinkArch = &riscv.LinkRISCV64

	arch.REGSP = riscv.REG_SP
	arch.MAXWIDTH = 1 << 50

	arch.Ginsnop = ginsnop
	arch.ZeroRange = zeroRange

	arch.SSAMarkMoves = ssaMarkMoves
	arch.SSAGenValue = ssaGenValue
	arch.SSAGenBlock = ssaGenBlock
	arch.LoadRegResult = loadRegResult
	arch.SpillArgReg = spillArgReg
}
