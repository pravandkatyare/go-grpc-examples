package main

import (
	"context"
	"fmt"
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

	doClientStreaming(c)
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting a client streaming RPC")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Pravand1",
				Lastname:  "Katyare1",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Pravand2",
				Lastname:  "Katyare2",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Pravand3",
				Lastname:  "Katyare3",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Pravand4",
				Lastname:  "Katyare4",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Pravand5",
				Lastname:  "Katyare5",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Pravand6",
				Lastname:  "Katyare6",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while caling LongGreet RPC: %v", err)
	}

	// we iterate over slice and send each message individually

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet RPC: %v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res)
}
