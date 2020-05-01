package queries

import (
	"context"
	"flock-support/back/proto"
	"fmt"
	"google.golang.org/grpc"
)


func GetAllUsers(ctx context.Context) ([]*proto.User , error){

	conn, err := grpc.Dial("localhost:8005", grpc.WithInsecure())
	if err != nil{
		fmt.Println(err)
	}
	
	client := proto.NewAddServiceClient(conn)

	req := &proto.GetAllUsersRequest{}
	response, err := client.GetAllUsers(ctx, req); 
	fmt.Println("result: ", response.Users)
	return response.Users,nil

}

func GetSingleUser(ctx context.Context, id int64)(*proto.User, error){

	conn, err := grpc.Dial("localhost:8005", grpc.WithInsecure())
	client := proto.NewAddServiceClient(conn)

	if (err != nil){
		fmt.Println(err)
	}
	
	req := &proto.GetSingleUserRequest{Id: int64(id)}
	response, err := client.GetSingleUser(ctx, req)
	fmt.Println("result: ", response)
	return response, nil
	
}
