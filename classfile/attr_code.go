package classfile

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocal       uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (attr *CodeAttribute) readInfo(reader *ClassReader) {
	attr.maxStack = reader.readUint16()
	attr.maxLocal = reader.readUint16()
	codeLength := reader.readUint32()
	attr.code = reader.readBytes(codeLength)
	attr.exceptionTable = readExceptionTable(reader)
	attr.attributes = readAttributes(reader, attr.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}
