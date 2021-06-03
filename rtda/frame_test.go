package rtda

import (
	"jvmingo/rtda/heap"
	"testing"
)

func TestFrame(t *testing.T) {
	frame := newFrame(NewThread(), &heap.Method{})
	testLocalVars(frame.localVars, t)
	testOperandStack(frame.operandStack, t)
}

func testLocalVars(vars LocalVars, t *testing.T) {
	t.Log("LocalVars testing: ")
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)
	t.Log(vars.GetInt(0))
	t.Log(vars.GetInt(1))
	t.Log(vars.GetLong(2))
	t.Log(vars.GetLong(4))
	t.Log(vars.GetFloat(6))
	t.Log(vars.GetDouble(7))
	t.Log(vars.GetRef(9))
}

func testOperandStack(ops *OperandStack, t *testing.T) {
	t.Log("OperandStack testing: ")
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	t.Log(ops.PopRef())
	t.Log(ops.PopDouble())
	t.Log(ops.PopFloat())
	t.Log(ops.PopLong())
	t.Log(ops.PopLong())
	t.Log(ops.PopInt())
	t.Log(ops.PopInt())
}
