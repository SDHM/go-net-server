package conn

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

type Conn struct {
	closed    bool          //连接是否已经关闭
	rb        *bufio.Reader //buffer Reader
	wb        io.Writer     //数据写入句柄
	conn      net.Conn      //tcp连接
	timestart time.Time
}

func NewConn(c net.Conn) *Conn {
	this := new(Conn)
	this.conn = c
	this.rb = bufio.NewReaderSize(c, 1024)
	this.wb = c
	this.closed = false
	this.SetReadDeadline(time.Now().Add(time.Second * 30))
	this.timestart = time.Now()
	return this
}

//设置读取超时时间
func (this *Conn) SetReadDeadline(t time.Time) {
	this.conn.SetReadDeadline(t)
}

//关闭连接
func (this *Conn) Close() {

	if !this.closed {
		this.closed = true
		this.conn.Close()
	}
}

//发送二进制消息
func (this *Conn) SendMessage(data []byte) error {
	if n, err := this.wb.Write(data); err != nil {
		return err
	} else if n != len(data) {
		return err
	} else {
		return nil
	}
}

//读取一字节长度，并转化为uint8
func (this *Conn) ReadUnsignedChar() (uint8, error) {
	buf := []byte{0}
	if _, err := io.ReadFull(this.rb, buf); err != nil {
		return 0, err
	}

	return uint8(buf[0]), nil
}

//读取两字节长度，并转化为大端字节序的uint16
func (this *Conn) ReadBigEndianUint16() (uint16, error) {
	buf := []byte{0, 0}
	if _, err := io.ReadFull(this.rb, buf); err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint16(buf), nil
}

//读取两字节长度，并转化为小端字节序的uint16
func (this *Conn) ReadLittleEndianUint16() (uint16, error) {
	buf := []byte{0, 0}
	if _, err := io.ReadFull(this.rb, buf); err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint16(buf), nil
}

//读取两字节长度，并转化为大端字节序的uint32
func (this *Conn) ReadBigEndianUint32() (uint32, error) {
	buf := []byte{0, 0, 0, 0}
	if _, err := io.ReadFull(this.rb, buf); err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint32(buf), nil
}

//读取两字节长度，并转化为小端字节序的uint32
func (this *Conn) ReadLittleEndianUint32() (uint32, error) {
	buf := []byte{0, 0, 0, 0}
	if _, err := io.ReadFull(this.rb, buf); err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint32(buf), nil
}

//读取两字节长度，并转化为大端字节序的uint64
func (this *Conn) ReadBigEndianUint64() (uint64, error) {
	buf := []byte{0, 0, 0, 0}
	if _, err := io.ReadFull(this.rb, buf); err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint64(buf), nil
}

//读取两字节长度，并转化为小端字节序的uint64
func (this *Conn) ReadLittleEndianUint64() (uint64, error) {
	buf := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	if _, err := io.ReadFull(this.rb, buf); err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint64(buf), nil
}

//读取长度值为length的二进制数据
func (this *Conn) ReadMessage(length int) ([]byte, error) {
	buf := make([]byte, length)

	if _, err := io.ReadFull(this.rb, buf); err != nil {
		return nil, err
	}

	return buf, nil
}

func (this *Conn) PrintWaitTime() {
	fmt.Println("wait time:%d", time.Now().Sub(this.timestart))
}
