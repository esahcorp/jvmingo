package control

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

/*
lookupswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
npairs1
npairs2
npairs3
npairs4
match-offset pairs...
*/

// Match value

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (inst *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	inst.defaultOffset = reader.ReadInt32()
	inst.npairs = reader.ReadInt32()
	inst.matchOffsets = reader.ReadInt32s(inst.npairs * 2)
}

func (inst *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < inst.npairs*2; i += 2 {
		if inst.matchOffsets[i] == key {
			offset := inst.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(inst.defaultOffset))
}
