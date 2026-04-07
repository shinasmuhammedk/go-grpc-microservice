package main

import (
	"log"
	"net"
	"user-service/db"
	"user-service/interceptor"
	pb "user-service/proto"
	"user-service/service"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	dbConn := db.ConnectDB()
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.AuthInterceptor()),
	)

	pb.RegisterUserServiceServer(grpcServer, &service.UserService{
		DB: dbConn,
	})

	log.Println("User servic running on :50051")
	grpcServer.Serve(lis)
}
