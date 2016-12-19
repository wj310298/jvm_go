package classfile

/**
 * EnclosingMethod属性是可选的定长属性,位于ClassFile结构的属性表.
 * 当且仅当Class为局部类或者匿名类时,才能具有EnclosingMethod属性.
 * 一个类最多只能有一个EnclosingMethod属性.
 * EnclosingMethod_attribute {
       u2 attribute_name_index;
       u4 attribute_length;
       u2 class_index;
       u2 method_index;
   }
 * attribute_name_index项的值必须是一个对常量池的有效索引.常量池在该索引处的项必须是CONSTANT_Utf8_info结构,表示字符串"EnclosingMethod".
 * attribute_length项的值固定为4.
 * class_index项的值必须是一个对常量池的有效索引.常量池在该索引出的项必须是CONSTANT_Class_info结构,表示包含当前类声明的最内层类.
 * 如果当前类不是在某个方法或初始化方法中直接包含(Enclosed),那么method_index值为0,
 * 否则method_index项的值必须是对常量池的一个有效索引,常量池在该索引处的成员必须是CONSTANT_NameAndType_info结构,表示由class_index属性引用的类的对应方法的方法名和方法类型.
 * Java编译器有责任在语法上保证通过method_index确定的方法是语法上最接近那个包含EnclosingMethod属性的类的方法(Closest Lexically Enclosing Method).
 */

type EnclosingMethodAttribute struct {
	cp		ConstantPool
	classIndex	uint16
	methodIndex	uint16
}

func (self *EnclosingMethodAttribute) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.methodIndex = reader.readUint16()
}

func (self *EnclosingMethodAttribute) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *EnclosingMethodAttribute) MethodNameAndDescriptor() (string, string) {
	if self.methodIndex > 0 {
		return self.cp.getNameAndType(self.methodIndex)
	} else {
		return "", ""
	}
}
