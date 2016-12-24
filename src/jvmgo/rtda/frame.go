package rtda

/*
    执行方法所需的局部变量表大小和操作数栈深度是由编译器预先计算好的,存储在class文件method_info结构的Code属性中.

 |----------|     |----------|     |----------|     |----------|
 |  Stack   |  |->|  Frame   |  |->|  Frame   |  |->|  Frame   |
 |----------|  |  |----------|  |  |----------|  |  |----------|
 |   _top   |--|  |  lower   |--|  |  lower   |--|  |  lower   |
 |__________|     |__________|     |__________|     |__________|

 */

type Frame struct {
	lower		*Frame
	localVars	LocalVars
	operandStack	*OperandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:	newLocalVars(maxLocals),
		operandStack:	newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
