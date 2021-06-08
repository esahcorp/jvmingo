package heap

import "jvmingo/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (ref *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if ref.method == nil {
		ref.resolveInterfaceMethodRef()
	}
	return ref.method
}

func (ref *InterfaceMethodRef) resolveInterfaceMethodRef() {
	invoker := ref.cp.class
	owner := ref.ResolvedClass()
	if !owner.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupInterfaceMethod(owner, ref.name, ref.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(invoker) {
		panic("java.lang.IllegalAccessError")
	}

	ref.method = method
}

func lookupInterfaceMethod(iface *Class, name, descriptor string) *Method {
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return lookupMethodInInterfaces(iface.interfaces, name, descriptor)
}
