package classfile

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (i *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	i.classIndex = reader.readUint16()
	i.nameAndTypeIndex = reader.readUint16()
}

func (i ConstantMemberrefInfo) ClassName() string {
	return i.cp.getClassName(i.classIndex)
}

func (i ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return i.cp.getNameAndType(i.nameAndTypeIndex)
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}
