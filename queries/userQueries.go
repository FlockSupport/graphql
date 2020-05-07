package queries

import (
	"context"
	"flock-support/back/proto"
	"google.golang.org/grpc"
	"github.com/pkg/errors"
)


func GetAllUsers(ctx context.Context) ([]*proto.User , error){

	conn, err := grpc.Dial("localhost:8005", grpc.WithInsecure())
	if err != nil{
		return nil, errors.Wrap(err, "GraphQL: Unable to connect to port 8005")
	}
	
	client := proto.NewAddServiceClient(conn)

	req := &proto.GetAllUsersRequest{}
	response, err := client.GetAllUsers(ctx, req); 
	if (err != nil){
		return nil, err
	} else {
		return response.Users,nil
	}
	

}

func GetSingleUser(ctx context.Context, uid string)(*proto.User, error){

	conn, err := grpc.Dial("localhost:8005", grpc.WithInsecure())
	client := proto.NewAddServiceClient(conn)

	if (err != nil){
		return nil, errors.Wrap(err, "GraphQL: Unable to connect to port 8005")
	}
	
	req := &proto.GetSingleUserRequest{Uid: uid}
	response, err := client.GetSingleUser(ctx, req)
	if (err != nil){
		return nil, err
	} else {
		return response, nil
	}
	
}
