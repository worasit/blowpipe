package server

import "google.golang.org/grpc"

type Config struct {
	Host string
	Port int
	Conn grpc.ClientConn
}

type Server struct {
	Config Config
}

func NewServer(Config config) *Server {
	server := Server()
	server.Config = server.Config
	return server
}

func (server Server) Status() string {

}

func (server Server) Connect() bool {
	conn, err := grpc.Dial(*serverAddress)
	server.Connection = conn
	defer conn.close()
	return true
}

func (server Server) Disconnect() string {
	defer server.conn.close()

}
