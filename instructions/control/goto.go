package control

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

type GOTO struct{ base.BranchInstruction }

func (inst *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, inst.Offset)
}
