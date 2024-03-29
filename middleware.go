package plugin

import (
	"github.com/vroomy/httpserve"
)

// SetUserMW is will check user permissions. Expects the following arguments:
//   - redirectOnFail (e.g. false)
//   - allowNonLoggedIn (e.g. false)
func (p *plugin) SetUserMW(args ...string) (h httpserve.Handler, err error) {
	var (
		redirectOnFail   bool
		allowNonLoggedIn bool
	)

	switch len(args) {
	case 2:
		redirectOnFail = args[0] == "true"
		allowNonLoggedIn = args[1] == "true"

	case 1:
		redirectOnFail = args[0] == "true"
	case 0:

	default:
		err = ErrInvalidSetUserArguments
		return
	}

	h = p.jump.NewSetUserIDMW(redirectOnFail, allowNonLoggedIn)
	return
}

// CheckPermissionsMW is will check user permissions. Expects the following arguments:
//   - groupName (e.g. users)
//   - paramKey (e.g. userID)
func (p *plugin) CheckPermissionsMW(args ...string) (h httpserve.Handler, err error) {
	var resourceName, paramKey string
	switch len(args) {
	case 1:
		resourceName = args[0]
	case 2:
		resourceName = args[0]
		paramKey = args[1]

	default:
		err = ErrInvalidCheckPermissionsArguments
		return
	}

	h = p.jump.NewCheckPermissionsMW(resourceName, paramKey)
	return
}

// GrantPermissionsMW is will grant user permissions. Expects the following arguments:
//   - groupName (e.g. users)
//   - paramKey (e.g. userID)
func (p *plugin) GrantPermissionsMW(args ...string) (h httpserve.Handler, err error) {
	if len(args) != 3 {
		err = ErrInvalidGrantPermissionsArguments
		return
	}

	resourceName := args[0]
	actions := getPermissions(args[1])
	adminActions := getPermissions(args[2])
	h = p.jump.NewGrantPermissionsMW(resourceName, actions, adminActions)
	return
}
