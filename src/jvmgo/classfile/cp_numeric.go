package classfile

import "math"

/**
 * CONSTANT_Integer_info使用4字节存储整数常量
   CONSTANT_Integer_info {
       u1 tag;
       u4 bytes;
   }
 * CONSTANT_Integer_info正好可以容纳一个Java的int型常量,但实际上比int更小的boolean,byte,short和char类型的常量也放在CONSTANT_Integer_info中.
 * 编译器给ClassFileTest类的INT字段生成了一个CONSTANT_Integer_info常量
 */

type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

/**
 * CONSTANT_Float_info使用4字节存储IEEE754单精度浮点数常量
 * CONSTANT_Float_info {
       u1 tag;
       u4 bytes;
   }
 */
type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

/**
 * CONSTANT_Long_info使用8字节存储整数常量
 * CONSTANT_Long_info {
       u1 tag;
       u4 high_bytes;
       u4 low_bytes;
   }
 */

type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

/**
 * CONSTANT_Double_info使用8字节存储IEEE754双精度浮点数
 * CONSTANT_Double_info {
       u1 tag;
       u4 high_bytes;
       u4 low_bytes;
   }
 */

type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}
