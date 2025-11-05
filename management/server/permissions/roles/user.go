package roles

import (
	"github.com/Bee-Bros-Software/r-vpn/management/server/permissions/operations"
	"github.com/Bee-Bros-Software/r-vpn/management/server/types"
)

var User = RolePermissions{
	Role: types.UserRoleUser,
	AutoAllowNew: map[operations.Operation]bool{
		operations.Read:   false,
		operations.Create: false,
		operations.Update: false,
		operations.Delete: false,
	},
}
