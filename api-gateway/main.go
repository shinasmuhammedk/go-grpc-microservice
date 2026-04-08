package main

import (
	"api-gateway/handler"
	pb "api-gateway/proto"
	"log"

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
    handler.RegisterUserRoutes(r, userClient)
    handler.RegisterOrderRoutes(r, orderClient)
    
    
    log.Println("Gateway running on : 8080")
    r.Run(":8080")
}
