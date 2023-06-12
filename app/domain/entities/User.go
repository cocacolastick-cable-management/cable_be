package entities

const (
	RoleAdmin      = "admin"
	RolePlanner    = "planner"
	RoleSupplier   = "supplier"
	RoleContractor = "contractor"
)

type User struct {
	EntityBase

	Role         string
	Email        string
	Name         string
	PasswordHash string
}
