package dtos

type CreateUserReq struct {
	Role  string `json:"role"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
