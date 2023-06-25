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

var (
	StatusList = []string{StatusNew, StatusReady, StatusCanceled, StatusCollected}
)

const (
	ActionCreate = "create"
	ActionUpdate = "update"
	ActionCancel = "cancel"

	ActionDisable = "disable"
	ActionEnable  = "enable"
)

var (
	ActionList = []string{ActionCancel, ActionUpdate, ActionCreate}
)

const (
	ObjTyRequest  = "request"
	ObjTyUser     = "user"
	ObjTyContract = "contract"
)
