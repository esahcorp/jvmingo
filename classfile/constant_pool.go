package classfile

import "fmt"

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpSize := int(reader.readUint16())
	cp := make([]ConstantInfo, cpSize)

	for i := 1; i < cpSize; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}

	return cp
}

func (c ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := c[index]; cpInfo != nil {
		return cpInfo
	}
	fmt.Printf("index: %v\n", index)
	panic("Invalid constant pool index!")
}

func (c ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := c.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := c.getUtf8(ntInfo.nameIndex)
	_type := c.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (c ConstantPool) getClassName(index uint16) string {
	classInfo := c.getConstantInfo(index).(*ConstantClassInfo)
	return c.getUtf8(classInfo.nameIndex)
}

func (c ConstantPool) getUtf8(index uint16) string {
	utf8Info := c.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.val
}
