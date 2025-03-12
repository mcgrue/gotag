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

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"pingpawn.com/gotag/gemini"

	pb "pingpawn.com/gotag/protos/tagger"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

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
	log.Printf("Received: %v", in.GetUnstructuredEntry())

	// // Shuffle the randomWords slice
	// for i := len(randomWords) - 1; i > 0; i-- {
	// 	j := rand.Intn(i + 1)
	// 	randomWords[i], randomWords[j] = randomWords[j], randomWords[i]
	// }

	// // Slice 2-5 entries from the shuffled slice
	// numTags := rand.Intn(4) + 2 // Random number between 2 and 5
	// selectedTags := randomWords[:numTags]

	// use the gemini client to get the tags
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("TagText GEMINI_API_KEY environment variable not set")
	}
	client := gemini.NewClient(apiKey)

	query := "Please analyze the following text and provide a list of 2-5 tags that best describe the content. The tags should be relevant and concise.  Format the tags as a json array.\n\n\n\n" + in.GetUnstructuredEntry()

	selectedTagsJsonStr, err := client.GenerateContent(query)
	if err != nil {
		log.Fatalf("could not generate content: %v", err)
	}

	// using regexp, remove the ```json
	re := regexp.MustCompile("^```json")
	cleanedStr := re.ReplaceAllString(selectedTagsJsonStr, "$1")

	re = regexp.MustCompile("\n```\n")
	cleanedStr = re.ReplaceAllString(cleanedStr, "$1")

	// parse selectedTagsJsonStr as json
	selectedTags := []string{}
	if err := json.Unmarshal([]byte(cleanedStr), &selectedTags); err != nil {
		log.Fatalf("could not unmarshal json: '%v'", err)
		log.Fatalf("Original response: '%v'", selectedTagsJsonStr)
	}
	log.Printf("Unmarshalled json: %v", selectedTags)

	// Create the TagReply with the selected tags
	// reply := &pb.TagReply{
	// 	Tags: selectedTags,
	// }

	// set reply to an empty string array
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
