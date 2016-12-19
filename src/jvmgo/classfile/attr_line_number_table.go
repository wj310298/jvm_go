package classfile

/**
 * LineNumberTable属性表存放方法的行号信息.属于调试信息,不是运行时必需的.
 * 在使用javac编译器编译Java程序时,默认会在class文件中生成这些信息.可以使用javac提供的-g:none选项来关闭这些信息的生成.
 * LineNumberTable_attribute {
       u2 attribute_name_index;
       u4 attribute_length;
       u2 line_number_table_length;
       {   u2 start;
           u2 line_number;
       } line_number_table[line_number_table_length];
   }
 */

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc		uint16
	lineNumber	uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:	reader.readUint16(),
			lineNumber:	reader.readUint16(),
		}
	}
}
