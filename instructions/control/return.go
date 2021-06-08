package control

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

// Return void from method

type RETURN struct{ base.NoOperandsInstruction }

func (inst *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

// Return reference from method

type ARETURN struct{ base.NoOperandsInstruction }

func (inst *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(ref)
}

// Return double from method

type DRETURN struct{ base.NoOperandsInstruction }

func (inst *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(val)
}

// Return float from method

type FRETURN struct{ base.NoOperandsInstruction }

func (inst *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(val)
}

// Return int from method

type IRETURN struct{ base.NoOperandsInstruction }

func (inst *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(val)
}

// Return double from method

type LRETURN struct{ base.NoOperandsInstruction }

func (inst *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(val)
}
