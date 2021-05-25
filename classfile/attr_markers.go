package classfile

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct {
}

func (attr *MarkerAttribute) readInfo(_ *ClassReader) {
	// read nothing
}
