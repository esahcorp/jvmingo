package heap

// other 类型可以赋值给 class

func (class *Class) isAssignableFrom(other *Class) bool {
	o, t := other, class
	if o == t {
		return true
	}
	if o.IsArray() && t.IsArray() {
		oc := o.ComponentClass()
		tc := t.ComponentClass()
		// oc 与 tc 是同一种类型
		// 或 oc 和 tc 都是引用类型，并且 oc 可以强制转换为 tc
		return oc == tc || tc.isAssignableFrom(oc)
	}
	if o.IsArray() && !t.IsArray() {
		if t.IsInterface() {
			return t.isJlCloneable() || t.isJlSerializable() // 数组可以强制转换为 Cloneable 和 Serializable
		} else {
			return t.isJlObject() // 数组可以强制转换为 Object
		}
	}
	if !o.IsArray() {
		if o.IsInterface() && t.IsInterface() {
			return t.isSuperInterfaceOf(o)
		}
		if o.IsInterface() && !t.IsInterface() {
			return t.isJlObject() // 接口可以强制转换为 Object
		}
		if !o.IsInterface() && t.IsInterface() {
			return o.IsImplements(t) // o 实现了 t
		}
		if !o.IsInterface() && !t.IsInterface() {
			return o.IsSubClassOf(t) // o 是 t 的子类
		}
	}
	return false
}

func (class *Class) IsSubClassOf(other *Class) bool {
	for c := class.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (class *Class) IsImplements(iface *Class) bool {
	for c := class; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

func (class *Class) isSubInterfaceOf(iface *Class) bool {
	for _, si := range class.interfaces {
		if si == iface || si.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

func (class *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(class)
}

func (class *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.isSuperInterfaceOf(class)
}
