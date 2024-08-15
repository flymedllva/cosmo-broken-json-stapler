package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"example/service1/graph/model"
)

// Test is the resolver for the test field.
func (r *mutationResolver) Test(ctx context.Context) (*model.TestResult, error) {
	return &model.TestResult{Data: "\u0015"}, nil
}

// Test is the resolver for the test field.
func (r *queryResolver) Test(ctx context.Context) (*model.TestResult, error) {
	return &model.TestResult{Data: "\u0015"}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
