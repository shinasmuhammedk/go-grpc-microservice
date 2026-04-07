package main

import (
	"api-gateway/handler"
	pb "api-gateway/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	//connect userservice
	userConn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
    userClient := pb.NewUserServiceClient(userConn)
    
    //connect orderservice
    orderConn,_ := grpc.Dial("localhost:50052", grpc.WithInsecure())
    orderClient := pb.NewOrderServiceClient(orderConn)
    
    r := gin.Default()
    
    //routes
    
}
