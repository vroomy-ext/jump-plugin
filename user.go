package main

import (
	"github.com/gdbu/jump/users"
	vroomy "github.com/vroomy/common"
)

const (
	errPasswordsDontMatch = "New password does not match confirmation"
)

// GetUserID will get the ID of the currently logged in user
func GetUserID(ctx vroomy.Context) (res vroomy.Response) {
	return ctx.NewJSONResponse(200, ctx.Get("userID"))
}

// GetUser will get a user by ID
func GetUser(ctx vroomy.Context) (res vroomy.Response) {
	var (
		user *users.User
		err  error
	)

	if user, err = p.jump.GetUser(ctx.Param("userID")); err != nil {
		return ctx.NewJSONResponse(400, err)
	}

	return ctx.NewJSONResponse(200, user)
}

// UpdateEmail will update a user's email address
func UpdateEmail(ctx vroomy.Context) (res vroomy.Response) {
	var (
		user users.User
		err  error
	)

	if err = ctx.BindJSON(&user); err != nil {
		return ctx.NewJSONResponse(400, err)
	}

	userID := ctx.Param("userID")

	if err = p.jump.UpdateEmail(userID, user.Email); err != nil {
		return ctx.NewJSONResponse(400, err)
	}

	return ctx.NewNoContentResponse()
}

// UpdatePassword is the update password handler
func UpdatePassword(ctx vroomy.Context) (res vroomy.Response) {
	var (
		user users.User
		err  error
	)

	if err = ctx.BindJSON(&user); err != nil {
		return ctx.NewJSONResponse(400, err)
	}

	userID := ctx.Param("userID")

	if err = p.jump.UpdatePassword(userID, user.Password); err != nil {
		return ctx.NewJSONResponse(400, err)
	}

	return ctx.NewNoContentResponse()
}

// ChangePassword accepts current, new, and confirm password fields
func ChangePassword(ctx vroomy.Context) (res vroomy.Response) {
	var (
		cpr changePasswordRequest
		err error
	)

	if err = ctx.BindJSON(&cpr); err != nil {
		return ctx.NewJSONResponse(400, err)
	}

	userID := ctx.Param("userID")

	if _, err = p.jump.Users().Match(userID, cpr.Current); err != nil {
		return ctx.NewJSONResponse(400, err)
	}

	if err = p.jump.UpdatePassword(userID, cpr.New); err != nil {
		return ctx.NewJSONResponse(400, err)
	}

	return ctx.NewNoContentResponse()
}

type changePasswordRequest struct {
	Current string `json:"current"`
	New     string `json:"new"`
}

// EnableUser is the handler for enabling a user
func EnableUser(ctx vroomy.Context) (res vroomy.Response) {
	var (
		userID string
		err    error
	)

	if err = p.jump.EnableUser(userID); err != nil {
		return ctx.NewJSONResponse(400, err)
	}

	return ctx.NewNoContentResponse()
}

// DisableUser is the handler for disabling a user
// Note: This will kill all active sessions for this user
func DisableUser(ctx vroomy.Context) (res vroomy.Response) {
	var (
		userID string
		err    error
	)

	if err = p.jump.DisableUser(userID); err != nil {
		return ctx.NewJSONResponse(400, err)
	}

	return ctx.NewNoContentResponse()
}
