package references

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
	"jvmingo/rtda/heap"
)

type NEW struct {
	base.Index16Instruction // 16 bit operand, index of runtime cp link to class ref
}

func (inst *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(inst.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
