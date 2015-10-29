package server

import (
	"fmt"
	"io"
	"net"
	"runtime"
	"strings"

	"github.com/SDHM/conn"
)

//tcp服务接口，抽象tcp服务的公共接口
type TcpServerInterface interface {
	listen(addr string) error                  //监听
	accept()                                   //accept客户端连接
	Run(addr string) error                     //运行tcp服务
	authLogin(connection *conn.Conn) error     //客户端权限鉴定
	onConnection(con net.Conn)                 //连接服务接口
	connectionRun(connection *conn.Conn)       //连接服务函数
	handleMessage(connection *conn.Conn) error //处理消息的接口
}

//tcp服务公共类
type TcpServer struct {
	tcpServerIFace TcpServerInterface
	closed         bool             //tcp服务是否关闭
	listener       *net.TCPListener //监听对象
}

//启动服务，包含监听和接受客户端连接
func (this *TcpServer) Run(addr string) error {

	if err := this.listen(addr); nil != err {
		return err
	}

	fmt.Println("listen addr:", addr)
	this.accept()

	return nil
}

//监听端口
func (this *TcpServer) listen(addr string) error {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if nil != err {
		return err
	}

	this.listener, err = net.ListenTCP("tcp", tcpAddr)

	if nil != err {
		return err
	}

	return nil
}

//接受客户端连接，并为每个连接启动协程
func (this *TcpServer) accept() {
	for {
		conn, err := this.listener.AcceptTCP()
		if err != nil {
			continue
		}
		conn.SetLinger(0)
		go this.onConnection(conn)
	}
}

//处理连接
func (this *TcpServer) onConnection(con net.Conn) {
	connection := conn.NewConn(con)
	this.connectionRun(connection)
}

//处理连接
func (this *TcpServer) connectionRun(connection *conn.Conn) {
	defer func() {
		if err := recover(); err != nil {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			fmt.Printf("Server run panic : %s\n%s", err, string(buf))
		}

		connection.Close()
	}()

	this.tcpServerIFace.authLogin(connection)

	for {
		if err := this.tcpServerIFace.handleMessage(connection); nil != err {
			if strings.EqualFold("client quit", err.Error()) {
				break
			} else if err == io.EOF {
				break
			}
		}
	}

}
