package stores

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

/* Store reference into local variable */

type DSTORE struct {
	base.Index8Instruction
}

func (inst *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame, inst.Index)
}

type DSTORE_0 struct {
	base.NoOperandsInstruction
}

func (inst *DSTORE_0) Execute(frame *rtda.Frame) {
	_dstore(frame, 0)
}

type DSTORE_1 struct {
	base.NoOperandsInstruction
}

func (inst *DSTORE_1) Execute(frame *rtda.Frame) {
	_dstore(frame, 1)
}

type DSTORE_2 struct {
	base.NoOperandsInstruction
}

func (inst *DSTORE_2) Execute(frame *rtda.Frame) {
	_dstore(frame, 2)
}

type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func (inst *DSTORE_3) Execute(frame *rtda.Frame) {
	_dstore(frame, 3)
}

func _dstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}
