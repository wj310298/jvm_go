package classfile

/**
 * CONSTANT_MethodHandle_info 结构用于表示方法句柄
 * CONSTANT_MethodHandle_info {
       u1 tag;
       u1 reference_kind; // 在1至9之间(包括1和9),它决定了方法句柄的类型
       u2 reference_index; // 对常量池的有效索引
   }
 */

type ConstantMethodHandleInfo struct {
	referenceKind uint8
	referenceIndex uint16
}

func (self *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	self.referenceKind = reader.readUint8()
	self.referenceIndex = reader.readUint16()
}

/**
 * CONSTANT_MethodType_info结构用于表示方法类型
 * CONSTANT_MethodType_info {
       u1 tag;
       u2 descriptor_index;
   }
 */

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (self *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	self.descriptorIndex = reader.readUint16()
}

/**
 * CONSTANT_InvokeDynamic_info用于表示invokedynamic指令所使用到的引导方法(Bootstrap Method),引导方法使用到动态调用名称(Dynamic Invocation Name),
 * 参数和请求返回类型,以及可以选择性的附加被称为静态参数(Static Arguments)的常量序列.
 * CONSTANT_InvokeDynamic_info {
       u1 tag;
       u2 bootstrap_method_attr_index;
       u2 name_and_type_index;
   }
 */

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex uint16
}

func (self *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	self.bootstrapMethodAttrIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}
