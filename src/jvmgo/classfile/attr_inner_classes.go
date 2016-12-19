package classfile

/**
 * InnerClasses属性是一个变长属性,位于ClassFile结构的属性表.
 * 为了方便说明特别定义一个表示类或接口的Class格式为C.
 * 如果C的常量池中包含某个CONSTANT_Class_info成员,且这个成员所表示的类或接口不属于任何一个包,那么C的ClassFile结构的属性表中就必须含有对应的InnerClasses属性.
 * InnerClasses属性是在JDK1.1中为了支持内部类和内部接口而引入的.
 * InnerClasses_attribute {
       u2 attribute_name_index;
       u4 attribute_length;
       u2 number_of_classes;
       {   u2 inner_class_info_index;
           u2 outer_class_info_index;
           u2 inner_name_index;
           u2 inner_class_access_flags;
       } classes[number_of_classes];
   }
 * attribute_length项的值给出了这个属性的长度,不包括前6个字节.
 * number_of_classes项的值表示classes[]数组的成员数量.
 * 常量池中的每个CONSTANT_Class_info结构如果表示的类或接口并非某个包的成员,则每个类或接口在classes[]数组中都有一个成员与之对应.
 * classes[]数组中每个成员包含以下4个项:
 * 1.inner_class_info_index项的值必须是一个对常量池的有效索引.常量池在该索引处的项必须是CONSTANT_Class_info结构,表示接口C.
 * 2.如果C不是类或接口的成员(也就是C为顶层类或接口,局部类或匿名类),那么outer_class_info_index项的值为0,
 * 否则这个项的值必须是对常量池的一个有效索引,常量池在该索引处的项必须是CONSTANT_Class_info结构,代表一个类或接口,C为这个类或接口的成员.
 * 3.如果C是匿名类,inner_name_index项的值则必须为0.
 * 否则这个项的值必须是对常量池的一个有效索引,常量池在该索引处的项必须CONSTANT_Utf8_info结构,表示C的Class文件在对应的源文件中定义的C的原始简单名称.
 * 4.inner_class_access_flags项的值是一个掩码标志,用于定义Class文件对应的源文件中C的访问权和基本属性.用于编译器在无法访问源文件时可以恢复C的原始信息.
 * 内部类访问全和基础属性标志
   标记名		值	含义
   ACC_PUBLIC		0x0001	源文件定义public
   ACC_PRIVATE		0x0002	源文件定义private
   ACC_PROTECTED	0x0004	源文件定义protected
   ACC_STATIC		0x0008	源文件定义static
   ACC_FINAL		0x0010	源文件定义final
   ACC_INTERFACE	0x0200	源文件定义interface
   ACC_ABSTRACT		0x0400	源文件定义abstract
   ACC_SYNTHETIC	0x1000	声明synthetic,非源文件定义
   ACC_ANNOTATION	0x2000	声明annotation
   ACC_ENUM		0x4000	声明enum
 * 如果Class文件的版本号为51.0或更高,属性表中有InnerClasses属性,并且InnerClasses属性的classes[]数组中的inner_name_index项的值为0,则它对应的outer_class_info_index项的值也必须为0.
 * Oracle的Java虚拟机实现不会检查InnerClasses属性和某个属性引用的类或接口的Class文件的一致性.
 */

type InnerClassesAttribute struct {
	classes []*InnerClassInfo
}

type InnerClassInfo struct {
	innerClassInfoIndex	uint16
	outerClassInfoIndex	uint16
	innerNameIndex		uint16
	innerClassAccessFlags	uint16
}

func (self *InnerClassesAttribute) readInfo(reader *ClassReader) {
	numberOfClasses := reader.readUint16()
	self.classes = make([]*InnerClassInfo, numberOfClasses)
	for i := range self.classes {
		self.classes[i] = &InnerClassInfo{
			innerClassInfoIndex:	reader.readUint16(),
			outerClassInfoIndex:	reader.readUint16(),
			innerNameIndex:		reader.readUint16(),
			innerClassAccessFlags:	reader.readUint16(),
		}
	}
}
