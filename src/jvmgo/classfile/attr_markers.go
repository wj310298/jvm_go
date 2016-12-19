package classfile

/**
 * Deprecated和Synthetic仅起标记作用,不包含任何数据.
 * 这两种属性都是JDK1.1引进的,可以出现在ClassFile,field_info和method_info结构中.
 * Deprecated_attribute {
       u2 attribute_name_index;
       u4 attribute_length;
   }
   Synthetic_attribute {
       u2 attribute_name_index;
       u4 attribute_length;
   }
 * 由于不包含任何数据,所以attribute_length的值必须是0.
 * Deprecated属性用于指出类,接口,字段或方法已经不建议使用,编辑器可以根据Deprecated属性输出警告信息.
 * J2SE 5.0之前可以使用Javadoc提供的@deprecated标签指示编译器给类,接口,字段或方法添加Deprecated属性,语法格式如下:
   /** @deprecated *\/
   public void oldMethod() {}
 * 从J2SE 5.0开始,也可以使用@Deprecated注解,语法格式如下
   @Deprecated
   public void oldMethod() {}
 * Synthetic属性用来标记源文件中不存在,由编译器生成的类成员,引入Synthetic属性主要是为了支持前套类和嵌套接口.
 */

type DeprecatedAttribute struct { MarkerAttribute }

type SyntheticAttribute struct { MarkerAttribute }

type MarkerAttribute struct {}

func (*MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}