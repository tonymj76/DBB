package main

import (
	"context"

	"github.com/tonymj76/DBB/generated/prisma-client"
)

type Resolver struct{
	Prisma		*prisma.Client
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}


type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, firstName string) (*prisma.User, error) {
	return r.Prisma.CreateUser(prisma.UserCreateInput{
		FirstName: firstName,
	}).Exec(ctx)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]prisma.User, error) {
	return r.Prisma.Users(nil).Exec(ctx)
}
func (r *queryResolver) User(ctx context.Context, ID string) (*prisma.User, error) {
	return r.Prisma.User(prisma.UserWhereUniqueInput{ID: &ID}).Exec(ctx)
}
