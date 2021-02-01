package cpu

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/pbaettig/go-6502/internal/util"
)

const (
	MemorySize  = 64 * 1024
	ResetVector = 0xFFFC
)

func init() {
	LoadInstructionSet()
}

type Memory [MemorySize]uint8

func (mem Memory) Zero() {
	for i := 0; i < MemorySize; i++ {
		mem[i] = 0
	}
}

func (mem Memory) Print(start, end uint16) {
	fmt.Printf(hex.Dump(mem[start:end]))
}

func (mem Memory) PrintLine(start uint16) {
	fmts := func(buf []byte) string {
		var sb strings.Builder
		for _, c := range buf {
			if c < 32 || c > 126 {
				sb.WriteString(".")
			} else {
				sb.WriteByte(c)
			}
		}
		return sb.String()
	}
	s := start
	fmt.Printf(
		"%04X: %02X %02X %02X %02X %02X %02X %02X %02X  %02X %02X %02X %02X %02X %02X %02X %02X |%s|\n",
		s, mem[s], mem[s+1], mem[s+2], mem[s+3], mem[s+4], mem[s+5], mem[s+6], mem[s+7], mem[s+8],
		mem[s+9], mem[s+10], mem[s+11], mem[s+12], mem[s+13], mem[s+14], mem[s+15], fmts(mem[s:s+15]))
}

func (mem Memory) GetWord(addr uint16) uint16 {
	hb := mem[addr]
	lb := mem[addr+1]

	return util.BytesToWord(lb, hb)
}

type StatusRegister uint8

func (s *StatusRegister) N() bool {
	return (uint8(*s) & (1 << 7)) > 0
}

func (s *StatusRegister) SetN(v bool) {
	*s |= StatusRegister((util.Btoi(v) << 7))
}

func (s *StatusRegister) V() bool {
	return (uint8(*s) & (1 << 6)) > 0
}

func (s *StatusRegister) SetV(v bool) {
	*s |= StatusRegister((util.Btoi(v) << 6))
}

func (s *StatusRegister) B() bool {
	return (uint8(*s) & (1 << 4)) > 0
}

func (s *StatusRegister) SetB(v bool) {
	*s |= StatusRegister((util.Btoi(v) << 4))
}

func (s *StatusRegister) D() bool {
	return (uint8(*s) & (1 << 3)) > 0
}

func (s *StatusRegister) SetD(v bool) {
	*s |= StatusRegister((util.Btoi(v) << 3))
}

func (s *StatusRegister) I() bool {
	return (uint8(*s) & (1 << 2)) > 0
}

func (s *StatusRegister) SetI(v bool) {
	*s |= StatusRegister((util.Btoi(v) << 2))
}

func (s *StatusRegister) Z() bool {
	return (uint8(*s) & (1 << 1)) > 0
}

func (s *StatusRegister) SetZ(v bool) {
	*s |= StatusRegister((util.Btoi(v) << 1))
}

func (s *StatusRegister) C() bool {
	return (uint8(*s) & 0x01) > 0
}

func (s *StatusRegister) SetC(v bool) {
	*s |= StatusRegister(0x01)
}

type CPU struct {
	halted bool

	CurrentInstruction *Instruction

	Cycles uint64

	Mem Memory

	PC uint16
	AC uint8
	XR uint8
	YR uint8
	SR *StatusRegister
	SP uint16
}

func (cpu *CPU) Reset() {
	cpu.SR = new(StatusRegister)

	cpu.PC = cpu.Mem.GetWord(ResetVector)

}

func (cpu *CPU) FetchByte(addr uint16) uint8 {
	defer func() { cpu.PC++; cpu.Cycles++ }()

	return cpu.Mem[addr]
}

func (cpu *CPU) FetchWord(addr uint16) uint16 {
	defer func() { cpu.PC += 2; cpu.Cycles += 2 }()

	return cpu.Mem.GetWord(addr)
}

func Decode(opcode uint8) (Instruction, bool) {
	ins, ok := Instructions[opcode]

	return ins, ok
}

func (cpu *CPU) FetchDecodeExecute() {
	// fetch
	opcode := cpu.FetchByte(cpu.PC)

	//decode
	ins, ok := Decode(opcode)
	if !ok {
		cpu.CurrentInstruction = nil
		return
	}
	cpu.CurrentInstruction = &ins
	// decrement Instruction cycles to accomodate the Fetch
	cpu.CurrentInstruction.Cycles--

	//execute
	fmt.Printf("Executing %s (at 0x%04X) \n", cpu.CurrentInstruction.Name, cpu.PC)
	cpu.CurrentInstruction.Execute(cpu)

}

func (cpu *CPU) Halt() {
	cpu.halted = true
}

func (cpu *CPU) PrintRegisters() {
	fmt.Printf("Current Instruction: ")
	if cpu.CurrentInstruction != nil {
		fmt.Printf("%s (0x%02X)\n", cpu.CurrentInstruction.Name, cpu.CurrentInstruction.Opcode)
	} else {
		fmt.Printf("<nil> (0x%02X)\n", cpu.Mem[cpu.PC])
	}
	sr := *cpu.SR
	fmt.Printf("PC: 0x%04X 0b%08b\n", cpu.PC, cpu.PC)
	fmt.Printf("SR: 0x%02X   0b%08b\n", sr, sr)
	fmt.Printf("AC: 0x%02X   0b%08b (%d)\n", cpu.AC, cpu.AC, cpu.AC)
	fmt.Printf("XR: 0x%02X   0b%08b (%d)\n", cpu.XR, cpu.XR, cpu.XR)
	fmt.Printf("YR: 0x%02X   0b%08b (%d)\n", cpu.YR, cpu.YR, cpu.YR)

}

func New() *CPU {
	cpu := new(CPU)
	mem := Memory{}

	cpu.SR = new(StatusRegister)

	mem.Zero()
	cpu.Mem = mem

	return cpu
}
