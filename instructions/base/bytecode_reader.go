package base

/* Java Bytecode Reader, read instruction
and operand behind instruction from bytecode */

type BytecodeReader struct {
	code []byte
	pc   int
}

func (reader *BytecodeReader) Reset(code []byte, pc int) {
	reader.code = code
	reader.pc = pc
}

func (reader *BytecodeReader) PC() int {
	return reader.pc
}

func (reader *BytecodeReader) ReadUint8() uint8 {
	i := reader.code[reader.pc]
	reader.pc++
	return i
}

func (reader *BytecodeReader) ReadUint16() uint16 {
	high := uint16(reader.ReadUint8())
	low := uint16(reader.ReadUint8())
	return (high << 8) | low
}

func (reader *BytecodeReader) ReadInt8() int8 {
	return int8(reader.ReadUint8())
}

func (reader *BytecodeReader) ReadInt16() int16 {
	return int16(reader.ReadUint16())
}

func (reader *BytecodeReader) ReadInt32() int32 {
	first := int32(reader.ReadUint8())
	second := int32(reader.ReadUint8())
	third := int32(reader.ReadUint8())
	fourth := int32(reader.ReadUint8())
	return (first << 24) | (second << 16) | (third << 8) | fourth
}

func (reader *BytecodeReader) SkipPadding() {
	for reader.pc%4 != 0 {
		reader.ReadUint8()
	}
}

func (reader *BytecodeReader) ReadInt32s(count int32) []int32 {
	ints := make([]int32, count)
	for i := range ints {
		ints[i] = reader.ReadInt32()
	}
	return ints
}
