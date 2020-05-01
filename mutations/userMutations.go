package mutations


import (
	"context"
	"flock-support/back/proto"
	"fmt"
	"google.golang.org/grpc"
)

func CreateUser(ctx context.Context, id int64, age int64, name string) {
	fmt.Println("REACHED!")
	conn, err := grpc.Dial("localhost:8005", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	client := proto.NewAddServiceClient(conn)

	req := &proto.AddUserRequest{Id: int64(id), Age: int64(age), Name: string(name)}
	if response, err := client.AddUser(ctx, req); err == nil {
		fmt.Println("result: ", response.Result)
	} else {
		fmt.Println(err)
	}

}