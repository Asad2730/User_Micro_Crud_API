package services

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Asad2730/User_Micro_Crud_API/proto"
	"github.com/Asad2730/User_Micro_Crud_API/user/data"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) CreateUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	data.Users_Arr = append(data.Users_Arr, user)
	return user, nil
}

func (s *server) GetUsers(ctx context.Context, in *pb.Empty) (*pb.Users, error) {
	return &pb.Users{List: data.Users_Arr}, nil
}

func (s *server) GetUser(ctx context.Context, id *pb.UserId) (*pb.User, error) {

	for _, user := range data.Users_Arr {
		if user.Id == id.Id {
			return user, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "User not found with id: %s", id.Id)
}

func (s *server) UpdateUser(ctx context.Context, updateUserDto *pb.User) (*pb.User, error) {
	for i, user := range data.Users_Arr {
		if user.Id == updateUserDto.Id {
			data.Users_Arr[i] = updateUserDto
			return updateUserDto, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "User not found with id: %s", updateUserDto.Id)
}

func (s *server) RemoveUser(ctx context.Context, id *pb.UserId) (*pb.Empty, error) {

	for i, user := range data.Users_Arr {
		if user.Id == id.Id {
			data.Users_Arr = append(data.Users_Arr[:i], data.Users_Arr[i+1:]...)
			return &pb.Empty{}, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "User not found with id: %s", id.Id)
}

func (s *server) StreamUsers(empty *pb.Empty, stream pb.UserService_StreamUsersServer) error {

	for _, user := range data.Users_Arr {
		if err := stream.Send(user); err != nil {
			return err
		}
	}

	return nil
}

func Start() {
	listners, err := net.Listen("tcp", "5001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	fmt.Println("gRPC server is running at :50051")
	if err := s.Serve(listners); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
