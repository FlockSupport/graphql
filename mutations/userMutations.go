package mutations


import (
	"context"
	"flock-support/back/proto"
	"google.golang.org/grpc"
)


func CreateUser(ctx context.Context, age int64, email string, uid string) (*proto.User, error){
	conn, err := grpc.Dial("localhost:8005", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := proto.NewAddServiceClient(conn)

	req := &proto.AddUserRequest{Age: int64(age), Email: email, Uid: uid}
	response, err := client.AddUser(ctx, req)
	return response, nil

}