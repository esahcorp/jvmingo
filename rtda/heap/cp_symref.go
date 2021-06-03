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
	d := sr.cp.class
	c := d.loader.LoadClass(sr.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	sr.class = c
}
