package conversions

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

// Convert float to x

type F2D struct{ base.NoOperandsInstruction }

func (inst *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := float64(f)
	stack.PushDouble(d)
}

type F2I struct{ base.NoOperandsInstruction }

func (inst *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	i := int32(f)
	stack.PushInt(i)
}

type F2L struct{ base.NoOperandsInstruction }

func (inst *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	l := int64(f)
	stack.PushLong(l)
}
