package service

import (
	"context"
	"database/sql"
	pb "order-service/proto"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	DB *sql.DB
}

// Create order
func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {

	var id int32
	err := s.DB.QueryRow(
		"INSERT INTO grpcOrdersss(user_id,product) VALUES ($1,$2) RETURNING id",
		&req.UserId, req.Product,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &pb.OrderResponse{
		Id:      id,
		UserId:  req.UserId,
		Product: req.Product,
	}, nil
}

// Get order
func (s *OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {

	var order pb.OrderResponse

	err := s.DB.QueryRow(
		"SELECT id,user_id,product FROM grpcOrdersss WHERE id = $1",
		req.Id,
	).Scan(&order.Id, &order.UserId, &order.Product)

	if err != nil {
		return nil, err
	}

	return &order, nil
}
