package main_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/mugund10/grpcuserservice/pb"
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

	c := pb.NewUserserviceClient(conn)

	req := pb.Requestid{
		Id: 2,
	}

	resp, err := c.Fetchuser(ctx, &req)
	errcheck(err)
	fmt.Println(resp)




one := pb.Requestid{Id: 3,}
two := pb.Requestid{Id: 7,}
three := pb.Requestid{Id: 9,}
four := pb.Requestid{Id: 12,}
multipleids := []*pb.Requestid{&one,&two,&three,&four}

	req2 := pb.Requestids{
		Id: multipleids,
	}

	ctx2, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp2, err := c.Fetchusers(ctx2, &req2)
	errcheck(err)
	fmt.Println(resp2)


}

func errcheck(err error) {
	if err != nil {
		log.Fatalf("error : %s", err)
	}
}
