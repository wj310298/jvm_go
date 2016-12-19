package classfile

import "fmt"

/**
 * 作为类(或者接口)信息的载体,每个class文件都完整地定义一个类.
 * 构成class文件的基本数据单位是字节,可以把整个class文件当成一个字节流来处理.稍大一些的数据由连续多个字节构成,这些数据在class文件中以大端(big-endian)方式存储.
 * 整个class文件被描述为一个ClassFile结构,代码如下:
 * ClassFile {
       u4		magic;
       u2		minor_version;
       u2		major_version;
       u2		constant_pool_count;
       cp_info		constant_pool[constant_pool_count - 1];
       u2		access_flags;
       u2		this_class;
       u2		super_class;
       u2		interfaces_count;
       u2		interfaces[interfaces_count];
       u2		fields_count;
       field_info	fields[fields_count];
       u2		methods_count;
       method_info	methods[methods_count];
       u2		attributes_count;
       attribute_info	attributes[attributes_count];
   }
 */

type ClassFile struct {
	//magic 		uint32
	minorVersion	uint16
	majorVersion	uint16
	constantPool	ConstantPool
	accessFlags	uint16		//类访问标志,bitmask,指出文件定义的是类还是接口,访问级别是public还是private,etc
	thisClass	uint16		//类的常量池索引
	superClass	uint16		//超类的常量池索引
	interfaces	[]uint16	//接口索引表,表中存放的是常量池索引,给出该类实现的所有接口的名字
	fields		[]*MemberInfo	//字段表
	methods		[]*MemberInfo	//方法表
	attributes	[]AttributeInfo
}

/**
 * 把[]byte解析成ClassFile结构体
 */
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{data:classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

/**
 * 依次调用其他方法解析class文件
 */
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

/**
 * 魔数 0xCAFEBABE
 */
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

/**
 * 检查版本号
 */
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

/**
 * 从常量池查找类名
 */
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

/**
 * 从常量池查找超类名
 */
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" // 只有java.lang.Object没有超类
}

/**
 * 从常量池查找接口名
 */
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

