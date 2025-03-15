package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"pingpawn.com/gotag/gemini"

	pb "pingpawn.com/gotag/protos/tagger"
)

var port = flag.Int("port", 50051, "The server port")

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.TaggerServer
}

// make a function that checks if there's a file named foo in the current working dir, and walk up a directory and check again if not.  return the relative path when found, error if not
func findFile(filename string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		path := filepath.Join(wd, filename)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}

		// Walk up a directory
		wd = filepath.Dir(wd)
		if wd == "/" || wd == "." {
			break
		}
	}

	return "", fmt.Errorf("file %s not found", filename)
}

func init() {
	pathToDotEnv, err := findFile(".env")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Found .env file at: %s", pathToDotEnv)
	}

	// check if a file named ".env" exists in the current working dir
	if _, err := os.Stat(pathToDotEnv); err == nil {
		if err := godotenv.Load(pathToDotEnv); err != nil {
			log.Printf("Warning: Error loading .env file: %v", err)
		} else {
			log.Printf("Loaded .env file from: %s", ".env")
		}
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("init - GEMINI_API_KEY environment variable not set")
	}
}

// implements tagger.TagText
func (s *server) TagText(_ context.Context, in *pb.UnstructuredText) (*pb.TagReply, error) {
	text := in.GetUnstructuredEntry()
	log.Printf("Received: %v", text)

	// Check for empty input
	if strings.TrimSpace(text) == "" {
		return nil, status.Error(codes.InvalidArgument, "empty input")
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, status.Error(codes.FailedPrecondition, "GEMINI_API_KEY environment variable not set")
	}

	client := gemini.NewClient(apiKey)
	query := "Return ONLY a JSON array of 2-5 tags that describe this text. No other text or formatting: \n\n" + text

	selectedTagsJsonStr, err := client.GenerateContent(query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate content: %v", err)
	}

	// Clean up the response - remove any markdown formatting
	cleanedStr := selectedTagsJsonStr
	re := regexp.MustCompile("(?s)^```(?:json)?\\s*(.+?)```\\s*$")
	if matches := re.FindStringSubmatch(cleanedStr); len(matches) > 1 {
		cleanedStr = matches[1]
	}
	cleanedStr = strings.TrimSpace(cleanedStr)

	var selectedTags []string
	if err := json.Unmarshal([]byte(cleanedStr), &selectedTags); err != nil {
		fmt.Fprintf(os.Stderr, "Raw Gemini API response:\n%s\n", selectedTagsJsonStr)
		return nil, status.Errorf(codes.Internal, "failed to parse tags from LLM response: %v", err)
	}

	return &pb.TagReply{Tags: selectedTags}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("main()")
	}

	s := grpc.NewServer()
	pb.RegisterTaggerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil { // s.Serve() is blocking
		log.Fatalf("failed to serve: %v", err)
	}
}
