package main

import (
	"github.com/Hatch1fy/httpserve"
	"github.com/Hatch1fy/jump/users"
)

const (
	errPasswordsDontMatch = "New password does not match confirmation"
)

// GetUserID will get the ID of the currently logged in user
func GetUserID(ctx *httpserve.Context) (res httpserve.Response) {
	return httpserve.NewJSONResponse(200, ctx.Get("userID"))
}

// GetUser will get a user by ID
func GetUser(ctx *httpserve.Context) (res httpserve.Response) {
	var (
		user *users.User
		err  error
	)

	if user, err = p.jump.GetUser(ctx.Param("userID")); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	return httpserve.NewJSONResponse(200, user)
}

// UpdateEmail will update a user's email address
func UpdateEmail(ctx *httpserve.Context) (res httpserve.Response) {
	var (
		user users.User
		err  error
	)

	if err = ctx.BindJSON(&user); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	userID := ctx.Param("userID")

	if err = p.jump.UpdateEmail(userID, user.Email); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	return httpserve.NewNoContentResponse()
}

// UpdatePassword is the update password handler
func UpdatePassword(ctx *httpserve.Context) (res httpserve.Response) {
	var (
		user users.User
		err  error
	)

	if err = ctx.BindJSON(&user); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	userID := ctx.Param("userID")

	if err = p.jump.UpdatePassword(userID, user.Password); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	return httpserve.NewNoContentResponse()
}

// ChangePassword accepts current, new, and confirm password fields
func ChangePassword(ctx *httpserve.Context) (res httpserve.Response) {
	var (
		cpr changePasswordRequest
		err error
	)

	if err = ctx.BindJSON(&cpr); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	userID := ctx.Param("userID")

	if _, err = p.jump.Users().Match(userID, cpr.Current); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	if err = p.jump.UpdatePassword(userID, cpr.New); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	return httpserve.NewNoContentResponse()
}

type changePasswordRequest struct {
	Current string `json:"current"`
	New     string `json:"new"`
}
