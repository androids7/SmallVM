package rtarea

type Frame struct {
	localVariables *LocalVariables
	operandStack   *OperandStack
	next           *Frame
}

func NewFrame(maxLocals uint, maxStack uint) *Frame {
	frame := &Frame{}
	frame.next = nil
	frame.localVariables = NewLocalVariables(maxLocals)
	frame.operandStack = NewOperandStack(maxStack)
	return frame
}