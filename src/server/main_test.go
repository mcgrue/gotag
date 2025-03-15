package main

import (
	"context"
	"net"
	"regexp"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	pb "pingpawn.com/gotag/protos/tagger"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterTaggerServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestParseTags(t *testing.T) {
	input := "```json\n"
	input += "[\n"
	input += "  \"hunter2\",\n"
	input += "  \"password\",\n"
	input += "  \"joke\",\n"
	input += "  \"security\",\n"
	input += " \"IRC\"\n"
	input += "]\n"
	input += "```\n"
	input += "\n"

	// print out the strting
	t.Log(input)

	// using regexp, remove the ```json
	re := regexp.MustCompile("^```json")
	cleanedStr := re.ReplaceAllString(input, "$1")

	re2 := regexp.MustCompile("\n```\n")
	cleanedStr2 := re2.ReplaceAllString(cleanedStr, "$1")

	// print out the strting
	t.Log(cleanedStr2)
}

// this is a bad test, but it's a start
func TestTagText(t *testing.T) {
	// Increase timeout to 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewTaggerClient(conn)

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:  "Empty input",
			input: "",
			// Empty input should return an error
			wantErr: true,
		},
		{
			name:    "Single line input",
			input:   "Hello world",
			wantErr: false,
		},
		{
			name:    "Multi line input",
			input:   "Hello\nworld\ntest",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &pb.UnstructuredText{UnstructuredEntry: tt.input}
			resp, err := client.TagText(ctx, req)

			if (err != nil) != tt.wantErr {
				t.Errorf("TagText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && resp == nil {
				t.Error("Expected non-nil response")
				return
			}

			if !tt.wantErr {
				// Check that we get between 2 and 5 tags
				if len(resp.Tags) < 2 || len(resp.Tags) > 5 {
					t.Errorf("Expected between 2 and 5 tags, got %d", len(resp.Tags))
				}

				// Check that all tags are non-empty
				for i, tag := range resp.Tags {
					if tag == "" {
						t.Errorf("Tag at index %d is empty", i)
					}
				}

				// Check for duplicate tags
				seen := make(map[string]bool)
				for _, tag := range resp.Tags {
					if seen[tag] {
						t.Errorf("Duplicate tag found: %s", tag)
					}
					seen[tag] = true
				}
			}
		})
	}
}

// TestServerImplementation ensures the server implements the required interface
func TestServerImplementation(t *testing.T) {
	s := &server{}
	var _ pb.TaggerServer = s // Verify that server implements the interface
}
