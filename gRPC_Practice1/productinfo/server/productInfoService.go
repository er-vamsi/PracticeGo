package main

import (
	"context"

	pb "productinfo/server/ecommerce"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	productMap map[string]*pb.Product
}

func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductId, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating ProductId", err)
	}
	in.Id = out.String()
	s.productMap[in.Id] = in
	return &pb.ProductId{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *server) GetProduct(ctx conext.Context, in *pb.ProductId) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist..", in.Value)
}
