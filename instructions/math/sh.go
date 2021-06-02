package math

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

// Shift left int

type ISHL struct {
	base.NoOperandsInstruction
}

func (inst *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	num := stack.PopInt()
	basic := stack.PopInt()
	s := uint32(num) & 0x1f // shift size is 4 bytes unsigned integer
	result := basic << s
	stack.PushInt(result)
}

// Shift right int

type ISHR struct {
	base.NoOperandsInstruction
}

func (inst *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	num := stack.PopInt()
	basic := stack.PopInt()
	s := uint32(num) & 0x1f // shift size is 4 bytes unsigned integer
	result := basic >> s
	stack.PushInt(result)
}

// Logical shift right int

type IUSHR struct {
	base.NoOperandsInstruction
}

func (inst *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	num := stack.PopInt()
	basic := stack.PopInt()
	s := uint32(num) & 0x1f
	result := int32(uint32(basic) >> s)
	stack.PushInt(result)
}

// Shift left long

type LSHL struct {
	base.NoOperandsInstruction
}

func (inst *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	num := stack.PopInt()
	basic := stack.PopLong()
	s := uint32(num) & 0x3f
	result := basic << s
	stack.PushLong(result)
}

type LSHR struct {
	base.NoOperandsInstruction
}

func (inst *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	num := stack.PopInt()
	basic := stack.PopLong()
	s := uint32(num) & 0x3f
	result := basic >> s
	stack.PushLong(result)
}

type LUSHR struct {
	base.NoOperandsInstruction
}

func (inst *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	num := stack.PopInt()
	basic := stack.PopLong()
	s := uint32(num) & 0x3f
	result := int64(uint64(basic) >> s)
	stack.PushLong(result)
}
