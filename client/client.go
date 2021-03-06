package main

import (
	"context"
	"log"
	"time"

	pb "github.com/uocxp/grpc_tut/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	var new_users = make(map[string]int32)

	new_users["Achim"] = 43
	new_users["Wolfgang"] = 42

	for name, age := range new_users {
		r, err := client.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}
		log.Printf(`User Details:
		NAME: %s
		AGE: %d
		ID: %d`, r.GetName(), r.GetAge(), r.GetId())
	}

}
