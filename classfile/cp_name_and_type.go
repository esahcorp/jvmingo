package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (i *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	i.nameIndex = reader.readUint16()
	i.descriptorIndex = reader.readUint16()
}
