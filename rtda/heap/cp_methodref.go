package heap

import "jvmingo/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

// 从类及其父类，类的接口及上层接口中查找指定方法，给 method 赋值

func (ref *MethodRef) ResolvedMethod() *Method {
	if ref.method == nil {
		ref.resolveMethodRef()
	}
	return ref.method
}

func (ref *MethodRef) resolveMethodRef() {
	invoker := ref.cp.class
	owner := ref.ResolvedClass()
	if owner.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupMethod(owner, ref.name, ref.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(invoker) {
		panic("java.lang.IllegalAccessError")
	}
	ref.method = method
}

func lookupMethod(owner *Class, name, descriptor string) *Method {
	method := LookupMethodInClass(owner, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(owner.interfaces, name, descriptor)
	}
	return method
}

func LookupMethodInClass(class *Class, name, descriptor string) *Method {
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

func lookupMethodInInterfaces(ifaces []*Class, name, descriptor string) *Method {
	for _, iface := range ifaces {
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}

		method := lookupMethodInInterfaces(iface.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}
	return nil
}
