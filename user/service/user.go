package service

import (
	"fmt"
	"log"
	"net"

	pb "github.com/Asad2730/User_Micro_Crud_API/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	users   []*pb.User
	userIdx int
}

func RunServer() {
	listners, err := net.Listen("tcp", "5001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	users := []*pb.User{
		{Id: "1", Name: "John Doe", Email: "john.doe@example.com", Degree: pb.Degree_BS},
		{Id: "2", Name: "Jane Smith", Email: "jane.smith@example.com", Degree: pb.Degree_MS},
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{users: users})
	fmt.Println("gRPC server is running at :50051")
	if err := s.Serve(listners); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
