package heap

import (
	"jvmingo/classfile"
	"strings"
)

// Class information

type Class struct {
	accessFlags       uint16
	name              string // thisClassName
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint    // 示例变量占用空间大小
	staticSlotCount   uint    // 静态变量占用空间大小
	staticVars        Slots   // 静态变量值
	initStarted       bool    // 类初始化标志
	jClass            *Object // java.lang.Class 实例
}

// Convert ClassFile to Class
func newClass(cf *classfile.ClassFile) *Class {
	class := new(Class)
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (class *Class) IsPublic() bool {
	return 0 != class.accessFlags&ACC_PUBLIC
}

func (class *Class) IsFinal() bool {
	return 0 != class.accessFlags&ACC_FINAL
}
func (class *Class) IsSuper() bool {
	return 0 != class.accessFlags&ACC_SUPER
}
func (class *Class) IsInterface() bool {
	return 0 != class.accessFlags&ACC_INTERFACE
}
func (class *Class) IsAbstract() bool {
	return 0 != class.accessFlags&ACC_ABSTRACT
}
func (class *Class) IsSynthetic() bool {
	return 0 != class.accessFlags&ACC_SYNTHETIC
}
func (class *Class) IsAnnotation() bool {
	return 0 != class.accessFlags&ACC_ANNOTATION
}
func (class *Class) IsEnum() bool {
	return 0 != class.accessFlags&ACC_ENUM
}
func (class *Class) Name() string {
	return class.name
}

// getter

func (class *Class) ConstantPool() *ConstantPool {
	return class.constantPool
}
func (class *Class) Loader() *ClassLoader {
	return class.loader
}
func (class *Class) Fields() []*Field {
	return class.fields
}
func (class *Class) Methods() []*Method {
	return class.methods
}
func (class *Class) SuperClass() *Class {
	return class.superClass
}
func (class *Class) StaticVars() Slots {
	return class.staticVars
}
func (class *Class) JClass() *Object {
	return class.jClass
}
func (class *Class) InitStarted() bool {
	return class.initStarted
}

func (class *Class) StartInit() {
	class.initStarted = true
}

func (class *Class) isAccessibleTo(other *Class) bool {
	return class.IsPublic() || class.GetPackageName() == other.GetPackageName()
}

func (class *Class) GetPackageName() string {
	if i := strings.LastIndex(class.name, "/"); i >= 0 {
		return class.name[:i]
	}
	return ""
}

func (class *Class) GetMainMethod() *Method {
	return class.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (class *Class) GetClinitMethod() *Method {
	return class.getStaticMethod("<clinit>", "()V")
}

func (class *Class) getStaticMethod(name, descriptor string) *Method {
	for _, m := range class.methods {
		if m.IsStatic() &&
			m.Name() == name && m.Descriptor() == descriptor {
			return m
		}
	}
	return nil
}

func (class *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.IsStatic() == isStatic &&
				method.name == name &&
				method.descriptor == descriptor {

				return method
			}
		}
	}
	return nil
}

func (class *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := class; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic &&
				field.name == name &&
				field.descriptor == descriptor {

				return field
			}
		}
	}
	return nil
}

func (class *Class) isJlObject() bool {
	return class.name == "java/lang/Object"
}

func (class *Class) isJlCloneable() bool {
	return class.name == "java/lang/Cloneable"
}

func (class *Class) isJlSerializable() bool {
	return class.name == "java/io/Serializable"
}

func (class *Class) NewObject() *Object {
	return newObject(class)
}

func (class *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(class.name)
	return class.loader.LoadClass(arrayClassName)
}

func (class *Class) JavaName() string {
	return strings.Replace(class.name, "/", ".", -1)
}
