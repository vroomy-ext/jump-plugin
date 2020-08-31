package main

import (
	"github.com/gdbu/apikeys"
	vroomy "github.com/vroomy/common"
)

// GetAPIKeysByUser is the handler for retrieving the api keys for a user
func GetAPIKeysByUser(ctx vroomy.Context) (res vroomy.Response) {
	var (
		as  []*apikeys.APIKey
		err error
	)

	if as, err = p.jump.GetAPIKeysByUser(ctx.Param("userID")); err != nil {
		return ctx.NewJSONResponse(400, err)
	}

	return ctx.NewJSONResponse(200, as)
}
