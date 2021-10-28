package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/uocxp/grpc_tut/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	user_name := in.GetName()
	user_age := in.GetAge()
	user_id := int32(rand.Intn(1000))
	log.Printf("Recieved: %v", user_name)
	return &pb.User{Name: user_name, Age: user_age, Id: user_id}, nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// creates a gRPC server which has no service registered and has not started to accept requests yet
	server := grpc.NewServer()
	// register the server as a new gRPC service
	pb.RegisterUserManagementServer(server, &UserManagementServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}
