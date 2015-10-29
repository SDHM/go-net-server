//二进制消息构造工具
package message

import (
	"encoding/binary"
)

type Msg struct {
	buffer []byte
}

func NewMessage() *Msg {
	this := new(Msg)
	this.buffer = make([]byte, 0, 50)
	return this
}

//向消息缓冲区写入一个字节
func (this *Msg) WriteUnsignedChar(uchar byte) {
	this.buffer = append(this.buffer, uchar)
}

//向消息缓冲区写入大端字节序的uint16
func (this *Msg) WriteBigEndianUint16(u16 uint16) {
	buf := []byte{0, 0}
	binary.BigEndian.PutUint16(buf[0:2], u16)
	this.buffer = append(this.buffer, buf...)
}

//向消息缓冲区写入大端字节序的uint32
func (this *Msg) WriteBigEndianUint32(u32 uint32) {
	buf := []byte{0, 0, 0, 0}
	binary.BigEndian.PutUint32(buf[0:2], u32)
	this.buffer = append(this.buffer, buf...)
}

//向消息缓冲区写入大端字节序的uint64
func (this *Msg) WriteBigEndianUint64(u64 uint64) {
	buf := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	binary.BigEndian.PutUint64(buf[0:2], u64)
	this.buffer = append(this.buffer, buf...)
}

//向消息缓冲区写入小端字节序的uint16
func (this *Msg) WriteLittleEndianUint16(u16 uint16) {
	buf := []byte{0, 0}
	binary.LittleEndian.PutUint16(buf[0:2], u16)
	this.buffer = append(this.buffer, buf...)
}

//向消息缓冲区写入小端字节序的uint32
func (this *Msg) WriteLittleEndianUint32(u32 uint32) {
	buf := []byte{0, 0, 0, 0}
	binary.LittleEndian.PutUint32(buf[0:2], u32)
	this.buffer = append(this.buffer, buf...)
}

//向消息缓冲区写入小端字节序的uint64
func (this *Msg) WriteLittleEndianUint64(u64 uint64) {
	buf := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	binary.LittleEndian.PutUint64(buf[0:2], u64)
	this.buffer = append(this.buffer, buf...)
}

//向消息缓冲区写入buffer
func (this *Msg) WriteBuffer(buf []byte) {
	this.buffer = append(this.buffer, buf...)
}

//得到最终的二进制缓存
func (this *Msg) GetBuffer() []byte {
	return this.buffer
}
