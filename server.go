package main

import (
	"context"
	"fmt"
	"gprc-test/idl"
	"net"

	"google.golang.org/grpc"
)

type HeartBeater struct {
	idl.UnimplementedHeartbeatServiceServer
}

func NewHeartBeater() *HeartBeater {
	return &HeartBeater{}
}

// Heartbeat register worker
func (hb *HeartBeater) Heartbeat(ctx context.Context, req *idl.HeartbeatReq) (*idl.HeartbeatReply, error) {
	reply := &idl.HeartbeatReply{
		Errno: 0,
	}

	return reply, nil
}

func (hb *HeartBeater) Run() {
	addr := fmt.Sprintf(":%d", 9991)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("failed to run hb server, err=%v\n", err)
		panic(err)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.MaxRecvMsgSize(maxMsgSize))
	opts = append(opts, grpc.MaxSendMsgSize(maxMsgSize))

	server := grpc.NewServer(opts...)
	idl.RegisterHeartbeatServiceServer(server, hb)
	server.Serve(lis)
	return
}

func (hb *HeartBeater) Start() {
	go hb.Run()
}
