package classfile

/**
 * LocalVariableTable是可选变长属性,位于Code属性的属性表中.
 * 它被调试器用于确定方法在执行过程中局部变量的信息.
 * 在Code属性的属性表中,LocalVariableTable属性可以按照任意顺序出现.Code属性中的每个局部变量最多只能有一个LocalVariableTable属性.
 * LocalVariableTable_attribute {
       u2 attribute_name_index;
       u4 attribute_length;
       u2 local_variable_table_length;
       {   u2 start_pc;
           u2 length;
           u2 name_index;
           u2 descriptor_index;
           u2 index;
       } local_variable_table[local_variable_table_length];
   }
 */

type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc		uint16
	length		uint16
	nameIndex	uint16
	descriptorIndex	uint16
	index		uint16
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	self.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range self.localVariableTable {
		self.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:	reader.readUint16(),
			length:		reader.readUint16(),
			nameIndex:	reader.readUint16(),
			descriptorIndex:reader.readUint16(),
			index:		reader.readUint16(),
		}
	}
}
