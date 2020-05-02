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
	return response.Users,nil

}

func GetSingleUser(ctx context.Context, uid string)(*proto.User, error){

	conn, err := grpc.Dial("localhost:8005", grpc.WithInsecure())
	client := proto.NewAddServiceClient(conn)

	if (err != nil){
		fmt.Println(err)
	}
	
	req := &proto.GetSingleUserRequest{Uid: uid}
	response, err := client.GetSingleUser(ctx, req)
	return response, nil
	
}
