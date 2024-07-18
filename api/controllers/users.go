package controllers

import (
	"context"

	pb "github.com/Asad2730/User_Micro_Crud_API/proto"
	"github.com/gin-gonic/gin"
)

type UserClient struct {
	grpcClient pb.UserServiceClient
}

func NewUserClient(client pb.UserServiceClient) *UserClient {
	return &UserClient{
		grpcClient: client,
	}
}

func (cleint *UserClient) CreateUser(c *gin.Context) {
	var user pb.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	res, err := cleint.grpcClient.CreateUser(context.Background(), &user)
	if err != nil {
		c.JSON(402, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, res)
}

func (cleint *UserClient) GetUsers(c *gin.Context) {
	res, err := cleint.grpcClient.GetUsers(context.Background(), &pb.Empty{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res.List)
}

func (cleint *UserClient) GetUser(c *gin.Context) {
	id := c.Param("id")
	res, err := cleint.grpcClient.GetUser(context.Background(), &pb.UserId{Id: id})
	if err != nil {
		c.JSON(402, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

func (cleint *UserClient) UpdateUser(c *gin.Context) {
	var user pb.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	res, err := cleint.grpcClient.UpdateUser(context.Background(), &user)
	if err != nil {
		c.JSON(402, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

func (cleint *UserClient) RemoveUser(c *gin.Context) {
	id := c.Param("id")
	res, err := cleint.grpcClient.RemoveUser(context.Background(), &pb.UserId{Id: id})
	if err != nil {
		c.JSON(402, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

func (cleint *UserClient) StreamUsers(c *gin.Context) {

	res, err := cleint.grpcClient.StreamUsers(context.Background(), &pb.Empty{})
	if err != nil {
		c.JSON(402, gin.H{"error": err.Error()})
		return
	}

	var users []*pb.User
	for {
		user, err := res.Recv()
		if err != nil {
			break
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		c.JSON(404, gin.H{"error": "No users found"})
		return
	}

	c.JSON(200, users)
}
