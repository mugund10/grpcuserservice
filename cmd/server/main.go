package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mugund10/grpcuserservice/internal/db"
	"github.com/mugund10/grpcuserservice/pb"
	"github.com/mugund10/grpcuserservice/pkg/anysolver"
	"github.com/mugund10/grpcuserservice/pkg/search"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// creates a tcp listener on port 9292
	addr := ":9292"
	lis, err := net.Listen("tcp", addr)
	iserr(err, "listen") // a function which checks for error

	// creates a new grpc server
	srv := grpc.NewServer()

	// registering grpc server
	var us Userser
	pb.RegisterUserserviceServer(srv, &us)

	// registering reflection service
	reflection.Register(srv)

	// now the server is ready
	log.Printf("info: server is ready on %s \n", addr)

	// serving request
	err = srv.Serve(lis)
	iserr(err, "serve") // a function which checks for error

}

func iserr(err error, s string) {
	if err != nil {
		log.Fatalf("error cant be %s on %s", s, err)
	}
}

type Userser struct {
	pb.UnimplementedUserserviceServer
}

// fetches user details based on a user id
func (u *Userser) Fetchuser(ctx context.Context, req *pb.Requestid) (*pb.User, error) {

	people := db.UsersList
	resp := pb.User{}
	//
	for _, value := range people {
		if value.Id == req.Id {
			resp = pb.User{
				Id:      value.Id,
				Fname:   value.Fname,
				City:    value.City,
				Phone:   value.Phone,
				Height:  value.Height,
				Married: value.Married,
			}
		}
	}
	//if the requested id is not prersent on the database then its returns a error
	if req.Id != resp.Id {
		return nil, fmt.Errorf("user %d is not found", req.Id)
	}
	log.Printf("FETCHUSER(endpoint): served user detail with id %d \n", req.Id)
	return &resp, nil
}

// retrieves a list of user details based on a list of user id's
func (u *Userser) Fetchusers(ctx context.Context, req *pb.Requestids) (*pb.Users, error) {
	people := db.UsersList
	usersli := []*pb.User{}

	for _, value1 := range req.Id {
		for _, value2 := range people {

			if value1.Id == value2.Id {
				usersli = append(usersli, &pb.User{
					Id:      value2.Id,
					Fname:   value2.Fname,
					City:    value2.City,
					Phone:   value2.Phone,
					Height:  value2.Height,
					Married: value2.Married,
				})
				break
			}
		}
	}
	log.Printf("FETCHUSERS(endpoint): served users details with ids %s \n", req.Id[:])

	resp := pb.Users{
		User: usersli,
	}

	return &resp, nil

}

func (u *Userser) Search(ctx context.Context, req *pb.Criteria) (*pb.Users, error) {

	users := db.UsersList
	search := search.Details{Users: users}
	usersli := []*pb.User{}

	field := req.Field
	value, err := anysolver.ConvertAnyToInterface(req.Value)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	muser := search.ByCriteria(field, value)
	log.Printf("SEARCH(endpoint): served users details with criteria where %s = %v:\n", field, value)
	for _, user := range muser {

		usersli = append(usersli, &pb.User{
			Id:      user.Id,
			Fname:   user.Fname,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		})
	}

	resp := pb.Users{
		User: usersli,
	}
	return &resp, nil
}
