package classfile

/**
 * 常量池存放着常量信息,包括数字和字符串常量,类和接口名,字段和方法名,etc
 * 常量池实际上是一个表.
 * 1.表头给出的常量池大小比实际大1.
 * 2.有效的常量池索引是1~n-1.0是无效索引,表示不指向任何常量.
 * 3.CONSTANT_Long_info和CONSTANT_Double_info各占两个位置.如果常量池中存在这两种常量,实际的常量数量比n-1还要少,而且1~n-1的某些数也会变成无效索引.
 */

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ //占两个位置
		}
	}
	return cp
}

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
