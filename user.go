package main

import (
	"github.com/gdbu/jump/users"
	"github.com/vroomy/common"
)

const (
	errPasswordsDontMatch = "New password does not match confirmation"
)

// GetUserID will get the ID of the currently logged in user
func GetUserID(ctx common.Context) {
	ctx.WriteJSON(200, ctx.Get("userID"))
	return
}

// GetUser will get a user by ID
func GetUser(ctx common.Context) {
	var (
		user *users.User
		err  error
	)

	if user, err = p.jump.GetUser(ctx.Param("userID")); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteJSON(200, user)
	return
}

// UpdateEmail will update a user's email address
func UpdateEmail(ctx common.Context) {
	var (
		user users.User
		err  error
	)

	if err = ctx.Bind(&user); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	userID := ctx.Param("userID")

	if err = p.jump.UpdateEmail(userID, user.Email); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteNoContent()
}

// UpdatePassword is the update password handler
func UpdatePassword(ctx common.Context) {
	var (
		user users.User
		err  error
	)

	if err = ctx.Bind(&user); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	userID := ctx.Param("userID")

	if err = p.jump.UpdatePassword(userID, user.Password); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteNoContent()
}

// ChangePassword accepts current, new, and confirm password fields
func ChangePassword(ctx common.Context) {
	var (
		cpr changePasswordRequest
		err error
	)

	if err = ctx.Bind(&cpr); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	userID := ctx.Param("userID")

	if _, err = p.jump.Users().Match(userID, cpr.Current); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	if err = p.jump.UpdatePassword(userID, cpr.New); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteNoContent()
}

type changePasswordRequest struct {
	Current string `json:"current"`
	New     string `json:"new"`
}

// EnableUser is the handler for enabling a user
func EnableUser(ctx common.Context) {
	var (
		userID string
		err    error
	)

	if err = p.jump.EnableUser(userID); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteNoContent()
}

// DisableUser is the handler for disabling a user
// Note: This will kill all active sessions for this user
func DisableUser(ctx common.Context) {
	var (
		userID string
		err    error
	)

	if err = p.jump.DisableUser(userID); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteNoContent()
}

// VerifyUser is the handler for verifying a user
func VerifyUser(ctx common.Context) {
	var (
		userID string
		err    error
	)

	if err = p.jump.VerifyUser(userID); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteNoContent()
}
