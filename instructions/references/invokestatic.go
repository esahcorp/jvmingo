package references

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
	"jvmingo/rtda/heap"
)

// Invoke a class (static) method

type INVOKE_STATIC struct {
	base.Index16Instruction
}

func (inst *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(inst.Index).(*heap.MethodRef)
	method := methodRef.ResolvedMethod()
	if !method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	class := method.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	base.InvokeMethod(frame, method)
}
