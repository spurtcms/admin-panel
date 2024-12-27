package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"spurt-cms/graphql/controller"
	"spurt-cms/graphql/graph"
	"spurt-cms/graphql/model"
)

// CategoryList is the resolver for the CategoryList field.
func (r *queryResolver) CategoryList(ctx context.Context, categoryFilter *model.CategoryFilter, commonFilter *model.Filter) (*model.CategoryDetails, error) {
	return controller.CategoryList(ctx, categoryFilter, commonFilter)
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
