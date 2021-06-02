package plugin

import (
	"github.com/gdbu/emailvalidator"
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
func (p *plugin) CreateUser(ctx common.Context) {
	var (
		resp CreateUserResponse
		err  error
	)

	if resp, err = createUser(ctx); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteJSON(200, resp)
}

// CreateUserMW is a middleware for creating a new user
func (p *plugin) CreateUserMW(ctx common.Context) {
	var err error
	if _, err = createUser(ctx); err != nil {
		ctx.WriteJSON(400, err)
		return
	}
}

// SignUp is a handler for a self sign-up for a new user
func (p *plugin) SignUp(ctx common.Context) {
	var (
		resp CreateUserResponse
		err  error
	)

	if resp, err = createUser(ctx); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

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
func (p *plugin) GetUsersList(ctx common.Context) {
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

func createUser(ctx common.Context) (resp CreateUserResponse, err error) {
	var req CreateUserRequest
	if err = ctx.Bind(&req); err != nil {
		return
	}

	if err = emailvalidator.Validate(req.Email); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	if resp.UserID, resp.APIKey, err = p.jump.CreateUser(req.Email, req.Password, "users"); err != nil {
		return
	}

	// Set createdUserID field
	ctx.Put("createdUserID", resp.UserID)
	return
}
