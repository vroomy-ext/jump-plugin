package plugin

import (
	"github.com/gdbu/jump/users"
	"github.com/vroomy/httpserve"
)

// GetUserID will get the ID of the currently logged in user
func (p *plugin) GetUserID(ctx *httpserve.Context) {
	ctx.WriteJSON(200, ctx.Get("userID"))
}

// GetUser will get a user by ID
func (p *plugin) GetUser(ctx *httpserve.Context) {
	var (
		user *users.User
		err  error
	)

	if user, err = p.jump.GetUser(ctx.Param("userID")); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteJSON(200, user)
}

// UpdateEmail will update a user's email address
func (p *plugin) UpdateEmail(ctx *httpserve.Context) {
	var (
		user changeRequest
		err  error
	)

	if err = ctx.Bind(&user); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	userID := ctx.Param("userID")

	if _, err = p.jump.UpdateEmail(userID, user.New); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteNoContent()
}

// UpdatePassword is the update password handler
func (p *plugin) UpdatePassword(ctx *httpserve.Context) {
	var (
		user users.User
		err  error
	)

	if err = ctx.Bind(&user); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	userID := ctx.Param("userID")

	if _, err = p.jump.UpdatePassword(userID, user.Password); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteNoContent()
}

// ChangePassword accepts current, new, and confirm password fields
func (p *plugin) ChangePassword(ctx *httpserve.Context) {
	var (
		cpr changeRequest
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

	if _, err = p.jump.UpdatePassword(userID, cpr.New); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteNoContent()
}

// EnableUser is the handler for enabling a user
func (p *plugin) EnableUser(ctx *httpserve.Context) {
	var err error
	userID := ctx.Param("userID")
	if err = p.jump.EnableUser(userID); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteNoContent()
}

// DisableUser is the handler for disabling a user
// Note: This will kill all active sessions for this user
func (p *plugin) DisableUser(ctx *httpserve.Context) {
	var err error
	userID := ctx.Param("userID")
	if err = p.jump.DisableUser(userID); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteNoContent()
}

// VerifyUser is the handler for verifying a user
func (p *plugin) VerifyUser(ctx *httpserve.Context) {
	var err error
	userID := ctx.Param("userID")
	if err = p.jump.VerifyUser(userID); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteNoContent()
}

type changeRequest struct {
	Current string `json:"current" form:"current"`
	New     string `json:"new" form:"new"`
}
