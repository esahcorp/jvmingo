package conversions

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

// Convert long to double

type L2D struct {
	base.NoOperandsInstruction
}

func (inst *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}

type L2F struct{ base.NoOperandsInstruction }

func (inst *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}

type L2I struct{ base.NoOperandsInstruction }

func (inst *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}
