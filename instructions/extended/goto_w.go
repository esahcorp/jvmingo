package extended

import (
	"jvmingo/instructions/base"
	"jvmingo/rtda"
)

type GOTO_W struct {
	offset int
}

func (inst *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	inst.offset = int(reader.ReadInt32())
}

func (inst *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, inst.offset)
}
