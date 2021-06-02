package stores

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

/* Store long into local variable */

type LSTORE struct {
	base.Index8Instruction
}

func (inst *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, inst.Index)
}

type LSTORE_0 struct {
	base.NoOperandsInstruction
}

func (inst *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}

type LSTORE_1 struct {
	base.NoOperandsInstruction
}

func (inst *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}

type LSTORE_2 struct {
	base.NoOperandsInstruction
}

func (inst *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}

type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func (inst *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}
