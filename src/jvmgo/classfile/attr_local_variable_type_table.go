package classfile

/**
 * LocalVariableTypeTable属性是可选变长属性,位于Code的属性表.
 * 它被用于给调试器确定方法在执行中局部变量的信息,在Code属性的属性表中,LocalVariableTable属性可以按照任意顺序出现.
 * Code属性中的每个局部变量最多只能有一个LocalVariableTable属性.
 * LocalVariableTypeTable属性和LocalVariableTable属性并不相同,LocalVariableTypeTable提供签名信息而不是描述符信息.
 * 这仅仅对泛型类型有意义.泛型类型的属性会同时出现在LocalVariableTable属性和LocalVariableTypeTable属性中,其他的属性仅出现在LocalVariableTable属性表中.
 * LocalVariableTypeTable_attribute {
       u2 attribute_name_index;
       u4 attribute_length;
       u2 local_variable_type_table_length;
       {   u2 start_pc;
           u2 length;
           u2 name_index;
           u2 signature_index;
           u2 index;
       } local_variable_type_table[local_variable_type_table_length];
   }
 */

type LocalVariableTypeTableAttribute struct {
	localVariableTypeTable []*LocalVariableTypeTableEntry
}

type LocalVariableTypeTableEntry struct {
	startPc        uint16
	length         uint16
	nameIndex      uint16
	signatureIndex uint16
	index          uint16
}

func (self *LocalVariableTypeTableAttribute) readInfo(reader *ClassReader) {
	localVariableTypeTableLength := reader.readUint16()
	self.localVariableTypeTable = make([]*LocalVariableTypeTableEntry, localVariableTypeTableLength)
	for i := range self.localVariableTypeTable {
		self.localVariableTypeTable[i] = &LocalVariableTypeTableEntry{
			startPc:        reader.readUint16(),
			length:         reader.readUint16(),
			nameIndex:      reader.readUint16(),
			signatureIndex: reader.readUint16(),
			index:          reader.readUint16(),
		}
	}
}