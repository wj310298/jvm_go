package classpath

import (
	"path/filepath"
	"os"
)

/**
 * Classpath结构体有三个字段,分别存放三种类路径.
 */

type Classpath struct {
	bootClasspath	Entry
	extClasspath	Entry
	userClasspath	Entry
}

/**
 * 使用-Xjre选项解析启动类路径和扩展类路径,使用-classpath/cp选项解析用户类路径.
 */
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

/**
 * 优先使用用户输入的-Xjre选项作为jre目录.如果没有输入该选项,则在当前目录下寻找jre目录.
 * 如果找不到,尝试使用JAVA_HOME环境变量.
 */
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

/**
 * 如果用户没有提供-classpath/cp选项,则使用当前目录作为用户类路径.
 * 依次从启动类路径,扩展类路径和用户类路径中搜索class文件
 */
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}
