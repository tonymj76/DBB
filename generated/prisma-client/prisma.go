// Code generated by Prisma CLI (https://github.com/prisma/prisma). DO NOT EDIT.

package prisma

import (
	"context"
	"errors"

	"os"

	"github.com/prisma/prisma-client-lib-go"

	"github.com/machinebox/graphql"
)

var ErrNoResult = errors.New("query returned no result")

func Str(v string) *string { return &v }
func Int32(v int32) *int32 { return &v }
func Bool(v bool) *bool    { return &v }

type BatchPayloadExec struct {
	exec *prisma.BatchPayloadExec
}

func (exec *BatchPayloadExec) Exec(ctx context.Context) (BatchPayload, error) {
	bp, err := exec.exec.Exec(ctx)
	return BatchPayload(bp), err
}

type BatchPayload struct {
	Count int64 `json:"count"`
}

type Aggregate struct {
	Count int64 `json:"count"`
}

type Client struct {
	Client *prisma.Client
}

type Options struct {
	Endpoint string
	Secret   string
}

func New(options *Options, opts ...graphql.ClientOption) *Client {
	endpoint := DefaultEndpoint
	secret := Secret
	if options != nil {
		endpoint = options.Endpoint
		secret = options.Secret
	}
	return &Client{
		Client: prisma.New(endpoint, secret, opts...),
	}
}

func (client *Client) GraphQL(ctx context.Context, query string, variables map[string]interface{}) (map[string]interface{}, error) {
	return client.Client.GraphQL(ctx, query, variables)
}

var DefaultEndpoint = "http://localhost:4467"
var Secret = os.Getenv("DBB_SECRET")

func (client *Client) User(params UserWhereUniqueInput) *UserExec {
	ret := client.Client.GetOne(
		nil,
		params,
		[2]string{"UserWhereUniqueInput!", "User"},
		"user",
		[]string{"id", "email", "firstName", "lastName", "userName", "gender", "isAdmin", "createdAt", "updatedAt"})

	return &UserExec{ret}
}

