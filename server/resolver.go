package main

import (
	"context"

	prisma "github.com/tonymj76/DBB/generated/prisma-client"
)

type Resolver struct {
	Prisma *prisma.Client
}

func (r *Resolver) Address() AddressResolver {
	return &addressResolver{r}
}
func (r *Resolver) Image() ImageResolver {
	return &imageResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type addressResolver struct{ *Resolver }

func (r *addressResolver) UserAddress(ctx context.Context, obj *prisma.Address) (*prisma.User, error) {
	return r.Prisma.Address(prisma.AddressWhereUniqueInput{ID: &obj.ID}).UserAddress().Exec(ctx)
}

type imageResolver struct{ *Resolver }

func (r *imageResolver) UserImage(ctx context.Context, obj *prisma.Image) (*prisma.User, error) {
	return r.Prisma.Image(prisma.ImageWhereUniqueInput{ID: &obj.ID}).UserImage().Exec(ctx)
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, firstName string, lastName string, email string, userName string, isAdmin bool, gender string) (*prisma.User, error) {
	return r.Prisma.CreateUser(prisma.UserCreateInput{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		UserName:  userName,
		IsAdmin:   &isAdmin,
		Gender:    gender,
	}).Exec(ctx)
}
func (r *mutationResolver) CreateImage(ctx context.Context, name string, userImage string) (*prisma.Image, error) {
	return r.Prisma.CreateImage(prisma.ImageCreateInput{
		Name: name,
		UserImage: &prisma.UserCreateOneWithoutImagesInput{
			Connect: &prisma.UserWhereUniqueInput{ID: &userImage},
		},
	}).Exec(ctx)
}
func (r *mutationResolver) CreateAddress(ctx context.Context, name string, city string, state string, country *string, userAddress string) (*prisma.Address, error) {
	return r.Prisma.CreateAddress(prisma.AddressCreateInput{
		AddressName: name,
		State:       state,
		Country:     country,
		City:        &city,
		UserAddress: &prisma.UserCreateOneWithoutAddressInput{
			Connect: &prisma.UserWhereUniqueInput{ID: &userAddress},
		},
	}).Exec(ctx)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]prisma.User, error) {
	return r.Prisma.Users(nil).Exec(ctx)
}
func (r *queryResolver) User(ctx context.Context, ID string) (*prisma.User, error) {
	return r.Prisma.User(prisma.UserWhereUniqueInput{ID: &ID}).Exec(ctx)
}

type userResolver struct{ *Resolver }

func (r *userResolver) Images(ctx context.Context, obj *prisma.User) ([]prisma.Image, error) {
	return r.Prisma.User(prisma.UserWhereUniqueInput{ID: &obj.ID}).Images(nil).Exec(ctx)
}
func (r *userResolver) Address(ctx context.Context, obj *prisma.User) (*prisma.Address, error) {
	return r.Prisma.User(prisma.UserWhereUniqueInput{ID: &obj.ID}).Address().Exec(ctx)
}
