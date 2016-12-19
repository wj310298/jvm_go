package classfile

/*
 * Signature属性是可选的定长属性,位于ClassFile,field_info或method_info结构的属性表中.
 * 在Java语言中,任何类,接口,初始化方法或成员的泛型签名如果包含了类型变量(Type Variables)或参数化类型(Parameterized Types),则Signature属性会为它记录泛型签名信息.
 * Signature_attribute {
       u2 attribute_name_index;
       u4 attribute_length;
       u2 signature_index;
   }
 * signature_index项的值必须是一个对常量池的有效索引.
 * 常量池在该索引处的项必须是CONSTANT_Utf8_info结构,表示类签名或方法类型签名或字段类型签名:
 * 如果当前的Signature属性是ClassFile结构的属性,则这个结构表示类签名,
 * 如果当前的Signature属性是method_info结构的属性,则这个结构表示方法类型签名,如果当前Signature属性是field_info结构的属性,则这个结构表示字段类型签名.
 */

type SignatureAttribute struct {
	cp             ConstantPool
	signatureIndex uint16
}

func (self *SignatureAttribute) readInfo(reader *ClassReader) {
	self.signatureIndex = reader.readUint16()
}

func (self *SignatureAttribute) Signature() string {
	return self.cp.getUtf8(self.signatureIndex)
}
