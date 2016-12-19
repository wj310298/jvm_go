package classfile

/**
 * 属性可以扩展,不同虚拟机实现可以定义自己的属性类型.由于这个原因,Java虚拟机规范没有使用tag,而是使用属性名来区别不同的属性.
 * 属性数据放在属性名之后的u1表中,这样Java虚拟机实现就可以跳过自己无法识别的属性.
 * tips: 属性表中存放的属性名实际上并不是编码后的字符串,而是常量池索引.
 * attribute_info {
       u2 attribute_name_index;
       u4 attribute_length;
       u1 info[attribute_length];
 */

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code": return &CodeAttribute{cp: cp}
	case "ConstantValue": return &ConstantValueAttribute{}
	case "Deprecated": return &DeprecatedAttribute{}
	case "Exceptions": return &ExceptionsAttribute{}
	case "LineNumberTable": return &LineNumberTableAttribute{}
	case "LocalVariableTable": return &LocalVariableTableAttribute{}
	case "SourceFile": return &SourceFileAttribute{cp: cp}
	case "Synthetic": return &SyntheticAttribute{}
	default: return &UnparsedAttribute{name: attrName, length: attrLen, info: nil}
	}
}
