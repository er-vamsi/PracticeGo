package main

import (
	"strings"
	pb "OrderManagement/service/ecommerce"
	"context"
	"net"

	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc"
)

const(
	port = ":8080"
)

type server struct {
	orderMap map[string]*pb.Order
}

func (s *server) GetOrder(ctx context.Context, orderID *wrapper.StingValue) (*pb.Order, error) {
	order, exists := s.orderMap[orderID.Value]
	if exists {
		return &order, status.New(codes.OK, "").Err()
	}
	return nil, status.Errof(codes.NotFound, "Order does not exist.", orderID)
}

func (s *server) SearchOrder(searchQuery *wrapper.StringValue, stream pb.OrderManagement_SearchOrderServer) error{
	for key, order := range s.orderMap{
		log.Print(key, order)
		for _,itemStr := range order.Items{
			if strings.Contains(itemStr,searchQuery.Value){
				err := stream.Send(&order)
				if err != nil {
					return fmt.Errorf("Error sending message to stream: %v",err)
				}
				log.Print("Matching order found... " + key)
				break
			}
		}
	}
	return nil
}

func main() {
	conn, err := net.Listen("tcp",port)
	if err := nil {
		lof.Fatalf("Failed to listen: %v",err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s,&server{})
	log.Printf("Server starting on port: %v",port)
	if err := s.Serve(conn); err != nil {
		log.Fatalf("Failed to start: %v",err)
	}
		
}
