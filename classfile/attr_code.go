package classfile

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocal       uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (attr *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return attr.exceptionTable
}

func (attr *CodeAttribute) MaxStack() uint {
	return uint(attr.maxStack)
}

func (attr *CodeAttribute) MaxLocal() uint {
	return uint(attr.maxLocal)
}

func (attr *CodeAttribute) Code() []byte {
	return attr.code
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

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (e *ExceptionTableEntry) StartPc() uint16 {
	return e.startPc
}

func (e *ExceptionTableEntry) EndPc() uint16 {
	return e.endPc
}

func (e *ExceptionTableEntry) HandlerPc() uint16 {
	return e.handlerPc
}

func (e *ExceptionTableEntry) CatchType() uint16 {
	return e.catchType
}
