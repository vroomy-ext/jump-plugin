package main

import (
	"fmt"

	"github.com/gdbu/jump/users"
	"github.com/hatchify/errors"
	"github.com/mojura/mojura"

	"github.com/vroomy/httpserve"
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

	contentType := ctx.Request.Header.Get("Content-Type")
	switch contentType {
	case "application/x-www-form-urlencoded":
		if err := ctx.Request.ParseForm(); err != nil {
			err = fmt.Errorf("error parsing form: %v", err)
			return httpserve.NewHTMLResponse(400, []byte(err.Error()))
		}

		login.Email = ctx.Request.Form.Get("email")
		login.Password = ctx.Request.Form.Get("password")

	default:
		if err = ctx.BindJSON(&login); err != nil {
			return httpserve.NewJSONResponse(400, err)
		}
	}

	if login.ID, err = p.jump.Login(ctx, login.Email, login.Password); err != nil {
		if err == mojura.ErrEntryNotFound {
			err = ErrNoLoginFound
		}

		// TODO: Respond differently based on content type
		return httpserve.NewJSONResponse(400, err)
	}

	var user *users.User
	if user, err = p.jump.GetUser(login.ID); err != nil {
		err = fmt.Errorf("error getting user %s: %v", login.ID, err)
		// TODO: Respond differently based on content type
		return httpserve.NewJSONResponse(400, err)
	}

	// Grab url values from request
	q := ctx.Request.URL.Query()

	// Check to see if redirect query value has been set
	if redirect := q.Get("redirect"); len(redirect) > 0 {
		// Client is expecting a redirect on success, return 302 to provided value
		return httpserve.NewRedirectResponse(302, redirect)
	}

	// TODO: Respond differently based on content type
	return httpserve.NewJSONResponse(200, user)
}

// Logout is the logout handler
func Logout(ctx *httpserve.Context) (res httpserve.Response) {
	var err error
	if err = p.jump.Logout(ctx); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	return ctx.NewNoContentResponse()
}