type UsersParams struct {
	Where   *UserWhereInput   `json:"where,omitempty"`
	OrderBy *UserOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

func (client *Client) Users(params *UsersParams) *UserExecArray {
	var wparams *prisma.WhereParams
	if params != nil {
		wparams = &prisma.WhereParams{
			Where:   params.Where,
			OrderBy: (*string)(params.OrderBy),
			Skip:    params.Skip,
			After:   params.After,
			Before:  params.Before,
			First:   params.First,
			Last:    params.Last,
		}
	}

	ret := client.Client.GetMany(
		nil,
		wparams,
		[3]string{"UserWhereInput", "UserOrderByInput", "User"},
		"users",
		[]string{"id", "email", "firstName", "lastName", "userName", "gender", "isAdmin", "createdAt", "updatedAt"})

	return &UserExecArray{ret}
}

type UsersConnectionParams struct {
	Where   *UserWhereInput   `json:"where,omitempty"`
	OrderBy *UserOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

func (client *Client) UsersConnection(params *UsersConnectionParams) UserConnectionExec {
	panic("not implemented")
}

func (client *Client) CreateUser(params UserCreateInput) *UserExec {
	ret := client.Client.Create(
		params,
		[2]string{"UserCreateInput!", "User"},
		"createUser",
		[]string{"id", "email", "firstName", "lastName", "userName", "gender", "isAdmin", "createdAt", "updatedAt"})

	return &UserExec{ret}
}

type UserUpdateParams struct {
	Data  UserUpdateInput      `json:"data"`
	Where UserWhereUniqueInput `json:"where"`
}

func (client *Client) UpdateUser(params UserUpdateParams) *UserExec {
	ret := client.Client.Update(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[3]string{"UserUpdateInput!", "UserWhereUniqueInput!", "User"},
		"updateUser",
		[]string{"id", "email", "firstName", "lastName", "userName", "gender", "isAdmin", "createdAt", "updatedAt"})

	return &UserExec{ret}
}

type UserUpdateManyParams struct {
	Data  UserUpdateManyMutationInput `json:"data"`
	Where *UserWhereInput             `json:"where,omitempty"`
}

func (client *Client) UpdateManyUsers(params UserUpdateManyParams) *BatchPayloadExec {
	exec := client.Client.UpdateMany(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[2]string{"UserUpdateManyMutationInput!", "UserWhereInput"},
		"updateManyUsers")
	return &BatchPayloadExec{exec}
}

type UserUpsertParams struct {
	Where  UserWhereUniqueInput `json:"where"`
	Create UserCreateInput      `json:"create"`
	Update UserUpdateInput      `json:"update"`
}

func (client *Client) UpsertUser(params UserUpsertParams) *UserExec {
	uparams := &prisma.UpsertParams{
		Where:  params.Where,
		Create: params.Create,
		Update: params.Update,
	}
	ret := client.Client.Upsert(
		uparams,
		[4]string{"UserWhereUniqueInput!", "UserCreateInput!", "UserUpdateInput!", "User"},
		"upsertUser",
		[]string{"id", "email", "firstName", "lastName", "userName", "gender", "isAdmin", "createdAt", "updatedAt"})

	return &UserExec{ret}
}

func (client *Client) DeleteUser(params UserWhereUniqueInput) *UserExec {
	ret := client.Client.Delete(
		params,
		[2]string{"UserWhereUniqueInput!", "User"},
		"deleteUser",
		[]string{"id", "email", "firstName", "lastName", "userName", "gender", "isAdmin", "createdAt", "updatedAt"})

	return &UserExec{ret}
}

func (client *Client) DeleteManyUsers(params *UserWhereInput) *BatchPayloadExec {
	exec := client.Client.DeleteMany(params, "UserWhereInput", "deleteManyUsers")
	return &BatchPayloadExec{exec}
}

type UserOrderByInput string

const (
	UserOrderByInputIDAsc         UserOrderByInput = "id_ASC"
	UserOrderByInputIDDesc        UserOrderByInput = "id_DESC"
	UserOrderByInputEmailAsc      UserOrderByInput = "email_ASC"
	UserOrderByInputEmailDesc     UserOrderByInput = "email_DESC"
	UserOrderByInputFirstNameAsc  UserOrderByInput = "firstName_ASC"
	UserOrderByInputFirstNameDesc UserOrderByInput = "firstName_DESC"
	UserOrderByInputLastNameAsc   UserOrderByInput = "lastName_ASC"
	UserOrderByInputLastNameDesc  UserOrderByInput = "lastName_DESC"
	UserOrderByInputUserNameAsc   UserOrderByInput = "userName_ASC"
	UserOrderByInputUserNameDesc  UserOrderByInput = "userName_DESC"
	UserOrderByInputGenderAsc     UserOrderByInput = "gender_ASC"
	UserOrderByInputGenderDesc    UserOrderByInput = "gender_DESC"
	UserOrderByInputIsAdminAsc    UserOrderByInput = "isAdmin_ASC"
	UserOrderByInputIsAdminDesc   UserOrderByInput = "isAdmin_DESC"
	UserOrderByInputCreatedAtAsc  UserOrderByInput = "createdAt_ASC"
	UserOrderByInputCreatedAtDesc UserOrderByInput = "createdAt_DESC"
	UserOrderByInputUpdatedAtAsc  UserOrderByInput = "updatedAt_ASC"
	UserOrderByInputUpdatedAtDesc UserOrderByInput = "updatedAt_DESC"
)

type MutationType string

const (
	MutationTypeCreated MutationType = "CREATED"
	MutationTypeUpdated MutationType = "UPDATED"
	MutationTypeDeleted MutationType = "DELETED"
)

type UserWhereUniqueInput struct {
	ID    *string `json:"id,omitempty"`
	Email *string `json:"email,omitempty"`
}

type UserWhereInput struct {
	ID                     *string          `json:"id,omitempty"`
	IDNot                  *string          `json:"id_not,omitempty"`
	IDIn                   []string         `json:"id_in,omitempty"`
	IDNotIn                []string         `json:"id_not_in,omitempty"`
	IDLt                   *string          `json:"id_lt,omitempty"`
	IDLte                  *string          `json:"id_lte,omitempty"`
	IDGt                   *string          `json:"id_gt,omitempty"`
	IDGte                  *string          `json:"id_gte,omitempty"`
	IDContains             *string          `json:"id_contains,omitempty"`
	IDNotContains          *string          `json:"id_not_contains,omitempty"`
	IDStartsWith           *string          `json:"id_starts_with,omitempty"`
	IDNotStartsWith        *string          `json:"id_not_starts_with,omitempty"`
	IDEndsWith             *string          `json:"id_ends_with,omitempty"`
	IDNotEndsWith          *string          `json:"id_not_ends_with,omitempty"`
	Email                  *string          `json:"email,omitempty"`
	EmailNot               *string          `json:"email_not,omitempty"`
	EmailIn                []string         `json:"email_in,omitempty"`
	EmailNotIn             []string         `json:"email_not_in,omitempty"`
	EmailLt                *string          `json:"email_lt,omitempty"`
	EmailLte               *string          `json:"email_lte,omitempty"`
	EmailGt                *string          `json:"email_gt,omitempty"`
	EmailGte               *string          `json:"email_gte,omitempty"`
	EmailContains          *string          `json:"email_contains,omitempty"`
	EmailNotContains       *string          `json:"email_not_contains,omitempty"`
	EmailStartsWith        *string          `json:"email_starts_with,omitempty"`
	EmailNotStartsWith     *string          `json:"email_not_starts_with,omitempty"`
	EmailEndsWith          *string          `json:"email_ends_with,omitempty"`
	EmailNotEndsWith       *string          `json:"email_not_ends_with,omitempty"`
	FirstName              *string          `json:"firstName,omitempty"`
	FirstNameNot           *string          `json:"firstName_not,omitempty"`
	FirstNameIn            []string         `json:"firstName_in,omitempty"`
	FirstNameNotIn         []string         `json:"firstName_not_in,omitempty"`
	FirstNameLt            *string          `json:"firstName_lt,omitempty"`
	FirstNameLte           *string          `json:"firstName_lte,omitempty"`
	FirstNameGt            *string          `json:"firstName_gt,omitempty"`
	FirstNameGte           *string          `json:"firstName_gte,omitempty"`
	FirstNameContains      *string          `json:"firstName_contains,omitempty"`
	FirstNameNotContains   *string          `json:"firstName_not_contains,omitempty"`
	FirstNameStartsWith    *string          `json:"firstName_starts_with,omitempty"`
	FirstNameNotStartsWith *string          `json:"firstName_not_starts_with,omitempty"`
	FirstNameEndsWith      *string          `json:"firstName_ends_with,omitempty"`
	FirstNameNotEndsWith   *string          `json:"firstName_not_ends_with,omitempty"`
	LastName               *string          `json:"lastName,omitempty"`
	LastNameNot            *string          `json:"lastName_not,omitempty"`
	LastNameIn             []string         `json:"lastName_in,omitempty"`
	LastNameNotIn          []string         `json:"lastName_not_in,omitempty"`
	LastNameLt             *string          `json:"lastName_lt,omitempty"`
	LastNameLte            *string          `json:"lastName_lte,omitempty"`
	LastNameGt             *string          `json:"lastName_gt,omitempty"`
	LastNameGte            *string          `json:"lastName_gte,omitempty"`
	LastNameContains       *string          `json:"lastName_contains,omitempty"`
	LastNameNotContains    *string          `json:"lastName_not_contains,omitempty"`
	LastNameStartsWith     *string          `json:"lastName_starts_with,omitempty"`
	LastNameNotStartsWith  *string          `json:"lastName_not_starts_with,omitempty"`
	LastNameEndsWith       *string          `json:"lastName_ends_with,omitempty"`
	LastNameNotEndsWith    *string          `json:"lastName_not_ends_with,omitempty"`
	UserName               *string          `json:"userName,omitempty"`
	UserNameNot            *string          `json:"userName_not,omitempty"`
	UserNameIn             []string         `json:"userName_in,omitempty"`
	UserNameNotIn          []string         `json:"userName_not_in,omitempty"`
	UserNameLt             *string          `json:"userName_lt,omitempty"`
	UserNameLte            *string          `json:"userName_lte,omitempty"`
	UserNameGt             *string          `json:"userName_gt,omitempty"`
	UserNameGte            *string          `json:"userName_gte,omitempty"`
	UserNameContains       *string          `json:"userName_contains,omitempty"`
	UserNameNotContains    *string          `json:"userName_not_contains,omitempty"`
	UserNameStartsWith     *string          `json:"userName_starts_with,omitempty"`
	UserNameNotStartsWith  *string          `json:"userName_not_starts_with,omitempty"`
	UserNameEndsWith       *string          `json:"userName_ends_with,omitempty"`
	UserNameNotEndsWith    *string          `json:"userName_not_ends_with,omitempty"`
	Gender                 *string          `json:"gender,omitempty"`
	GenderNot              *string          `json:"gender_not,omitempty"`
	GenderIn               []string         `json:"gender_in,omitempty"`
	GenderNotIn            []string         `json:"gender_not_in,omitempty"`
	GenderLt               *string          `json:"gender_lt,omitempty"`
	GenderLte              *string          `json:"gender_lte,omitempty"`
	GenderGt               *string          `json:"gender_gt,omitempty"`
	GenderGte              *string          `json:"gender_gte,omitempty"`
	GenderContains         *string          `json:"gender_contains,omitempty"`
	GenderNotContains      *string          `json:"gender_not_contains,omitempty"`
	GenderStartsWith       *string          `json:"gender_starts_with,omitempty"`
	GenderNotStartsWith    *string          `json:"gender_not_starts_with,omitempty"`
	GenderEndsWith         *string          `json:"gender_ends_with,omitempty"`
	GenderNotEndsWith      *string          `json:"gender_not_ends_with,omitempty"`
	IsAdmin                *bool            `json:"isAdmin,omitempty"`
	IsAdminNot             *bool            `json:"isAdmin_not,omitempty"`
	CreatedAt              *string          `json:"createdAt,omitempty"`
	CreatedAtNot           *string          `json:"createdAt_not,omitempty"`
	CreatedAtIn            []string         `json:"createdAt_in,omitempty"`
	CreatedAtNotIn         []string         `json:"createdAt_not_in,omitempty"`
	CreatedAtLt            *string          `json:"createdAt_lt,omitempty"`
	CreatedAtLte           *string          `json:"createdAt_lte,omitempty"`
	CreatedAtGt            *string          `json:"createdAt_gt,omitempty"`
	CreatedAtGte           *string          `json:"createdAt_gte,omitempty"`
	UpdatedAt              *string          `json:"updatedAt,omitempty"`
	UpdatedAtNot           *string          `json:"updatedAt_not,omitempty"`
	UpdatedAtIn            []string         `json:"updatedAt_in,omitempty"`
	UpdatedAtNotIn         []string         `json:"updatedAt_not_in,omitempty"`
	UpdatedAtLt            *string          `json:"updatedAt_lt,omitempty"`
	UpdatedAtLte           *string          `json:"updatedAt_lte,omitempty"`
	UpdatedAtGt            *string          `json:"updatedAt_gt,omitempty"`
	UpdatedAtGte           *string          `json:"updatedAt_gte,omitempty"`
	And                    []UserWhereInput `json:"AND,omitempty"`
	Or                     []UserWhereInput `json:"OR,omitempty"`
	Not                    []UserWhereInput `json:"NOT,omitempty"`
}

type UserCreateInput struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
	Gender    string `json:"gender"`
	IsAdmin   *bool  `json:"isAdmin,omitempty"`
}

type UserUpdateInput struct {
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	UserName  *string `json:"userName,omitempty"`
	Gender    *string `json:"gender,omitempty"`
	IsAdmin   *bool   `json:"isAdmin,omitempty"`
}

type UserUpdateManyMutationInput struct {
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	UserName  *string `json:"userName,omitempty"`
	Gender    *string `json:"gender,omitempty"`
	IsAdmin   *bool   `json:"isAdmin,omitempty"`
}

type UserSubscriptionWhereInput struct {
	MutationIn                 []MutationType               `json:"mutation_in,omitempty"`
	UpdatedFieldsContains      *string                      `json:"updatedFields_contains,omitempty"`
	UpdatedFieldsContainsEvery []string                     `json:"updatedFields_contains_every,omitempty"`
	UpdatedFieldsContainsSome  []string                     `json:"updatedFields_contains_some,omitempty"`
	Node                       *UserWhereInput              `json:"node,omitempty"`
	And                        []UserSubscriptionWhereInput `json:"AND,omitempty"`
	Or                         []UserSubscriptionWhereInput `json:"OR,omitempty"`
	Not                        []UserSubscriptionWhereInput `json:"NOT,omitempty"`
}

type UserExec struct {
	exec *prisma.Exec
}

func (instance UserExec) Exec(ctx context.Context) (*User, error) {
	var v User
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserExecArray struct {
	exec *prisma.Exec
}

func (instance UserExecArray) Exec(ctx context.Context) ([]User, error) {
	var v []User
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
	Gender    string `json:"gender"`
	IsAdmin   bool   `json:"isAdmin"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserConnectionExec struct {
	exec *prisma.Exec
}

func (instance *UserConnectionExec) PageInfo() *PageInfoExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "PageInfo"},
		"pageInfo",
		[]string{"hasNextPage", "hasPreviousPage", "startCursor", "endCursor"})

	return &PageInfoExec{ret}
}

func (instance *UserConnectionExec) Edges() *UserEdgeExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "UserEdge"},
		"edges",
		[]string{"cursor"})

	return &UserEdgeExec{ret}
}

func (instance *UserConnectionExec) Aggregate(ctx context.Context) (Aggregate, error) {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "AggregateUser"},
		"aggregate",
		[]string{"count"})

	var v Aggregate
	_, err := ret.Exec(ctx, &v)
	return v, err
}

func (instance UserConnectionExec) Exec(ctx context.Context) (*UserConnection, error) {
	var v UserConnection
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserConnectionExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserConnectionExecArray struct {
	exec *prisma.Exec
}

func (instance UserConnectionExecArray) Exec(ctx context.Context) ([]UserConnection, error) {
	var v []UserConnection
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserConnection struct {
}

type PageInfoExec struct {
	exec *prisma.Exec
}

func (instance PageInfoExec) Exec(ctx context.Context) (*PageInfo, error) {
	var v PageInfo
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance PageInfoExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type PageInfoExecArray struct {
	exec *prisma.Exec
}

func (instance PageInfoExecArray) Exec(ctx context.Context) ([]PageInfo, error) {
	var v []PageInfo
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor,omitempty"`
	EndCursor       *string `json:"endCursor,omitempty"`
}

type UserEdgeExec struct {
	exec *prisma.Exec
}

func (instance *UserEdgeExec) Node() *UserExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "User"},
		"node",
		[]string{"id", "email", "firstName", "lastName", "userName", "gender", "isAdmin", "createdAt", "updatedAt"})

	return &UserExec{ret}
}

