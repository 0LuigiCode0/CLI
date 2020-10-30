package main

import (
	"context"
	"encoding/json"
	"fmt"
	"test/hub"
	"time"

	"google.golang.org/grpc"
)

const addr = ":50051"

type JSON struct {
	Message string `json:"msg"`
	Success bool   `json:"success"`
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	client := hub.NewHubClient(conn)
	time.Sleep(time.Second * 2)
	msg := new(JSON)
	msg.Message = "sdffsdgbdf"
	msg.Success = true
	bmsg, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	req := new(hub.MessageJSON)
	req.Body = []byte(bmsg)
	resp, err := client.RunRoute(context.TODO(), req)
	if err != nil {
		fmt.Println(err)
		return
	}
	out := new(JSON)
	err = json.Unmarshal(resp.Body, out)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)
	for {

	}
}
