package comparisons

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

// Branch if int comparison with zero succeeds

type IFEQ struct{ base.BranchInstruction }

func (inst *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, inst.Offset)
	}
}

type IFNE struct{ base.BranchInstruction }

func (inst *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, inst.Offset)
	}
}

type IFLT struct{ base.BranchInstruction }

func (inst *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, inst.Offset)
	}
}

type IFLE struct{ base.BranchInstruction }

func (inst *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, inst.Offset)
	}
}

type IFGT struct{ base.BranchInstruction }

func (inst *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, inst.Offset)
	}
}

type IFGE struct{ base.BranchInstruction }

func (inst *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, inst.Offset)
	}
}
