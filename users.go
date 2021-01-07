package main

import (
	"fmt"

	"github.com/gdbu/jump/users"
	"github.com/vroomy/common"
)

// CreateUserRequest is the request used to create a user
type CreateUserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password,omitempty" form:"password"`
}

// CreateUserResponse is returned after a user is created
type CreateUserResponse struct {
	UserID string `json:"userID"`
	APIKey string `json:"apiKey"`
}

// CreateUser is a handler for creating a new user
func CreateUser(ctx common.Context) {
	var (
		req CreateUserRequest
		err error
	)

	if err = ctx.Bind(&req); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	fmt.Println("REQ?", req, err)

	var resp CreateUserResponse
	if resp.UserID, resp.APIKey, err = p.jump.CreateUser(req.Email, req.Password, "users"); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	// Set createdUserID field
	ctx.Put("createdUserID", resp.UserID)

	// Grab url values from request
	q := ctx.Request().URL.Query()

	// Check to see if redirect query value has been set
	if redirect := q.Get("redirect"); len(redirect) > 0 {
		ctx.Redirect(302, redirect)
		return
	}

	ctx.WriteJSON(200, resp)
}

// GetUsersList will get the current users list
func GetUsersList(ctx common.Context) {
	var (
		us  []*users.User
		err error
	)

	if us, err = p.jump.GetUsersList(); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteJSON(200, us)
}
