package roles

import (
	"github.com/Bee-Bros-Software/r-vpn/management/server/permissions/operations"
	"github.com/Bee-Bros-Software/r-vpn/management/server/types"
)

var Owner = RolePermissions{
	Role: types.UserRoleOwner,
	AutoAllowNew: map[operations.Operation]bool{
		operations.Read:   true,
		operations.Create: true,
		operations.Update: true,
		operations.Delete: true,
	},
}
