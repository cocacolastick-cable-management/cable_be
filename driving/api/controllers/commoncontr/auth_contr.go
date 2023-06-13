package commoncontr

import (
	"github.com/cable_management/cable_be/app/usecases/commomcase"
	"github.com/cable_management/cable_be/driving/api/_share/constants"
	"github.com/cable_management/cable_be/driving/api/_share/responses"
	"github.com/cable_management/cable_be/driving/api/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IAuthController interface {
	SignIn(ctx *gin.Context)
}

type AuthController struct {
	signInCase commomcase.ISignIn
}

func NewAuthController(signInCase commomcase.ISignIn) *AuthController {
	return &AuthController{signInCase: signInCase}
}

func (a AuthController) SignIn(ctx *gin.Context) {

	var reqBody = ctx.MustGet(middlewares.BodyKey).(*commomcase.SignInRequest)

	res, err := a.signInCase.Handle(reqBody)
	if err != nil {
		ctx.Set(constants.ErrKey, err)
		ctx.Next()
		return
	}

	ctx.JSON(http.StatusOK, responses.ResponseType{
		Code:    "OK",
		Message: "sign in successfully",
		Payload: res,
	})
	return
}
