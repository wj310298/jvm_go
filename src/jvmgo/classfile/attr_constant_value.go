package classfile

/**
 * ConstantValue是定长属性,只会出现在field_info结构中,用于表示常量表达式的值.
 * ConstantValue_attribute {
       u2 attribute_name_index;
       u4 attribute_length;
       u2 constantValue_index;
   }
 * attrbute_length的值必须是2.constantValue_index是常量池索引,但具体指向哪种常量因字段类型而异.
 */

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
