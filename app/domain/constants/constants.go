package constants

const (
	RoleAdmin      = "admin"
	RolePlanner    = "planner"
	RoleSupplier   = "supplier"
	RoleContractor = "contractor"
)

var (
	RoleList = []string{RoleAdmin, RolePlanner, RoleSupplier, RoleContractor}
)

const (
	StatusNew       = "new"
	StatusReady     = "ready"
	StatusCollected = "collected"
	StatusCanceled  = "canceled"
)

const (
	ActionCreate = "create"
	ActionUpdate = "update"
	ActionCancel = "cancel"
)
