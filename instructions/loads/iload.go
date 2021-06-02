package loads

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

/* Load int operand from local variable */

type ILOAD struct {
	base.Index8Instruction
}

func (inst *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, inst.Index)
}

type ILOAD_0 struct {
	base.NoOperandsInstruction
}

func (inst *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

type ILOAD_1 struct {
	base.NoOperandsInstruction
}

func (inst *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct {
	base.NoOperandsInstruction
}

func (inst *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func (inst *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
