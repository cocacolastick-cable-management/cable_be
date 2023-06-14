package plannercase

import (
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/domain/constants"
	"github.com/cable_management/cable_be/app/domain/services"
	"github.com/go-playground/validator/v10"
)

type ICreateRequest interface {
	Handle(accessToken string, req *dtos.CreateRequestReq) error
}

type CreateRequest struct {
	authorService services.IAuthorizeService
	validator     *validator.Validate
	requestFac    services.IRequestFactory
	requestRepo   repos.IRequestRepo
}

func NewCreateRequest(authorService services.IAuthorizeService, validator *validator.Validate, requestFac services.IRequestFactory, requestRepo repos.IRequestRepo) *CreateRequest {
	return &CreateRequest{authorService: authorService, validator: validator, requestFac: requestFac, requestRepo: requestRepo}
}

func (c CreateRequest) Handle(accessToken string, req *dtos.CreateRequestReq) error {

	// authorize
	_, err := c.authorService.Authorize(accessToken, []string{constants.RolePlanner}, nil)
	if err != nil {
		return err
	}

	// validate req
	err = c.validator.Struct(req)
	if err != nil {
		return err
	}

	// create req
	request, err := c.requestFac.CreateRequest(req.CableAmount, req.ContractCounter, req.ContractorEmail)
	if err != nil {
		return err
	}

	// insert database
	_ = c.requestRepo.Insert(request)

	// send email notif

	// sse

	return nil
}
