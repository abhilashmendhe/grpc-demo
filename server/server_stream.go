package main

import (
	"log"
	"time"

	pb "github.com/abhilashmendhe/grpc-demo/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetServer_SayHelloServerStreamingServer) error {
	log.Printf("Got request with names: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hi " + name,
		}
		if err := stream.SendMsg(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
