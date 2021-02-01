package cpu

type Instruction struct {
	Name    string
	Opcode  uint8
	Execute func(*CPU)
	Cycles  uint8
}

var (
	Instructions = make(map[byte]Instruction)
)

func addInstruction(ins Instruction) {
	Instructions[ins.Opcode] = ins
}

func LoadInstructionSet() {
	addInstruction(Instruction{
		Name:   "LDA_Immediate",
		Opcode: 0xA9,
		Cycles: 2,
		Execute: func(cpu *CPU) {
			v := cpu.FetchByte(cpu.PC)

			cpu.SR.SetZ(v == 0)
			cpu.SR.SetN(v&(1<<7) > 0)

			cpu.AC = v
		},
	})

	addInstruction(Instruction{
		Name:   "LDX_Immediate",
		Opcode: 0xA2,
		Cycles: 2,
		Execute: func(cpu *CPU) {
			v := cpu.FetchByte(cpu.PC)

			cpu.SR.SetZ(v == 0)
			cpu.SR.SetN(v&(1<<7) > 0)

			cpu.XR = v
		},
	})

	addInstruction(Instruction{
		Name:   "LDY_Immediate",
		Opcode: 0xA0,
		Cycles: 2,
		Execute: func(cpu *CPU) {
			v := cpu.FetchByte(cpu.PC)

			cpu.SR.SetZ(v == 0)
			cpu.SR.SetN(v&(1<<7) > 0)

			cpu.YR = v
		},
	})
	addInstruction(Instruction{
		Name:   "NOP",
		Opcode: 0xEA,
		Cycles: 1,
		Execute: func(cpu *CPU) {
		},
	})
	addInstruction(Instruction{
		Name:   "DEC_Absolute",
		Opcode: 0xCE,
		Cycles: 6,
		Execute: func(cpu *CPU) {
			addr := cpu.FetchWord(cpu.PC)

			cpu.Mem[addr]--

			cpu.SR.SetZ(cpu.Mem[addr] == 0)
			cpu.SR.SetN(cpu.Mem[addr]&(1<<7) > 0)
		},
	})
	addInstruction(Instruction{
		Name:   "INC_Absolute",
		Opcode: 0xEE,
		Cycles: 6,
		Execute: func(cpu *CPU) {
			addr := cpu.FetchWord(cpu.PC)

			cpu.Mem[addr]++

			cpu.SR.SetZ(cpu.Mem[addr] == 0)
			cpu.SR.SetN(cpu.Mem[addr]&(1<<7) > 0)
		},
	})
	addInstruction(Instruction{
		Name:   "STA_Absolute",
		Opcode: 0x8D,
		Cycles: 4,
		Execute: func(cpu *CPU) {
			addr := cpu.FetchWord(cpu.PC)

			cpu.Mem[addr] -= cpu.AC
		},
	})
	addInstruction(Instruction{
		Name:   "STX_Absolute",
		Opcode: 0x8E,
		Cycles: 4,
		Execute: func(cpu *CPU) {
			addr := cpu.FetchWord(cpu.PC)

			cpu.Mem[addr] -= cpu.XR
		},
	})
	addInstruction(Instruction{
		Name:   "STY_Absolute",
		Opcode: 0x8C,
		Cycles: 4,
		Execute: func(cpu *CPU) {
			addr := cpu.FetchWord(cpu.PC)

			cpu.Mem[addr] -= cpu.YR
		},
	})
	// addInstruction(Instruction)
}
