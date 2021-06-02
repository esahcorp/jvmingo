package math

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

// Subtract double

type DSUB struct {
	base.NoOperandsInstruction
}

func (inst *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 - v2
	stack.PushDouble(result)
}

type FSUB struct {
	base.NoOperandsInstruction
}

func (inst *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 - v2
	stack.PushFloat(result)
}

type ISUB struct {
	base.NoOperandsInstruction
}

func (inst *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 - v2
	stack.PushInt(result)
}

type LSUB struct {
	base.NoOperandsInstruction
}

func (inst *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 - v2
	stack.PushLong(result)
}
