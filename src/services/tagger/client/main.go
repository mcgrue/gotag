package main

import (
	"bufio"
	"context"
	"flag"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "pingpawn.com/gotag/protos/tagger"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func readStdin() string {
	reader := bufio.NewReader(os.Stdin)
	var content string
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			content += line
			break
		}
		if err != nil {
			log.Fatalf("Error reading stdin: %v", err)
		}
		content += line
	}
	return content
}

func main() {
	flag.Parse()

	// Read content from stdin
	content := readStdin()

	// print the content
	log.Printf("Content: %s", content)

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTaggerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.TagText(ctx, &pb.UnstructuredText{UnstructuredEntry: content})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Tags: %s", strings.Join(r.GetTags(), ", "))
}
