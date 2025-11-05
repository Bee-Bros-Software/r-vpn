package users

import (
	"github.com/Bee-Bros-Software/r-vpn/management/server/permissions/roles"
	"github.com/Bee-Bros-Software/r-vpn/management/server/types"
)

// Wrapped UserInfo with Role Permissions
type UserInfoWithPermissions struct {
	*types.UserInfo

	Permissions roles.Permissions
	Restricted  bool
}
