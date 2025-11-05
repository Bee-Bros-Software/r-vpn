package roles

import (
	"github.com/Bee-Bros-Software/r-vpn/management/server/permissions/modules"
	"github.com/Bee-Bros-Software/r-vpn/management/server/permissions/operations"
	"github.com/Bee-Bros-Software/r-vpn/management/server/types"
)

type RolePermissions struct {
	Role         types.UserRole
	Permissions  Permissions
	AutoAllowNew map[operations.Operation]bool
}

type Permissions map[modules.Module]map[operations.Operation]bool

var RolesMap = map[types.UserRole]RolePermissions{
	types.UserRoleOwner:        Owner,
	types.UserRoleAdmin:        Admin,
	types.UserRoleUser:         User,
	types.UserRoleAuditor:      Auditor,
	types.UserRoleNetworkAdmin: NetworkAdmin,
}
