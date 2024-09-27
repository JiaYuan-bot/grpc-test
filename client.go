package main

import (
	"context"
	"fmt"
	"gprc-test/idl"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var GlobalData []int64

const maxMsgSize = 600 * 1024 * 1024

type client struct {
	hb   idl.HeartbeatServiceClient
	conn *grpc.ClientConn
}

func NewClient() (c *client, err error) {
	c = &client{}
	// gRPC heartbeat client
	if c.conn, err = grpc.NewClient(":9991", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(maxMsgSize),
		grpc.MaxCallSendMsgSize(maxMsgSize),
	)); err != nil {
		fmt.Printf("Failed to setup conn, addr=%s\n", "node")
		return
	}
	c.hb = idl.NewHeartbeatServiceClient(c.conn)
	return
}

func sendHeartbeat(client idl.HeartbeatServiceClient) {
	for {
		// Send the heartbeat every 2 seconds
		time.Sleep(10)

		start := time.Now()
		req := &idl.HeartbeatReq{Data: GlobalData}
		_, err := client.Heartbeat(context.Background(), req)
		if err != nil {
			fmt.Printf("Failed to send heartbeat: %v\n", err)
			continue
		}
		fmt.Printf("send time is %v\n", time.Since(start))
	}
}

func (c *client) StartHeartbeat() {
	go sendHeartbeat(c.hb)
}
