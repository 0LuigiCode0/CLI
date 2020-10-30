package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"test/hub"

	"github.com/gorilla/mux"

	"google.golang.org/grpc"
)

const addr = ":8080"
const gaddr = ":50051"

type JSON struct {
	Message string `json:"msg"`
	Success bool   `json:"success"`
}

type server struct {
	client hub.HubClient
}

func main() {
	srv := new(server)

	conn, err := grpc.Dial(gaddr, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	client := hub.NewHubClient(conn)
	srv.client = client

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	route := mux.NewRouter()
	route.HandleFunc("/", srv.Helloo)

	err = http.Serve(listen, route)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {

	}
}

func (s *server) Helloo(w http.ResponseWriter, r *http.Request) {

	data, _ := ioutil.ReadAll(r.Body)

	msg := new(hub.MessageJSON)
	msg.Body = data

	resp, err := s.client.RunRoute(context.TODO(), msg)
	if err != nil {
		fmt.Println(err)
		msgE := []byte(err.Error())
		w.Write(msgE)
		return
	}

	w.Write(resp.Body)
}
