package main

import (
	"fmt"

	"github.com/Hatch1fy/errors"
	"github.com/Hatch1fy/httpserve"
	"github.com/Hatch1fy/jump/users"
	core "github.com/gdbu/dbl"
)

const (
	// ErrNoLoginFound is returned when an email address is provided that is not found within the system
	ErrNoLoginFound = errors.Error("no login was found for the provided email address")
)

// Login is the login handler
func Login(ctx *httpserve.Context) (res httpserve.Response) {
	var (
		login users.User
		err   error
	)

	if err = ctx.BindJSON(&login); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	if login.ID, err = p.jump.Login(ctx, login.Email, login.Password); err != nil {
		if err == core.ErrEntryNotFound {
			err = ErrNoLoginFound
		}

		return httpserve.NewJSONResponse(400, err)
	}

	var user *users.User
	if user, err = p.jump.GetUser(login.ID); err != nil {
		err = fmt.Errorf("error getting user %s: %v", login.ID, err)
		return httpserve.NewJSONResponse(400, err)
	}

	return httpserve.NewJSONResponse(200, user)
}

// Logout is the logout handler
func Logout(ctx *httpserve.Context) (res httpserve.Response) {
	var err error
	if err = p.jump.Logout(ctx); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	return httpserve.NewNoContentResponse()
}
