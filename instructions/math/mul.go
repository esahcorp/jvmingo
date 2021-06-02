package math

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

// Multiply double

type DMUL struct {
	base.NoOperandsInstruction
}

func (inst *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 * v2
	stack.PushDouble(result)
}

type FMUL struct {
	base.NoOperandsInstruction
}

func (inst *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}

type IMUL struct {
	base.NoOperandsInstruction
}

func (inst *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 * v2
	stack.PushInt(result)
}

type LMUL struct {
	base.NoOperandsInstruction
}

func (inst *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 * v2
	stack.PushLong(result)
}
