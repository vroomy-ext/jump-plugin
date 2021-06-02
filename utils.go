package plugin

import (
	"fmt"
	"strings"

	"github.com/gdbu/jump/permissions"
)

func newResourceKey(name, userID string) (resourceKey string) {
	if len(userID) == 0 {
		return name
	}

	return fmt.Sprintf("%s::%s", name, userID)
}

func getPermissions(permsStr string) (a permissions.Action) {
	hasRead := strings.Contains(permsStr, "r")
	hasWrite := strings.Contains(permsStr, "w")
	hasDelete := strings.Contains(permsStr, "d")

	if hasRead {
		a |= permissions.ActionRead
	}

	if hasWrite {
		a |= permissions.ActionWrite
	}

	if hasDelete {
		a |= permissions.ActionDelete
	}

	return
}
