package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/pravandkatyare/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm client")

	cc, err := grpc.Dial("localhost:50050", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doBiDirectionalStreaming(c)
}

func doBiDirectionalStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting a BiDirectional streaming RPC")

	// Create a stream by nvoking the client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Pravand1",
				Lastname:  "Katyare1",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Pravand2",
				Lastname:  "Katyare2",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Pravand3",
				Lastname:  "Katyare3",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Pravand4",
				Lastname:  "Katyare4",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Pravand5",
				Lastname:  "Katyare5",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Pravand6",
				Lastname:  "Katyare6",
			},
		},
	}

	waitc := make(chan struct{})

	// Send a bunch of messages to the server (go routine)
	go func() {
		// function to send messages
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// Receive a bunch of messages from server (go routine)
	go func() {
		// function to receive messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Received: %v", res.GetResult())
		}
		close(waitc)
	}()

	// Block until everything is done
	<-waitc
}
