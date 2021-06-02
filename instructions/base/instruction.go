package base

import "jvmingo/rtda"

/* Instruction interface of all */

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

/* Instruction without operand */

type NoOperandsInstruction struct{}

func (NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// Nothing to do
}

type BranchInstruction struct {
	Offset int
}

func (i *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	i.Offset = int(reader.ReadInt16())
}

/* Instruction which operand is in constant pool, such as store and load */

type Index8Instruction struct {
	Index uint // index of constant pool
}

func (i *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	i.Index = uint(reader.ReadUint8())
}

/* 16 bit Index of constant pool is from Operands Stack */

type Index16Instruction struct {
	Index uint
}

func (i *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	i.Index = uint(reader.ReadUint16())
}
