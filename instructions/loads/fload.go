package loads

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

/* Load float from local variable */

type FLOAD struct {
	base.Index8Instruction
}

func (inst *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, inst.Index)
}

type FLOAD_0 struct {
	base.NoOperandsInstruction
}

func (inst *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

type FLOAD_1 struct {
	base.NoOperandsInstruction
}

func (inst *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

type FLOAD_2 struct {
	base.NoOperandsInstruction
}

func (inst *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func (inst *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}

func _fload(frame *rtda.Frame, index uint) {
	float := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(float)
}
