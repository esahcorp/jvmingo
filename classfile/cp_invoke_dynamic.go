package classfile

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/

type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (i *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	i.referenceKind = reader.readUint8()
	i.referenceIndex = reader.readUint16()
}

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (i *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	i.descriptorIndex = reader.readUint16()
}

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (i *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	i.bootstrapMethodAttrIndex = reader.readUint16()
	i.nameAndTypeIndex = reader.readUint16()
}
