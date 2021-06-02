package stack

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

// Swap the top two operand stack value order

type SWAP struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
          \/
          /\
         V  V
[...][c][a][b]
*/

func (inst *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	top := stack.PopSlot()
	blow := stack.PopSlot()
	stack.PushSlot(top)
	stack.PushSlot(blow)
}
