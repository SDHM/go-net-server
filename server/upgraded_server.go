package server

import "github.com/SDHM/conn"

type UpGradeServer struct {
	TcpServer
}

func NewUpGradeServer() *UpGradeServer {
	this := new(UpGradeServer)
	this.TcpServer.tcpServerIFace = this
	return this
}

//客户端权限鉴定，默认空实现
func (this *UpGradeServer) authLogin(connection *conn.Conn) error {
	return nil
}

func (this *UpGradeServer) handleMessage(connection *conn.Conn) error {
	
	return nil
}
