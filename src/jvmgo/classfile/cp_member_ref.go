package classfile

/**
 * CONSTANT_Fieldref_info表示字段符号引用,CONSTANT_Methodref_info表示普通(非接口)方法符号引用,CONSTANT_Methodref_info表示接口方法符号引用.
 * 这三种常量接口一样
 * CONSTANT_Fieldref_info {
       u1 tag;
       u2 class_index; // 表示一个类或接口,当前字段或方法是这个类或接口的成员
       u2 name_and_type_index;
   }
 */

type ConstantMemberrefInfo struct {
	cp			ConstantPool
	classIndex		uint16
	nameAndTypeIndex	uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}
