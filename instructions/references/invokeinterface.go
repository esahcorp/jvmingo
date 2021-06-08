package references

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
	"jvmingo/rtda/heap"
)

type INVOKE_INTERFACE struct {
	index uint
	//count uint8 // 参数对应的 slot 数
	//zero  uint8
}

func (inst *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	inst.index = uint(reader.ReadUint16())
	reader.ReadUint8() // count
	reader.ReadUint8() // must be 0
}

func (inst *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(inst.index).(*heap.InterfaceMethodRef)
	method := methodRef.ResolvedInterfaceMethod()
	if method.IsStatic() || method.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(method.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}
