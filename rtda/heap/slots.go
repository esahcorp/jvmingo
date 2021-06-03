package heap

import "math"

type Slot struct {
	num int32 // Java int is 4 byte == 32 bit
	ref *Object
}

type Slots []Slot

func newSlots(slotCount uint) Slots {
	if slotCount > 0 {
		return make([]Slot, slotCount)
	}
	return nil
}

func (s Slots) SetInt(index uint, val int32) {
	s[index].num = val
}

func (s Slots) GetInt(index uint) int32 {
	return s[index].num
}

func (s Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	s[index].num = int32(bits)
}

func (s Slots) GetFloat(index uint) float32 {
	bits := uint32(s[index].num)
	return math.Float32frombits(bits)
}

func (s Slots) SetLong(index uint, val int64) {
	s[index].num = int32(val)
	s[index+1].num = int32(val >> 32) // A Slot store 32 bit, so 64 bit number need 2 Slot
}

func (s Slots) GetLong(index uint) int64 {
	low := uint32(s[index].num)
	high := uint32(s[index+1].num)
	return int64(high)<<32 | int64(low)
}

func (s Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	s.SetLong(index, int64(bits))
}

func (s Slots) GetDouble(index uint) float64 {
	long := s.GetLong(index)
	bits := uint64(long)
	return math.Float64frombits(bits)
}

func (s Slots) SetRef(index uint, ref *Object) {
	s[index].ref = ref
}

func (s Slots) GetRef(index uint) *Object {
	return s[index].ref
}
