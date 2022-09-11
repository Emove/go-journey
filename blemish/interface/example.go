package _interface

type Buffer struct {
	b []byte
}

// Write 向 Buffer 中写入bytes
// 在这个方法里，我实现了向 Buffer 中写入字节的功能
// 但是在无意间，我实现了 io.Writer
// 若我这时，想要查看那些地方引用了 Write 方法
// 会找出一堆 io.Writer.Write 方法的引用，而无法准确的找出对于 Buffer.Write 的引用
// 这是隐式实现机制的一个弊端，而非IDE设计不完善的锅
func (buf *Buffer) Write(b []byte) (n int, err error) {
	buf.b = append(buf.b, b...)
	return len(b), nil
}
