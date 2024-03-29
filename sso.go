package plugin

import (
	"time"

	"github.com/vroomy/httpserve"
)

// SSOLogin is the handler for logging in with SSO
func (p *plugin) SSOLogin(ctx *httpserve.Context) {
	var err error
	if ctx.Get("userID") != "" {
		// User is already logged in, write no content success and return
		ctx.WriteNoContent()
		return
	}

	// Get login code from URL param
	loginCode := ctx.Param("loginCode")

	// Attempt to login with code
	if err = p.jump.SSOLogin(ctx, loginCode); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteNoContent()
}

// SSOMultiLogin is the handler for logging in with SSO
func (p *plugin) SSOMultiLogin(ctx *httpserve.Context) {
	var err error
	if ctx.Get("userID") != "" {
		// User is already logged in, write no content success and return
		ctx.WriteNoContent()
		return
	}

	// Get login code from URL param
	loginCode := ctx.Param("loginCode")

	// Attempt to login with code
	if err = p.jump.SSOMultiLogin(ctx, loginCode, time.Minute*5); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteNoContent()
}
