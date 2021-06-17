package heap

/* Java Object */

type Object struct {
	class *Class
	data  interface{} // 实例变量，包括数组，静态变量在类变量里面
	extra interface{}
}

func (obj *Object) Extra() interface{} {
	return obj.extra
}

func (obj *Object) SetExtra(extra interface{}) {
	obj.extra = extra
}

func (obj *Object) Class() *Class {
	return obj.class
}

func (obj *Object) Fields() Slots {
	return obj.data.(Slots)
}

func (obj *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(obj.class)
}

func (obj *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := obj.class.getField(name, descriptor, false)
	slots := obj.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (obj *Object) GetRefVar(name, descriptor string) *Object {
	field := obj.class.getField(name, descriptor, false)
	slots := obj.data.(Slots)
	return slots.GetRef(field.slotId)
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}
