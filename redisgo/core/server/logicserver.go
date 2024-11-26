package server

import (
	"context"
	"net"
	"redisgo/core/db"
	"redisgo/core/parser"
)

type LogicServer struct {
}

func NewLogicServer() *LogicServer {
	return &LogicServer{}
}

func (s *LogicServer) Handle(ctx context.Context, conn net.Conn) {
	//1,解析数据
	parser.ParserStream(ctx, conn)

	//2,处理命令
	db.Exec()

	//3,将db处理结果写回客户端
	conn.Write([]byte("+Ok\r\n"))
}

func (s *LogicServer) Close() error {
	return nil
}
