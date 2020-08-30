package main

import (
	"github.com/gdbu/jump/users"
	"github.com/vroomy/httpserve"
)

// CreateUserRequest is the request used to create a user
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUserResponse is returned after a user is created
type CreateUserResponse struct {
	UserID string `json:"userID"`
	APIKey string `json:"apiKey"`
}

// CreateUser is a handler for creating a new user
func CreateUser(ctx *httpserve.Context) (res httpserve.Response) {
	var (
		req CreateUserRequest
		err error
	)

	if err = ctx.BindJSON(&req); err != nil {
		ctx.NewJSONResponse(400, err)
	}

	var resp CreateUserResponse
	if resp.UserID, resp.APIKey, err = p.jump.CreateUser(req.Email, req.Password, "users"); err != nil {
		return ctx.NewJSONResponse(400, err)
	}

	return ctx.NewJSONResponse(200, resp)
}

// GetUsersList will get the current users list
func GetUsersList(ctx *httpserve.Context) (res httpserve.Response) {
	var (
		us  []*users.User
		err error
	)

	if us, err = p.jump.GetUsersList(); err != nil {
		return ctx.NewJSONResponse(400, err)
	}

	return ctx.NewJSONResponse(200, us)
}
