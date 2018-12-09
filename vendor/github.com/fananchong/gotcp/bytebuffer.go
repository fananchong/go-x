package gotcp

const (
	presize  = 0
	initsize = 16
)

type ByteBuffer struct {
	_buffer      []byte
	_prependSize int
	_readerIndex int
	_writerIndex int
}

func NewByteBuffer() *ByteBuffer {
	return &ByteBuffer{
		_buffer:      make([]byte, presize+initsize),
		_prependSize: presize,
		_readerIndex: presize,
		_writerIndex: presize,
	}
}

func (this *ByteBuffer) Append(buff []byte) {
	size := len(buff)
	if size == 0 {
		return
	}
	this.WrGrow(size)
	copy(this._buffer[this._writerIndex:], buff)
	this.WrFlip(size)
}

func (this *ByteBuffer) WrBuf() []byte {
	if this._writerIndex >= len(this._buffer) {
		return nil
	}
	return this._buffer[this._writerIndex:]
}

func (this *ByteBuffer) WrSize() int {
	return len(this._buffer) - this._writerIndex
}

func (this *ByteBuffer) WrFlip(size int) {
	this._writerIndex += size
}

func (this *ByteBuffer) WrGrow(size int) {
	if size > this.WrSize() {
		this.wrreserve(size)
	}
}

func (this *ByteBuffer) RdBuf() []byte {
	if this._readerIndex >= len(this._buffer) {
		return nil
	}
	return this._buffer[this._readerIndex:]
}

func (this *ByteBuffer) RdReady() bool {
	return this._writerIndex > this._readerIndex
}

func (this *ByteBuffer) RdSize() int {
	return this._writerIndex - this._readerIndex
}

func (this *ByteBuffer) RdFlip(size int) {
	if size < this.RdSize() {
		this._readerIndex += size
	} else {
		this.Reset()
	}
}

func (this *ByteBuffer) Reset() {
	this._readerIndex = this._prependSize
	this._writerIndex = this._prependSize
}

func (this *ByteBuffer) wrreserve(size int) {
	if this.WrSize()+this._readerIndex < size+this._prependSize {
		newsize := this.RdSize() + this.WrSize()
		for newsize < this._writerIndex+size {
			newsize <<= 1
		}
		tmpbuff := make([]byte, newsize+this._prependSize)
		copy(tmpbuff, this._buffer)
		this._buffer = tmpbuff
	} else {
		readable := this.RdSize()
		copy(this._buffer[this._prependSize:], this._buffer[this._readerIndex:this._writerIndex])
		this._readerIndex = this._prependSize
		this._writerIndex = this._readerIndex + readable
	}
}
