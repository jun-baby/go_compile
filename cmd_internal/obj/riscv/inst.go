// Code generated by parse.py -go rv64_a rv64_d rv64_f rv64_i rv64_m rv64_q rv_a rv_d rv_f rv_i rv_m rv_q rv_s rv_system rv_zicsr; DO NOT EDIT.
package riscv

import "compile/cmd_internal/obj"

type inst struct {
	opcode uint32
	funct3 uint32
	rs2    uint32
	csr    int64
	funct7 uint32
}

func encode(a obj.As) *inst {
	switch a {
	case AADD:
		return &inst{0x33, 0x0, 0x0, 0, 0x0}
	case AADDI:
		return &inst{0x13, 0x0, 0x0, 0, 0x0}
	case AADDIW:
		return &inst{0x1b, 0x0, 0x0, 0, 0x0}
	case AADDW:
		return &inst{0x3b, 0x0, 0x0, 0, 0x0}
	case AAMOADDD:
		return &inst{0x2f, 0x3, 0x0, 0, 0x0}
	case AAMOADDW:
		return &inst{0x2f, 0x2, 0x0, 0, 0x0}
	case AAMOANDD:
		return &inst{0x2f, 0x3, 0x0, 1536, 0x30}
	case AAMOANDW:
		return &inst{0x2f, 0x2, 0x0, 1536, 0x30}
	case AAMOMAXD:
		return &inst{0x2f, 0x3, 0x0, -1536, 0x50}
	case AAMOMAXW:
		return &inst{0x2f, 0x2, 0x0, -1536, 0x50}
	case AAMOMAXUD:
		return &inst{0x2f, 0x3, 0x0, -512, 0x70}
	case AAMOMAXUW:
		return &inst{0x2f, 0x2, 0x0, -512, 0x70}
	case AAMOMIND:
		return &inst{0x2f, 0x3, 0x0, -2048, 0x40}
	case AAMOMINW:
		return &inst{0x2f, 0x2, 0x0, -2048, 0x40}
	case AAMOMINUD:
		return &inst{0x2f, 0x3, 0x0, -1024, 0x60}
	case AAMOMINUW:
		return &inst{0x2f, 0x2, 0x0, -1024, 0x60}
	case AAMOORD:
		return &inst{0x2f, 0x3, 0x0, 1024, 0x20}
	case AAMOORW:
		return &inst{0x2f, 0x2, 0x0, 1024, 0x20}
	case AAMOSWAPD:
		return &inst{0x2f, 0x3, 0x0, 128, 0x4}
	case AAMOSWAPW:
		return &inst{0x2f, 0x2, 0x0, 128, 0x4}
	case AAMOXORD:
		return &inst{0x2f, 0x3, 0x0, 512, 0x10}
	case AAMOXORW:
		return &inst{0x2f, 0x2, 0x0, 512, 0x10}
	case AAND:
		return &inst{0x33, 0x7, 0x0, 0, 0x0}
	case AANDI:
		return &inst{0x13, 0x7, 0x0, 0, 0x0}
	case AAUIPC:
		return &inst{0x17, 0x0, 0x0, 0, 0x0}
	case ABEQ:
		return &inst{0x63, 0x0, 0x0, 0, 0x0}
	case ABGE:
		return &inst{0x63, 0x5, 0x0, 0, 0x0}
	case ABGEU:
		return &inst{0x63, 0x7, 0x0, 0, 0x0}
	case ABLT:
		return &inst{0x63, 0x4, 0x0, 0, 0x0}
	case ABLTU:
		return &inst{0x63, 0x6, 0x0, 0, 0x0}
	case ABNE:
		return &inst{0x63, 0x1, 0x0, 0, 0x0}
	case ACSRRC:
		return &inst{0x73, 0x3, 0x0, 0, 0x0}
	case ACSRRCI:
		return &inst{0x73, 0x7, 0x0, 0, 0x0}
	case ACSRRS:
		return &inst{0x73, 0x2, 0x0, 0, 0x0}
	case ACSRRSI:
		return &inst{0x73, 0x6, 0x0, 0, 0x0}
	case ACSRRW:
		return &inst{0x73, 0x1, 0x0, 0, 0x0}
	case ACSRRWI:
		return &inst{0x73, 0x5, 0x0, 0, 0x0}
	case ADIV:
		return &inst{0x33, 0x4, 0x0, 32, 0x1}
	case ADIVU:
		return &inst{0x33, 0x5, 0x0, 32, 0x1}
	case ADIVUW:
		return &inst{0x3b, 0x5, 0x0, 32, 0x1}
	case ADIVW:
		return &inst{0x3b, 0x4, 0x0, 32, 0x1}
	case ADRET:
		return &inst{0x73, 0x0, 0x12, 1970, 0x3d}
	case AEBREAK:
		return &inst{0x73, 0x0, 0x1, 1, 0x0}
	case AECALL:
		return &inst{0x73, 0x0, 0x0, 0, 0x0}
	case AFADDD:
		return &inst{0x53, 0x0, 0x0, 32, 0x1}
	case AFADDQ:
		return &inst{0x53, 0x0, 0x0, 96, 0x3}
	case AFADDS:
		return &inst{0x53, 0x0, 0x0, 0, 0x0}
	case AFCLASSD:
		return &inst{0x53, 0x1, 0x0, -480, 0x71}
	case AFCLASSQ:
		return &inst{0x53, 0x1, 0x0, -416, 0x73}
	case AFCLASSS:
		return &inst{0x53, 0x1, 0x0, -512, 0x70}
	case AFCVTDL:
		return &inst{0x53, 0x0, 0x2, -734, 0x69}
	case AFCVTDLU:
		return &inst{0x53, 0x0, 0x3, -733, 0x69}
	case AFCVTDQ:
		return &inst{0x53, 0x0, 0x3, 1059, 0x21}
	case AFCVTDS:
		return &inst{0x53, 0x0, 0x0, 1056, 0x21}
	case AFCVTDW:
		return &inst{0x53, 0x0, 0x0, -736, 0x69}
	case AFCVTDWU:
		return &inst{0x53, 0x0, 0x1, -735, 0x69}
	case AFCVTLD:
		return &inst{0x53, 0x0, 0x2, -990, 0x61}
	case AFCVTLQ:
		return &inst{0x53, 0x0, 0x2, -926, 0x63}
	case AFCVTLS:
		return &inst{0x53, 0x0, 0x2, -1022, 0x60}
	case AFCVTLUD:
		return &inst{0x53, 0x0, 0x3, -989, 0x61}
	case AFCVTLUQ:
		return &inst{0x53, 0x0, 0x3, -925, 0x63}
	case AFCVTLUS:
		return &inst{0x53, 0x0, 0x3, -1021, 0x60}
	case AFCVTQD:
		return &inst{0x53, 0x0, 0x1, 1121, 0x23}
	case AFCVTQL:
		return &inst{0x53, 0x0, 0x2, -670, 0x6b}
	case AFCVTQLU:
		return &inst{0x53, 0x0, 0x3, -669, 0x6b}
	case AFCVTQS:
		return &inst{0x53, 0x0, 0x0, 1120, 0x23}
	case AFCVTQW:
		return &inst{0x53, 0x0, 0x0, -672, 0x6b}
	case AFCVTQWU:
		return &inst{0x53, 0x0, 0x1, -671, 0x6b}
	case AFCVTSD:
		return &inst{0x53, 0x0, 0x1, 1025, 0x20}
	case AFCVTSL:
		return &inst{0x53, 0x0, 0x2, -766, 0x68}
	case AFCVTSLU:
		return &inst{0x53, 0x0, 0x3, -765, 0x68}
	case AFCVTSQ:
		return &inst{0x53, 0x0, 0x3, 1027, 0x20}
	case AFCVTSW:
		return &inst{0x53, 0x0, 0x0, -768, 0x68}
	case AFCVTSWU:
		return &inst{0x53, 0x0, 0x1, -767, 0x68}
	case AFCVTWD:
		return &inst{0x53, 0x0, 0x0, -992, 0x61}
	case AFCVTWQ:
		return &inst{0x53, 0x0, 0x0, -928, 0x63}
	case AFCVTWS:
		return &inst{0x53, 0x0, 0x0, -1024, 0x60}
	case AFCVTWUD:
		return &inst{0x53, 0x0, 0x1, -991, 0x61}
	case AFCVTWUQ:
		return &inst{0x53, 0x0, 0x1, -927, 0x63}
	case AFCVTWUS:
		return &inst{0x53, 0x0, 0x1, -1023, 0x60}
	case AFDIVD:
		return &inst{0x53, 0x0, 0x0, 416, 0xd}
	case AFDIVQ:
		return &inst{0x53, 0x0, 0x0, 480, 0xf}
	case AFDIVS:
		return &inst{0x53, 0x0, 0x0, 384, 0xc}
	case AFENCE:
		return &inst{0xf, 0x0, 0x0, 0, 0x0}
	case AFENCETSO:
		return &inst{0xf, 0x0, 0x13, -1997, 0x41}
	case AFEQD:
		return &inst{0x53, 0x2, 0x0, -1504, 0x51}
	case AFEQQ:
		return &inst{0x53, 0x2, 0x0, -1440, 0x53}
	case AFEQS:
		return &inst{0x53, 0x2, 0x0, -1536, 0x50}
	case AFLD:
		return &inst{0x7, 0x3, 0x0, 0, 0x0}
	case AFLED:
		return &inst{0x53, 0x0, 0x0, -1504, 0x51}
	case AFLEQ:
		return &inst{0x53, 0x0, 0x0, -1440, 0x53}
	case AFLES:
		return &inst{0x53, 0x0, 0x0, -1536, 0x50}
	case AFLQ:
		return &inst{0x7, 0x4, 0x0, 0, 0x0}
	case AFLTD:
		return &inst{0x53, 0x1, 0x0, -1504, 0x51}
	case AFLTQ:
		return &inst{0x53, 0x1, 0x0, -1440, 0x53}
	case AFLTS:
		return &inst{0x53, 0x1, 0x0, -1536, 0x50}
	case AFLW:
		return &inst{0x7, 0x2, 0x0, 0, 0x0}
	case AFMADDD:
		return &inst{0x43, 0x0, 0x0, 32, 0x1}
	case AFMADDQ:
		return &inst{0x43, 0x0, 0x0, 96, 0x3}
	case AFMADDS:
		return &inst{0x43, 0x0, 0x0, 0, 0x0}
	case AFMAXD:
		return &inst{0x53, 0x1, 0x0, 672, 0x15}
	case AFMAXQ:
		return &inst{0x53, 0x1, 0x0, 736, 0x17}
	case AFMAXS:
		return &inst{0x53, 0x1, 0x0, 640, 0x14}
	case AFMIND:
		return &inst{0x53, 0x0, 0x0, 672, 0x15}
	case AFMINQ:
		return &inst{0x53, 0x0, 0x0, 736, 0x17}
	case AFMINS:
		return &inst{0x53, 0x0, 0x0, 640, 0x14}
	case AFMSUBD:
		return &inst{0x47, 0x0, 0x0, 32, 0x1}
	case AFMSUBQ:
		return &inst{0x47, 0x0, 0x0, 96, 0x3}
	case AFMSUBS:
		return &inst{0x47, 0x0, 0x0, 0, 0x0}
	case AFMULD:
		return &inst{0x53, 0x0, 0x0, 288, 0x9}
	case AFMULQ:
		return &inst{0x53, 0x0, 0x0, 352, 0xb}
	case AFMULS:
		return &inst{0x53, 0x0, 0x0, 256, 0x8}
	case AFMVDX:
		return &inst{0x53, 0x0, 0x0, -224, 0x79}
	case AFMVSX:
		return &inst{0x53, 0x0, 0x0, -256, 0x78}
	case AFMVWX:
		return &inst{0x53, 0x0, 0x0, -256, 0x78}
	case AFMVXD:
		return &inst{0x53, 0x0, 0x0, -480, 0x71}
	case AFMVXS:
		return &inst{0x53, 0x0, 0x0, -512, 0x70}
	case AFMVXW:
		return &inst{0x53, 0x0, 0x0, -512, 0x70}
	case AFNMADDD:
		return &inst{0x4f, 0x0, 0x0, 32, 0x1}
	case AFNMADDQ:
		return &inst{0x4f, 0x0, 0x0, 96, 0x3}
	case AFNMADDS:
		return &inst{0x4f, 0x0, 0x0, 0, 0x0}
	case AFNMSUBD:
		return &inst{0x4b, 0x0, 0x0, 32, 0x1}
	case AFNMSUBQ:
		return &inst{0x4b, 0x0, 0x0, 96, 0x3}
	case AFNMSUBS:
		return &inst{0x4b, 0x0, 0x0, 0, 0x0}
	case AFRCSR:
		return &inst{0x73, 0x2, 0x3, 3, 0x0}
	case AFRFLAGS:
		return &inst{0x73, 0x2, 0x1, 1, 0x0}
	case AFRRM:
		return &inst{0x73, 0x2, 0x2, 2, 0x0}
	case AFSCSR:
		return &inst{0x73, 0x1, 0x3, 3, 0x0}
	case AFSD:
		return &inst{0x27, 0x3, 0x0, 0, 0x0}
	case AFSFLAGS:
		return &inst{0x73, 0x1, 0x1, 1, 0x0}
	case AFSFLAGSI:
		return &inst{0x73, 0x5, 0x1, 1, 0x0}
	case AFSGNJD:
		return &inst{0x53, 0x0, 0x0, 544, 0x11}
	case AFSGNJQ:
		return &inst{0x53, 0x0, 0x0, 608, 0x13}
	case AFSGNJS:
		return &inst{0x53, 0x0, 0x0, 512, 0x10}
	case AFSGNJND:
		return &inst{0x53, 0x1, 0x0, 544, 0x11}
	case AFSGNJNQ:
		return &inst{0x53, 0x1, 0x0, 608, 0x13}
	case AFSGNJNS:
		return &inst{0x53, 0x1, 0x0, 512, 0x10}
	case AFSGNJXD:
		return &inst{0x53, 0x2, 0x0, 544, 0x11}
	case AFSGNJXQ:
		return &inst{0x53, 0x2, 0x0, 608, 0x13}
	case AFSGNJXS:
		return &inst{0x53, 0x2, 0x0, 512, 0x10}
	case AFSQ:
		return &inst{0x27, 0x4, 0x0, 0, 0x0}
	case AFSQRTD:
		return &inst{0x53, 0x0, 0x0, 1440, 0x2d}
	case AFSQRTQ:
		return &inst{0x53, 0x0, 0x0, 1504, 0x2f}
	case AFSQRTS:
		return &inst{0x53, 0x0, 0x0, 1408, 0x2c}
	case AFSRM:
		return &inst{0x73, 0x1, 0x2, 2, 0x0}
	case AFSRMI:
		return &inst{0x73, 0x5, 0x2, 2, 0x0}
	case AFSUBD:
		return &inst{0x53, 0x0, 0x0, 160, 0x5}
	case AFSUBQ:
		return &inst{0x53, 0x0, 0x0, 224, 0x7}
	case AFSUBS:
		return &inst{0x53, 0x0, 0x0, 128, 0x4}
	case AFSW:
		return &inst{0x27, 0x2, 0x0, 0, 0x0}
	case AJAL:
		return &inst{0x6f, 0x0, 0x0, 0, 0x0}
	case AJALR:
		return &inst{0x67, 0x0, 0x0, 0, 0x0}
	case ALB:
		return &inst{0x3, 0x0, 0x0, 0, 0x0}
	case ALBU:
		return &inst{0x3, 0x4, 0x0, 0, 0x0}
	case ALD:
		return &inst{0x3, 0x3, 0x0, 0, 0x0}
	case ALH:
		return &inst{0x3, 0x1, 0x0, 0, 0x0}
	case ALHU:
		return &inst{0x3, 0x5, 0x0, 0, 0x0}
	case ALRD:
		return &inst{0x2f, 0x3, 0x0, 256, 0x8}
	case ALRW:
		return &inst{0x2f, 0x2, 0x0, 256, 0x8}
	case ALUI:
		return &inst{0x37, 0x0, 0x0, 0, 0x0}
	case ALW:
		return &inst{0x3, 0x2, 0x0, 0, 0x0}
	case ALWU:
		return &inst{0x3, 0x6, 0x0, 0, 0x0}
	case AMRET:
		return &inst{0x73, 0x0, 0x2, 770, 0x18}
	case AMUL:
		return &inst{0x33, 0x0, 0x0, 32, 0x1}
	case AMULH:
		return &inst{0x33, 0x1, 0x0, 32, 0x1}
	case AMULHSU:
		return &inst{0x33, 0x2, 0x0, 32, 0x1}
	case AMULHU:
		return &inst{0x33, 0x3, 0x0, 32, 0x1}
	case AMULW:
		return &inst{0x3b, 0x0, 0x0, 32, 0x1}
	case AOR:
		return &inst{0x33, 0x6, 0x0, 0, 0x0}
	case AORI:
		return &inst{0x13, 0x6, 0x0, 0, 0x0}
	case APAUSE:
		return &inst{0xf, 0x0, 0x10, 16, 0x0}
	case ARDCYCLE:
		return &inst{0x73, 0x2, 0x0, -1024, 0x60}
	case ARDCYCLEH:
		return &inst{0x73, 0x2, 0x0, -896, 0x64}
	case ARDINSTRET:
		return &inst{0x73, 0x2, 0x2, -1022, 0x60}
	case ARDINSTRETH:
		return &inst{0x73, 0x2, 0x2, -894, 0x64}
	case ARDTIME:
		return &inst{0x73, 0x2, 0x1, -1023, 0x60}
	case ARDTIMEH:
		return &inst{0x73, 0x2, 0x1, -895, 0x64}
	case AREM:
		return &inst{0x33, 0x6, 0x0, 32, 0x1}
	case AREMU:
		return &inst{0x33, 0x7, 0x0, 32, 0x1}
	case AREMUW:
		return &inst{0x3b, 0x7, 0x0, 32, 0x1}
	case AREMW:
		return &inst{0x3b, 0x6, 0x0, 32, 0x1}
	case ASB:
		return &inst{0x23, 0x0, 0x0, 0, 0x0}
	case ASBREAK:
		return &inst{0x73, 0x0, 0x1, 1, 0x0}
	case ASCD:
		return &inst{0x2f, 0x3, 0x0, 384, 0xc}
	case ASCW:
		return &inst{0x2f, 0x2, 0x0, 384, 0xc}
	case ASCALL:
		return &inst{0x73, 0x0, 0x0, 0, 0x0}
	case ASD:
		return &inst{0x23, 0x3, 0x0, 0, 0x0}
	case ASFENCEVMA:
		return &inst{0x73, 0x0, 0x0, 288, 0x9}
	case ASH:
		return &inst{0x23, 0x1, 0x0, 0, 0x0}
	case ASLL:
		return &inst{0x33, 0x1, 0x0, 0, 0x0}
	case ASLLI:
		return &inst{0x13, 0x1, 0x0, 0, 0x0}
	case ASLLIW:
		return &inst{0x1b, 0x1, 0x0, 0, 0x0}
	case ASLLW:
		return &inst{0x3b, 0x1, 0x0, 0, 0x0}
	case ASLT:
		return &inst{0x33, 0x2, 0x0, 0, 0x0}
	case ASLTI:
		return &inst{0x13, 0x2, 0x0, 0, 0x0}
	case ASLTIU:
		return &inst{0x13, 0x3, 0x0, 0, 0x0}
	case ASLTU:
		return &inst{0x33, 0x3, 0x0, 0, 0x0}
	case ASRA:
		return &inst{0x33, 0x5, 0x0, 1024, 0x20}
	case ASRAI:
		return &inst{0x13, 0x5, 0x0, 1024, 0x20}
	case ASRAIW:
		return &inst{0x1b, 0x5, 0x0, 1024, 0x20}
	case ASRAW:
		return &inst{0x3b, 0x5, 0x0, 1024, 0x20}
	case ASRET:
		return &inst{0x73, 0x0, 0x2, 258, 0x8}
	case ASRL:
		return &inst{0x33, 0x5, 0x0, 0, 0x0}
	case ASRLI:
		return &inst{0x13, 0x5, 0x0, 0, 0x0}
	case ASRLIW:
		return &inst{0x1b, 0x5, 0x0, 0, 0x0}
	case ASRLW:
		return &inst{0x3b, 0x5, 0x0, 0, 0x0}
	case ASUB:
		return &inst{0x33, 0x0, 0x0, 1024, 0x20}
	case ASUBW:
		return &inst{0x3b, 0x0, 0x0, 1024, 0x20}
	case ASW:
		return &inst{0x23, 0x2, 0x0, 0, 0x0}
	case AWFI:
		return &inst{0x73, 0x0, 0x5, 261, 0x8}
	case AXOR:
		return &inst{0x33, 0x4, 0x0, 0, 0x0}
	case AXORI:
		return &inst{0x13, 0x4, 0x0, 0, 0x0}
	}
	return nil
}
