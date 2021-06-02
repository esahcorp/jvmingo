package constants

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

/* No Operands Constant Instruction */

type NOP struct {
	base.NoOperandsInstruction
}

func (N *NOP) Execute(frame *rtda.Frame) {
	// nothing to do
}
