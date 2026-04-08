package handler

import (

	pb "api-gateway/proto"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	client pb.UserServiceClient
}

func RegisterUserRoutes(r *gin.Engine, client pb.UserServiceClient) {
	h := &UserHandler{client}

	r.POST("/register", h.CreateUser)
	r.POST("/login", h.Login)
}

// Register
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	c.BindJSON(&req)

	res, err := h.client.CreateUser(c, &pb.CreateUserRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

// Login
func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	c.BindJSON(&req)

	res, err := h.client.Login(c, &pb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": res.Token})
}