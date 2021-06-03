package heap

/* Java Object */

type Object struct {
	class  *Class
	fields Slots // 示例变量值，静态变量在类变量里面
}

func (o *Object) Class() *Class {
	return o.class
}

func (o *Object) Fields() Slots {
	return o.fields
}

func (o *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(o.class)
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}
