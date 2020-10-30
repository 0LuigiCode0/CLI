package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"test/hub"

	"google.golang.org/grpc"
)

const gaddr = ":50051"

type server struct {
	hub.HubServer
}

type JSON struct {
	Message string `json:"msg"`
	Success bool   `json:"success"`
}

func (s *server) RunRoute(ctx context.Context, req *hub.MessageJSON) (resp *hub.MessageJSON, err error) {
	resp = new(hub.MessageJSON)
	msg := new(JSON)
	err = json.Unmarshal(req.Body, msg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	msg.Success = false
	bmsg, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp.Body = []byte(bmsg)
	return resp, nil
}

func main() {

	glisten, err := net.Listen("tcp", gaddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	gServer := grpc.NewServer()
	hub.RegisterHubServer(gServer, &server{})
	err = gServer.Serve(glisten)
	if err != nil {
		fmt.Println(err)
		return
	}
}
