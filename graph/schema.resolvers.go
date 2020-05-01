package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"flock-support/graphql/mutations"
	"flock-support/graphql/queries"
	"fmt"

	"github.com/FlockSupport/graphql/graph/generated"
	"github.com/FlockSupport/graphql/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	mutations.CreateUser(context.Background(), int64(input.ID), int64(input.Age), input.Name)

	user := &model.User{
		ID:   input.ID,
		Name: input.Name,
		Age:  input.Age,
	}

	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	response, err := queries.GetAllUsers(context.Background())
	var users []*model.User // an empty list

	if err != nil {
		fmt.Println(err)
	}
	for _, record := range response {
		fmt.Printf("Name: %s Age: %d\n", record.Name, record.Age)
		fmt.Println("")
		users = append(users, &model.User{ID: int(record.Id), Age: int(record.Age), Name: record.Name})
	}

	return users, nil
}

func (r *queryResolver) SingleUser(ctx context.Context, input model.IDInput) (*model.User, error) {
	response, err := queries.GetSingleUser(context.Background(), int64(input.ID))

	if err != nil {
		fmt.Println(err)
	}

	user := &model.User{ID: int(response.Id), Age: int(response.Age), Name: response.Name}

	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
