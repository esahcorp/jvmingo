package extended

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

type IFNULL struct{ base.BranchInstruction }

func (inst *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, inst.Offset)
	}
}

type IFNONNULL struct{ base.BranchInstruction }

func (inst *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, inst.Offset)
	}
}
