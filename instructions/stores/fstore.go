package stores

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

/* Store reference into local variable */

type FSTORE struct {
	base.Index8Instruction
}

func (inst *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, inst.Index)
}

type FSTORE_0 struct {
	base.NoOperandsInstruction
}

func (inst *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct {
	base.NoOperandsInstruction
}

func (inst *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct {
	base.NoOperandsInstruction
}

func (inst *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func (inst *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}

func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}
