package main

import (
	"log"
	"net"
	"order-service/db"
	"order-service/interceptor"
	pb "order-service/proto"
	"order-service/service"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}

	dbConn := db.ConnectDB()

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.AuthInterceptor()),
	)

	pb.RegisterOrderServiceServer(grpcServer, &service.OrderService{
		DB: dbConn,
	})

	log.Println("Order service running on :50052")

	grpcServer.Serve(lis)
}
