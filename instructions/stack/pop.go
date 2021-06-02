package stack

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

/* Pop 32 bit type from Operand Stack */

type POP struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
            |
            V
[...][c][b]
*/

func (inst *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

/* Pop 64 bit type from Operand Stack */

type POP2 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
         |  |
         V  V
[...][c]
*/

func (inst *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
