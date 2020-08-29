package main

import (
	"github.com/Hatch1fy/httpserve"
	"github.com/gdbu/apikeys"
)

// GetAPIKeysByUser is the handler for retrieving the api keys for a user
func GetAPIKeysByUser(ctx *httpserve.Context) (res httpserve.Response) {
	var (
		as  []*apikeys.APIKey
		err error
	)

	if as, err = p.jump.GetAPIKeysByUser(ctx.Param("userID")); err != nil {
		return ctx.NewJSONResponse(400, err)
	}

	return ctx.NewJSONResponse(200, as)
}
