# command-line-arguments
main.main STEXT nosplit size=14 args=0x0 locals=0x0 funcid=0x0 align=0x0
	0x0000 00000 (/home/miki/work/ardanlabs/353words/looping/asm.go:3)	TEXT	main.main(SB), NOSPLIT|NOFRAME|ABIInternal, $0-0
	0x0000 00000 (/home/miki/work/ardanlabs/353words/looping/asm.go:3)	FUNCDATA	$0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (/home/miki/work/ardanlabs/353words/looping/asm.go:3)	FUNCDATA	$1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (/home/miki/work/ardanlabs/353words/looping/asm.go:3)	XORL	AX, AX
	0x0002 00002 (/home/miki/work/ardanlabs/353words/looping/asm.go:4)	JMP	7
	0x0004 00004 (/home/miki/work/ardanlabs/353words/looping/asm.go:4)	INCQ	AX
	0x0007 00007 (/home/miki/work/ardanlabs/353words/looping/asm.go:4)	CMPQ	AX, $3
	0x000b 00011 (/home/miki/work/ardanlabs/353words/looping/asm.go:4)	JLT	4
	0x000d 00013 (/home/miki/work/ardanlabs/353words/looping/asm.go:6)	RET
	0x0000 31 c0 eb 03 48 ff c0 48 83 f8 03 7c f7 c3        1...H..H...|..
go:cuinfo.producer.main SDWARFCUINFO dupok size=0
	0x0000 72 65 67 61 62 69                                regabi
go:cuinfo.packagename.main SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
main..inittask SNOPTRDATA size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
gclocals·g2BeySu+wFnoycgXfElmcg== SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
