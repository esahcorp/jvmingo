package heap

// other 类型可以赋值给 class

func (class *Class) isAssignableFrom(other *Class) bool {
	s, t := other, class
	if s == t {
		return true
	}
	if !t.IsInterface() {
		return s.IsSubClassOf(t)
	} else {
		return s.IsImplements(t)
	}
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
