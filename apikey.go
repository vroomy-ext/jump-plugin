package plugin

import (
	"github.com/gdbu/jump/apikeys"
	"github.com/vroomy/common"
)

// GetAPIKeysByUser is the handler for retrieving the api keys for a user
func (p *plugin) GetAPIKeysByUser(ctx common.Context) {
	var (
		as  []*apikeys.APIKey
		err error
	)

	if as, err = p.jump.GetAPIKeysByUser(ctx.Param("userID")); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteJSON(200, as)
}
