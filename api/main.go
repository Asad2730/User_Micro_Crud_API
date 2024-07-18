package main

import (
	"log"

	"github.com/Asad2730/User_Micro_Crud_API/api/controllers"
	"github.com/Asad2730/User_Micro_Crud_API/api/routes"
	pb "github.com/Asad2730/User_Micro_Crud_API/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	grpcClient := pb.NewUserServiceClient(conn)

	userController := controllers.NewUserClient(grpcClient)

	r := gin.Default()
	routes.RegisterUserRoutes(r, userController)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
