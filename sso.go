package main

import (
	"github.com/vroomy/common"
)

// SSOLogin is the handler for logging in with SSO
func SSOLogin(ctx common.Context) {
	var err error
	loginCode := ctx.Param("loginCode")

	if err = p.jump.SSOLogin(ctx, loginCode); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteNoContent()
}
