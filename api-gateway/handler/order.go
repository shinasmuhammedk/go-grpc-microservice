package handler

import (
	pb "api-gateway/proto"
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

type OrderHandler struct{
    client pb.OrderServiceClient
}

func RegisterOrderRoutes(r *gin.Engine, client pb.OrderServiceClient){
    h := &OrderHandler{client}
    
    r.POST("/order", h.CreateOrder)
    r.GET("/order/:id", h.GetOrder)
}

func (h *OrderHandler) CreateOrder(c *gin.Context){
    var req struct{
        UserId int32 `json:user_id`
        Product string `json:"product"`
    }
    
    c.BindJSON(&req)
    
    //token from header
    token := c.GetHeader("Authorization")
    
    //send token as metadata
    md := metadata.New(map[string]string{
        "authorization":token,
    })
    
    ctx := metadata.NewOutgoingContext(context.Background(), md)
    
    res,err := h.client.CreateOrder(ctx, &pb.CreateOrderRequest{
        UserId: req.UserId,
        Product: req.Product,
    })
    
    if err != nil{
        c.JSON(401, gin.H{
            "error":err.Error(),
        })
        return
    }
    
    c.JSON(200, res)
}


//Get order
func (h *OrderHandler) GetOrder(c *gin.Context){
    
    idParam := c.Param("id")
    id,_ := strconv.Atoi(idParam)
    
    token := c.GetHeader("Authorization")
    
    md := metadata.New(map[string]string{
        "authorization":token,
    })
    
    ctx := metadata.NewOutgoingContext(context.Background(), md)
    
    res, err := h.client.GetOrder(ctx, &pb.GetOrderRequest{
        Id: int32(id),
    })
    
    if err != nil{
        c.JSON(401, gin.H{
            "error":err.Error(),
        })
        return
    }
    
    c.JSON(200, res)
}