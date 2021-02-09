package main

import (
	"context"
	"fmt"
	"log"

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

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			Firstname: "Pravand",
			Lastname:  "Katyare",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while caling greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
