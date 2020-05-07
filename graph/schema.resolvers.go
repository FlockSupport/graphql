package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"flock-support/graphql/mutations"
	"flock-support/graphql/queries"

	"github.com/FlockSupport/graphql/graph/generated"
	"github.com/FlockSupport/graphql/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	createdUser, err := mutations.CreateUser(context.Background(), int64(input.Age), input.Email, input.UID)

	if err != nil {
		return nil, err
	}

	user := &model.User{ID: int(createdUser.Id), Age: int(createdUser.Age), Email: createdUser.Email, UID: createdUser.Uid}

	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	response, err := queries.GetAllUsers(context.Background())
	var users []*model.User // an empty list

	if err != nil {
		return nil, err
	}
	for _, record := range response {
		users = append(users, &model.User{ID: int(record.Id), Age: int(record.Age), Email: record.Email, UID: record.Uid})
	}

	return users, nil
}

func (r *queryResolver) SingleUser(ctx context.Context, input model.UIDInput) (*model.User, error) {
	response, err := queries.GetSingleUser(context.Background(), input.UID)

	if err != nil {
		return nil, err
	}

	user := &model.User{ID: int(response.Id), Age: int(response.Age), Email: response.Email, UID: response.Uid}

	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
