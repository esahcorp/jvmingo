package heap

/* Java Object */

type Object struct {
	class *Class
	data  interface{} // 实例变量，包括数组，静态变量在类变量里面
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

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}
