package references

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

// hack!

func (inst *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}
