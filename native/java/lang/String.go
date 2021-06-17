package lang

import (
	"jvmingo/native"
	"jvmingo/rtda"
	"jvmingo/rtda/heap"
)

func init() {
	native.Register("java/lang/String", "intern",
		"()Ljava/lang/String;", intern)
}

// public native String intern();
func intern(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
