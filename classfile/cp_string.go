package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (i *ConstantStringInfo) readInfo(reader *ClassReader) {
	i.stringIndex = reader.readUint16()
}

func (i *ConstantStringInfo) String() string {
	return i.cp.getUtf8(i.stringIndex)
}
