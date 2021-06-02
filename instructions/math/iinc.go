package math

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

// Increment local variable by constant value

type IINC struct {
	Index uint
	Const int32
}

func (inst *IINC) FetchOperands(reader *base.BytecodeReader) {
	inst.Index = uint(reader.ReadUint8())
	inst.Const = int32(reader.ReadInt8())
}

func (inst *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(inst.Index)
	val += inst.Const
	localVars.SetInt(inst.Index, val)
}
