package main

import (
	"fmt"

	"github.com/pbaettig/go-6502/internal/cpu"
)

func main() {
	p := cpu.New()
	p.Mem[cpu.ResetVector] = 0x00
	p.Mem[cpu.ResetVector+1] = 0x00
	p.Mem[0x00] = 0xA9
	p.Mem[0x01] = 0xFF

	p.Mem[0x02] = 0xA2 // LDX 42
	p.Mem[0x03] = 0xFE //

	p.Mem[0x04] = 0xA0 // LDY 41
	p.Mem[0x05] = 0xFD

	p.Mem[0xABCD] = 0x43
	p.Mem[0xABCE] = 0x41
	p.Mem[0x06] = 0xCE // DEC abcd
	p.Mem[0x07] = 0xAB
	p.Mem[0x08] = 0xCD

	p.Mem[0x09] = 0xEE // INC abce
	p.Mem[0x0A] = 0xAB
	p.Mem[0x0B] = 0xCE

	p.Reset()

	p.PrintRegisters()
	fmt.Println()

	p.FetchDecodeExecute()

	p.PrintRegisters()
	fmt.Println()

	p.FetchDecodeExecute()

	p.PrintRegisters()
	fmt.Println()

	p.FetchDecodeExecute()

	p.PrintRegisters()
	fmt.Println()

	p.Mem.PrintLine(0xABCD)

	p.FetchDecodeExecute()
	p.PrintRegisters()
	fmt.Println()

	p.Mem.PrintLine(0xABCD)
	fmt.Println()

	p.FetchDecodeExecute()
	p.PrintRegisters()
	fmt.Println()

	p.Mem.PrintLine(0xABCD)
	fmt.Println()

}
