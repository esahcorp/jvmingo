package loads

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

/* Load long from local variable */

type LLOAD struct {
	base.Index8Instruction
}

func (inst *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, inst.Index)
}

type LLOAD_0 struct {
	base.NoOperandsInstruction
}

func (inst *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

type LLOAD_1 struct {
	base.NoOperandsInstruction
}

func (inst *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

type LLOAD_2 struct {
	base.NoOperandsInstruction
}

func (inst *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

type LLOAD_3 struct {
	base.NoOperandsInstruction
}

func (inst *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}

func _lload(frame *rtda.Frame, index uint) {
	long := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(long)
}
