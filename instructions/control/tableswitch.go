package control

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

/*
tableswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
lowbyte1
lowbyte2
lowbyte3
lowbyte4
highbyte1
highbyte2
highbyte3
highbyte4
jump offsets...
*/

// Index

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (inst *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	// 0 ~ 3 byte padding behind tableswitch opcode
	reader.SkipPadding()
	inst.defaultOffset = reader.ReadInt32()
	inst.low = reader.ReadInt32()
	inst.high = reader.ReadInt32()
	jumpOffserCount := inst.high - inst.low + 1
	inst.jumpOffsets = reader.ReadInt32s(jumpOffserCount)
}

func (inst *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= inst.low && index <= inst.high {
		offset = int(inst.jumpOffsets[index-inst.low])
	} else {
		offset = int(inst.defaultOffset)
	}

	base.Branch(frame, offset)
}
