package math

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

// Negate double

type DNEG struct {
	base.NoOperandsInstruction
}

func (inst *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

type FNEG struct {
	base.NoOperandsInstruction
}

func (inst *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

type INEG struct {
	base.NoOperandsInstruction
}

func (inst *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

type LNEG struct {
	base.NoOperandsInstruction
}

func (inst *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
