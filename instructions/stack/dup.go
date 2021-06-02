package stack

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

/* Duplicate the top value of operand stack */

type DUP struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
             \_
               |
               V
[...][c][b][a][a]
*/

func (inst *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

// Duplicate the top operand stack value and insert two values down

type DUP_X1 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
          __/
         |
         V
[...][c][a][b][a]
*/

func (inst *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	top := stack.PopSlot()
	blow := stack.PopSlot()
	stack.PushSlot(top)
	stack.PushSlot(blow)
	stack.PushSlot(top)
}

// Duplicate the top operand stack value and insert three values down

type DUP_X2 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
       _____/
      |
      V
[...][a][c][b][a]
*/

func (inst *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	top := stack.PopSlot()
	blow := stack.PopSlot()
	third := stack.PopSlot()
	stack.PushSlot(top)
	stack.PushSlot(third)
	stack.PushSlot(blow)
	stack.PushSlot(top)
}

// Duplicate the top one and two operand stack values

type DUP2 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]____
          \____   |
               |  |
               V  V
[...][c][b][a][b][a]
*/

func (inst *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	top := stack.PopSlot()
	blow := stack.PopSlot()
	stack.PushSlot(blow)
	stack.PushSlot(top)
	stack.PushSlot(blow)
	stack.PushSlot(top)
}

// Duplicate the top one and two operand stack values and insert two or three values down

type DUP2_X1 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
       _/ __/
      |  |
      V  V
[...][b][a][c][b][a]
*/

func (inst *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	top := stack.PopSlot()
	blow := stack.PopSlot()
	third := stack.PopSlot()
	stack.PushSlot(blow)
	stack.PushSlot(top)
	stack.PushSlot(third)
	stack.PushSlot(blow)
	stack.PushSlot(top)
}

// Duplicate the top one or two operand stack values and insert two, three, or four values down

type DUP2_X2 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][d][c][b][a]
       ____/ __/
      |   __/
      V  V
[...][b][a][d][c][b][a]
*/

func (inst *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	top := stack.PopSlot()
	blow := stack.PopSlot()
	third := stack.PopSlot()
	fourth := stack.PopSlot()
	stack.PushSlot(blow)
	stack.PushSlot(top)
	stack.PushSlot(fourth)
	stack.PushSlot(third)
	stack.PushSlot(blow)
	stack.PushSlot(top)
}
