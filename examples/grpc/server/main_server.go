package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"lark/examples/pb_auth"
	"net"
)

func main() {
	s := new(AuthServer)
	s.Run()
}

type AuthServer struct {
	pb_auth.UnimplementedAuthServer
}

func (s *AuthServer) Run() {
	var (
		listener net.Listener
		srv      *grpc.Server
		err      error
	)
	listener, err = net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	srv = grpc.NewServer()
	pb_auth.RegisterAuthServer(srv, s)

	err = srv.Serve(listener)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *AuthServer) SignUp(ctx context.Context, req *pb_auth.SignUpReq) (resp *pb_auth.SignUpResp, err error) {
	fmt.Println(req.Nickname)
	resp = new(pb_auth.SignUpResp)
	resp.Msg = "Sign up succeed"
	return
}
