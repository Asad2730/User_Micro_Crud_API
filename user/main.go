package main

import (
	"log"

	"github.com/Asad2730/User_Micro_Crud_API/user/services"
)

func main() {
	grpcServer := services.NewServer(":8000")
	if err := grpcServer.Start(); err != nil {
		log.Fatalf("Failed to serve %v", err.Error())
	}
}
