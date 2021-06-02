package math

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

// Boolean XOR int

type IXOR struct{ base.NoOperandsInstruction }

func (inst *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

type LXOR struct{ base.NoOperandsInstruction }

func (inst *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}
