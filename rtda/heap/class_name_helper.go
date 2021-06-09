package heap

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}

func toDescriptor(className string) string {
	if className[0] == '[' {
		return className
	}
	if d, ok := primitiveTypes[className]; ok {
		return d
	}
	return "L" + className + ";"
}

/*
根据多维数组的类名得到数组元素的类名
[D -> double
[[D -> [D
[Ljava/lang/String; -> java/lang/Sting
D -> panic
*/
func getComponentClassName(className string) string {
	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: " + className)
}

func toClassName(descriptor string) string {
	if descriptor[0] == '[' {
		return descriptor
	}
	if descriptor[0] == 'L' {
		return descriptor[1 : len(descriptor)-1] // 'L' 'java/lang/String' ';'
	}
	for className, d := range primitiveTypes {
		if d == descriptor {
			return className
		}
	}
	panic("Invalid descriptor: " + descriptor)
}
