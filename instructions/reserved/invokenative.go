package reserved

import (
	"jvmingo/instructions/base"
	"jvmingo/native"
	_ "jvmingo/native/java/lang"
	"jvmingo/rtda"
)

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (inst *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
