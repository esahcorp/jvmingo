package classfile

type ExceptionAttribute struct {
	exceptionIndexTable []uint16
}

func (e *ExceptionAttribute) readInfo(reader *ClassReader) {
	e.exceptionIndexTable = reader.readUint16s()
}

func (e *ExceptionAttribute) ExceptionIndexTable() []uint16 {
	return e.exceptionIndexTable
}
