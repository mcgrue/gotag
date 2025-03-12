package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"
	// pb "pingpawn.com/gotag/protos/helloworld"
	pb "pingpawn.com/gotag/protos/tagger"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

var randomWords = [10]string{
	"cat",
	"sun",
	"blue",
	"jump",
	"quick",
	"fox",
	"lazy",
	"dog",
	"moon",
	"tree",
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.TaggerServer
}

// implements tagger.TagText
func (s *server) TagText(_ context.Context, in *pb.UnstructuredText) (*pb.TagReply, error) {
	log.Printf("Received: %v", in.GetUnstructuredEntry())

	log.Printf("Received: %v", in.GetUnstructuredEntry())

	// Shuffle the randomWords slice
	for i := len(randomWords) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		randomWords[i], randomWords[j] = randomWords[j], randomWords[i]
	}

	// Slice 2-5 entries from the shuffled slice
	numTags := rand.Intn(4) + 2 // Random number between 2 and 5
	selectedTags := randomWords[:numTags]

	// Create the TagReply with the selected tags
	reply := &pb.TagReply{
		Tags: selectedTags,
	}

	// set reply to an empty array
	// reply.Tags = []string{}

	return &pb.TagReply{Tags: reply.Tags}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTaggerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
