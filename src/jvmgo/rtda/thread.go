package rtda

/*
    和堆一样,Java虚拟机规范对Java虚拟机栈的约束也相当宽松.Java虚拟机栈可以是连续的空间,也可以不连续;可以是固定大小,也可以在运行时动态扩展.
如果Java虚拟机栈由大小限制,且执行线程所需的栈空间超出了这个限制,会导致StackOverflowError异常抛出.
如果Java虚拟机栈可以动态扩展,但是内存已经耗尽,会导致OutOfMemoryError异常抛出.
    Java命令提供了-Xss选项来设置Java虚拟机栈大小.
 */

type Thread struct {
	pc	int
	stack	*Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) PC() int { return self.pc }

func (self *Thread) SetPC(pc int) { self.pc = pc }

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
