package constants

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

/* Read byte or short operand then convert to int */

type BIPUSH struct {
	val int8 // Push byte
}

func (inst *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	inst.val = reader.ReadInt8()
}

func (inst *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(inst.val)
	frame.OperandStack().PushInt(i)
}

type SIPUSH struct {
	val int16 // Push short
}

func (inst *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	inst.val = reader.ReadInt16()
}

func (inst *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(inst.val)
	frame.OperandStack().PushInt(i)
}
