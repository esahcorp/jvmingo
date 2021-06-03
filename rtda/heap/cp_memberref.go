package heap

import "jvmingo/classfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (ref *MemberRef) Name() string {
	return ref.name
}

func (ref *MemberRef) Descriptor() string {
	return ref.descriptor
}

func (ref *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	ref.className = refInfo.ClassName()
	ref.name, ref.descriptor = refInfo.NameAndDescriptor()
}
