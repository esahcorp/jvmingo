package rtda

import (
	"jvmingo/rtda/heap"
	"math"
)

/* Local Variable Table of JVM Stack */

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (lv LocalVars) SetInt(index uint, val int32) {
	lv[index].num = val
}

func (lv LocalVars) GetInt(index uint) int32 {
	return lv[index].num
}

func (lv LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	lv[index].num = int32(bits)
}

func (lv LocalVars) GetFloat(index uint) float32 {
	bits := uint32(lv[index].num)
	return math.Float32frombits(bits)
}

func (lv LocalVars) SetLong(index uint, val int64) {
	lv[index].num = int32(val)
	lv[index+1].num = int32(val >> 32) // A Slot store 32 bit, so 64 bit number need 2 Slot
}

func (lv LocalVars) GetLong(index uint) int64 {
	low := uint32(lv[index].num)
	high := uint32(lv[index+1].num)
	return int64(high)<<32 | int64(low)
}

func (lv LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	lv.SetLong(index, int64(bits))
}

func (lv LocalVars) GetDouble(index uint) float64 {
	long := lv.GetLong(index)
	bits := uint64(long)
	return math.Float64frombits(bits)
}

func (lv LocalVars) SetRef(index uint, ref *heap.Object) {
	lv[index].ref = ref
}

func (lv LocalVars) GetRef(index uint) *heap.Object {
	return lv[index].ref
}
