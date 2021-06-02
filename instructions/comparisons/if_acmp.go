package comparisons

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

// Branch if reference comparison succeeds

type IF_ACMPEQ struct{ base.BranchInstruction }

func (inst *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	if _acmp(frame) {
		base.Branch(frame, inst.Offset)
	}
}

type IF_ACMPNE struct{ base.BranchInstruction }

func (inst *IF_ACMPNE) Execute(frame *rtda.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, inst.Offset)
	}
}

func _acmp(frame *rtda.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2
}