func (instance UserEdgeExec) Exec(ctx context.Context) (*UserEdge, error) {
	var v UserEdge
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserEdgeExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserEdgeExecArray struct {
	exec *prisma.Exec
}

func (instance UserEdgeExecArray) Exec(ctx context.Context) ([]UserEdge, error) {
	var v []UserEdge
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserEdge struct {
	Cursor string `json:"cursor"`
}

type UserSubscriptionPayloadExec struct {
	exec *prisma.Exec
}

func (instance *UserSubscriptionPayloadExec) Node() *UserExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "User"},
		"node",
		[]string{"id", "email", "firstName", "lastName", "userName", "gender", "isAdmin", "createdAt", "updatedAt"})

	return &UserExec{ret}
}

func (instance *UserSubscriptionPayloadExec) PreviousValues() *UserPreviousValuesExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "UserPreviousValues"},
		"previousValues",
		[]string{"id", "email", "firstName", "lastName", "userName", "gender", "isAdmin", "createdAt", "updatedAt"})

	return &UserPreviousValuesExec{ret}
}

func (instance UserSubscriptionPayloadExec) Exec(ctx context.Context) (*UserSubscriptionPayload, error) {
	var v UserSubscriptionPayload
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserSubscriptionPayloadExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserSubscriptionPayloadExecArray struct {
	exec *prisma.Exec
}

func (instance UserSubscriptionPayloadExecArray) Exec(ctx context.Context) ([]UserSubscriptionPayload, error) {
	var v []UserSubscriptionPayload
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserSubscriptionPayload struct {
	UpdatedFields []string `json:"updatedFields,omitempty"`
}

type UserPreviousValuesExec struct {
	exec *prisma.Exec
}

func (instance UserPreviousValuesExec) Exec(ctx context.Context) (*UserPreviousValues, error) {
	var v UserPreviousValues
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserPreviousValuesExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserPreviousValuesExecArray struct {
	exec *prisma.Exec
}

func (instance UserPreviousValuesExecArray) Exec(ctx context.Context) ([]UserPreviousValues, error) {
	var v []UserPreviousValues
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserPreviousValues struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
	Gender    string `json:"gender"`
	IsAdmin   bool   `json:"isAdmin"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
