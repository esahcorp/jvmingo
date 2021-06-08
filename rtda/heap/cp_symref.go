package heap

type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (sr *SymRef) ResolvedClass() *Class {
	if sr.class == nil {
		sr.resolveClassRef()
	}
	return sr.class
}

func (sr *SymRef) resolveClassRef() {
	visitor := sr.cp.class
	owner := visitor.loader.LoadClass(sr.className)
	if !owner.isAccessibleTo(visitor) {
		panic("java.lang.IllegalAccessError")
	}
	sr.class = owner
}
