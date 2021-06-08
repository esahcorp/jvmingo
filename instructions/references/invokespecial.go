package references

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
	"jvmingo/rtda/heap"
)

// 私有方法和构造函数不存在继承链，不需要动态绑定，比较简单

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

// Invoke instance method

func (inst *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(inst.Index).(*heap.MethodRef)
	owner := methodRef.ResolvedClass()
	method := methodRef.ResolvedMethod()

	if method.Name() == "<init>" && method.Class() != owner {
		panic("java.lang.NoSuchMethodError")
	}
	if method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(method.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	if method.IsProtected() &&
		method.Class().IsSuperClassOf(currentClass) &&
		method.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {

		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked := method
	if currentClass.IsSuper() &&
		owner.IsSuperClassOf(currentClass) &&
		method.Name() != "<init>" {
		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(),
			methodRef.Name(), methodRef.Descriptor())
	}

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}
