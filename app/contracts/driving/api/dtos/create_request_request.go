package dtos

type CreateRequestReq struct {
	CableAmount     uint   `json:"cableAmount"`
	ContractCounter uint   `json:"contractCounter"`
	ContractorEmail string `json:"contractorEmail"`
}
