package classpath

/*
 * Java虚拟机规范并没有规定虚拟机应该从哪儿寻找类,因此不同的虚拟机实现可以采用不同的方法.Oracle的Java虚拟机实现根据类路径来搜索类.
 * 按照搜索的先后顺序,类路径可以分为一下三个部分:
 * 1.启动类路径(bootstrap classpath)
 * 2.扩展类路径(extension classpath)
 * 3.用户类路径(user classpath)
 * Entry接口表示类路径项.
 * 常量pathListSeparator是string类型,存放路径分隔符,后面会用到.
 * Entry接口中有两个方法.readClass方法负责寻找和加载class文件;String方法的作用相当于Java中的toString,用于返回变量的字符串表示.
 * readClass方法的参数是class文件的相对路径,路径之间用/分隔,文件名有.class后缀.返回值是读取到的字节数据,最终定位到class文件的Entry,以及错误信息.
 */

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
	strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
