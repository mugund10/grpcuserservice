package main_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/mugund10/grpcuserservice/pb"
	"github.com/mugund10/grpcuserservice/pkg/anysolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Test(t *testing.T) {
	addr := "localhost:9292"
	creds := insecure.NewCredentials()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(creds))
	errcheck(err)
	defer conn.Close()

	log.Printf("connected to %s", addr)

	client := pb.NewUserserviceClient(conn)

	req := pb.Requestid{
		Id: 2,
	}

	resp, err := client.Fetchuser(ctx, &req)
	errcheck(err)
	fmt.Println(resp)

	one := pb.Requestid{Id: 3}
	two := pb.Requestid{Id: 7}
	three := pb.Requestid{Id: 9}
	four := pb.Requestid{Id: 12}
	multipleids := []*pb.Requestid{&one, &two, &three, &four}

	req2 := pb.Requestids{
		Id: multipleids,
	}

	ctx2, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp2, err := client.Fetchusers(ctx2, &req2)
	errcheck(err)
	fmt.Println(resp2)

	value, err := anysolver.ConvertInterfaceToAny("MUMBAI")
	errcheck(err)

	

	searchRequest := &pb.Criteria{
		Field: "city",
		Value: value,
	}
	

	users, err := client.Search(context.Background(), searchRequest)
	if err != nil {
		log.Fatalf("Search request failed: %v", err)
	}

	// Process the response from the server.
	for _, user := range users.User {
		log.Printf("User ID: %d, First Name: %s, City: %s", user.Id, user.Fname, user.City)
	}

	value1, err := anysolver.ConvertInterfaceToAny(true)
	errcheck(err)

	

	searchRequest1 := &pb.Criteria{
		Field: "married",
		Value: value1, // Set ny as the value for the Any field
	}
	

	users1, err := client.Search(context.Background(), searchRequest1)
	if err != nil {
		log.Fatalf("Search request failed: %v", err)
	}

	// Process the response from the server.
	for _, user := range users1.User {
		log.Printf("User ID: %d, First Name: %s, City: %s , married: %t", user.Id, user.Fname, user.City, user.Married)
	}

}

func errcheck(err error) {
	if err != nil {
		log.Fatalf("error : %s", err)
	}
}
