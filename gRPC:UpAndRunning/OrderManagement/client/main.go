package main

import (
	pb "OrderManagement/client/ecommerce"
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dail(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Not able to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background, time.Second*5)
	defer cancel()

	//Get Order
	order, err := client.GetOrder(ctx, &wrapper.StringValue{Value: "106"})
	if err != nil {
		log.Fatalf("Error retriving order: %v", err)
	}
	log.Print("GetOrder Response -> : ", retrievedOrder)

	//Search Order
	searchStream, _ := client.SearchOrder(ctx, &wrapper.StringValue{Value: "Google"})
	for {
		searchOrder, err := searchStream.Recv()
		if err == io.EOF {
			break
		}
		log.Print("Search result: ", searchOrder)
	}

}
