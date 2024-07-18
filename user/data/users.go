package data

import pb "github.com/Asad2730/User_Micro_Crud_API/proto"

var Users_Arr = []*pb.User{
	{Id: "1", Name: "Alice", Email: "alice@example.com", Degree: pb.Degree_BS},
	{Id: "2", Name: "Bob", Email: "bob@example.com", Degree: pb.Degree_MS},
}
