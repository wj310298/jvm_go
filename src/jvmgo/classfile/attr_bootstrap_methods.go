package classfile

/**
 * BootstrapMethods属性是一个变长属性,位于ClassFile结构的属性表中.用于保存invokedynamic指令引用的引导方法限定符.
 * 如果某个ClassFile结构的常量池中有至少一个CONSTANT_InvokeDynamic_info项,那么这个ClassFile结构的属性表中必须有一个明确的BootstrapMethods属性.
 * ClassFile结构的属性表中最多只能有一个BootstrapMethods属性.
 * BootstrapMethods_attribute {
       u2 attribute_name_index;
       u4 attribute_length;
       u2 num_bootstrap_methods;
       {   u2 bootstrap_method_ref;
           u2 num_bootstrap_arguments;
           u2 bootstrap_arguments[num_bootstrap_arguments];
       } bootstrap_methods[num_bootstrap_methods];
 * num_bootstrap_methods项的值给出了bootstrap_methods[]数组中的引导方法限定符的数量.
 * bootstrap_methods[]数组的每个成员包含一个指向CONSTANT_MethodHandle结构的索引值,它代表了一个引导方法.还包含了这个引导方法静态参数的序列(可能为空).
 * bootstrap_method_ref项的值必须是一个对常量池的有效索引.常量池在该索引处的值必须是一个CONSTANT_MethodHandle_info结构.
 * 此CONSTANT_MethodHandle_info结构的reference_kind项应为值6(REF_invokeStatic)或8(REF_newInvokeSpecial),否则在invokedynamic指令解析调用点限定符时,引导方法会执行失败.
 * num_bootstrap_arguments项的值给出了bootstrap_arguments[]数组成员的数量.
 * bootstrap_arguments[]数组的每个成员必须是一个对常量池的有效索引.
 * 常量池在该索引出必须是下列结构之一:
 * CONSTANT_String_info,CONSTANT_Class_info,CONSTANT_Integer_info,CONSTANT_Long_info,CONSTANT_Float_info,CONSTANT_Double_info,
 * CONSTANT_MethodHandle_info或CONSTANT_MethodType_info
 */

type BootstrapMethodsAttribute struct {
	bootstrapMethods []*BootstrapMethod
}

func (self *BootstrapMethodsAttribute) readInfo(reader *ClassReader) {
	numBootstrapMethods := reader.readUint16()
	self.bootstrapMethods = make([]*BootstrapMethod, numBootstrapMethods)
	for i := range self.bootstrapMethods {
		self.bootstrapMethods[i] = &BootstrapMethod{
			bootstrapMethodRef:	reader.readUint16(),
			bootstrapArguments:	reader.readUint16s(),
		}
	}
}

type BootstrapMethod struct {
	bootstrapMethodRef uint16
	bootstrapArguments []uint16
}
