package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/FlockSupport/graphql/graph/generated"
	"github.com/FlockSupport/graphql/graph/model"
	"fmt"
	"flock-support/back/proto"
	"google.golang.org/grpc"
	"strconv"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:       input.ID,
		Name:     input.Name,
		Quantity: input.Quantity,
	}

	id, err := strconv.ParseInt(input.ID, 10, 64)
	
	if (err != nil){
		fmt.Println(err)		
	} else {
		createUser(context.Background(), id, int64(input.Quantity), input.Name)
	}
	
	r.users = append(r.users, user)
	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }


func createUser(ctx context.Context, id int64, quantity int64, name string){
	fmt.Println("REACHED!")
	conn, err := grpc.Dial("localhost:8005", grpc.WithInsecure())
	if err != nil{
		fmt.Println(err)
	}
	
	client := proto.NewAddServiceClient(conn)

		req := &proto.AddUserRequest{Id: int64(id), Quantity: int64(quantity), Name: string(name)}
		if response, err := client.AddUser(ctx, req); err == nil {
			fmt.Println("result: ", response.Result)
		} else {
			fmt.Println(err)
		}


}

