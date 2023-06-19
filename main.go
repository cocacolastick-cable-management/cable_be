package main

import (
	"github.com/cable_management/cable_be/_share/env"
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driven/email"
	"github.com/cable_management/cable_be/app/contracts/driving/api/controllers/admincontr"
	"github.com/cable_management/cable_be/app/contracts/driving/api/controllers/commoncontr"
	"github.com/cable_management/cable_be/app/contracts/driving/api/controllers/plannercontr"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/validations"
	"github.com/cable_management/cable_be/app/domain/services"
	"github.com/cable_management/cable_be/app/usecases/admincase"
	"github.com/cable_management/cable_be/app/usecases/commomcase"
	"github.com/cable_management/cable_be/app/usecases/plannercase"
	"github.com/cable_management/cable_be/driven/database"
	imrepos "github.com/cable_management/cable_be/driven/database/repos"
	imemail "github.com/cable_management/cable_be/driven/email"
	imadmincontr "github.com/cable_management/cable_be/driving/api/controllers/admincontr"
	imcommoncontr "github.com/cable_management/cable_be/driving/api/controllers/commoncontr"
	implannercontr "github.com/cable_management/cable_be/driving/api/controllers/plannercontr"
	"github.com/cable_management/cable_be/driving/api/routers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log"
)

func main() {

	// NOTE - should not change the order of the calling methods
	BuildEnv()

	StartDb()

	StartEmail()

	BuildValidator()

	BuildDomain()

	StartApi()
}

// env
var environments env.Env

func BuildEnv() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//
	//ENV.DbDsn = os.Getenv("DB_DSN")

	environments.JwtSecret = "123467890qwertuiopasdfghjkl"

	environments.DbDsn = "host=localhost user=postgres password=postgrespw dbname=cable_db_2 port=32768 sslmode=disable TimeZone=Asia/Shanghai"

	environments.SmtpEmail = "vuphamlethanh@gmail.com"
	environments.SmtpHost = "smtp.gmail.com"
	environments.SmtpPort = "587"
	environments.SmtpPassword = "fzjugwhesxnbpixp"
}

// database
var (
	db                 *gorm.DB
	userRepo           repos.IUserRepo
	contractRepo       repos.IContractRepo
	requestRepo        repos.IRequestRepo
	requestHistoryRepo repos.IRequestHistoryRepo
)

func StartDb() {
	db = database.Init(environments.DbDsn)
	userRepo = imrepos.NewUserRepo(db)
	contractRepo = imrepos.NewContractRepo(db)
	requestRepo = imrepos.NewRequestRepo(db)
	requestHistoryRepo = imrepos.NewRequestHistoryRepo(db)
}

// email driven
var (
	emailDriven email.IEmail
)

func StartEmail() {
	emailDriven = imemail.NewEmail(imemail.Config{
		MailHost: environments.SmtpEmail,
		Host:     environments.SmtpHost,
		Port:     environments.SmtpPort,
		Password: environments.SmtpPassword,
	})
}

// validations
var (
	validation         *validator.Validate
	vlCreateUserReq    *validations.VlCreateUserReq
	vlCreateRequestReq *validations.VlCreateRequestReq
)

func BuildValidator() {

	validation = validator.New()

	vlCreateUserReq = validations.NewVlCreateUserReq(userRepo)
	vlCreateRequestReq = validations.NewVlCreateRequestReq(contractRepo, userRepo)

	validation.RegisterStructValidation(vlCreateUserReq.Handle, dtos.CreateUserReq{})
	validation.RegisterStructValidation(vlCreateRequestReq.Handle, dtos.CreateRequestReq{})
}

// domain
var (
	// services
	passwordService   services.IPasswordService
	userFac           services.IUserFactory
	tokenService      services.IAuthTokenService
	authorService     services.IAuthorizeService
	requestFac        services.IRequestFactory
	mailDataFac       services.IMailDataFactory
	requestHistoryFac services.IRequestHistoryFactory

	// usecases
	createUserCase          admincase.ICreateUser
	signInCase              commomcase.ISignIn
	updateUserIsActiveCase  admincase.IUpdateUserIsActive
	createRequestCase       plannercase.ICreateRequest
	updateRequestStatusCase commomcase.IUpdateRequestStatus
	getRequestListCase      plannercase.IGetRequestList
)

func BuildDomain() {

	// services
	passwordService = services.NewPasswordHash(validation)
	userFac = services.NewUserFactory(passwordService, userRepo, validation)
	tokenService = services.NewAuthTokenService(environments.JwtSecret)
	authorService = services.NewAuthorizeService(tokenService, userRepo)
	requestFac = services.NewRequestFactory(contractRepo, userRepo)
	mailDataFac = services.NewMailDataFactory(userRepo, requestHistoryRepo)
	requestHistoryFac = services.NewRequestHistoryFactory(requestRepo)

	// usecases
	createUserCase = admincase.NewCreateUser(userRepo, userFac, validation, authorService, emailDriven, passwordService)
	signInCase = commomcase.NewSignIn(userRepo, tokenService, passwordService)
	updateUserIsActiveCase = admincase.NewUpdateUserIsActive(userRepo, authorService, emailDriven)
	createRequestCase = plannercase.NewCreateRequest(authorService, validation, requestFac, requestRepo, requestHistoryFac, mailDataFac, requestHistoryRepo, emailDriven)
	updateRequestStatusCase = commomcase.NewUpdateRequestStatus(authorService, requestRepo, requestHistoryRepo, userRepo, validation, requestHistoryFac, mailDataFac, emailDriven)
	getRequestListCase = plannercase.NewGetRequestList(requestRepo, authorService)
}

// api
var (
	// controllers
	authContr           commoncontr.IAuthController
	adminUserContr      admincontr.IUserController
	plannerRequestContr plannercontr.IRequestContr
	commonRequestContr  commoncontr.IRequestController

	// routers
	baseRouters    routers.IRouterBase
	commonRouters  routers.IRouterBase
	adminRouters   routers.IRouterBase
	plannerRouters routers.IRouterBase
)

func StartApi() {

	// controllers
	authContr = imcommoncontr.NewAuthController(signInCase)
	adminUserContr = imadmincontr.NewUserController(createUserCase, updateUserIsActiveCase)
	plannerRequestContr = implannercontr.NewRequestContr(createRequestCase, getRequestListCase)
	commonRequestContr = imcommoncontr.NewRequestController(updateRequestStatusCase)

	// routers
	baseRouters = routers.NewRouterBase()
	commonRouters = routers.NewCommonRouters(authContr, commonRequestContr)
	adminRouters = routers.NewAdminRouters(adminUserContr)
	plannerRouters = routers.NewPlannerRouters(plannerRequestContr)

	// init
	engine := gin.Default()

	baseRouters.Register(engine)
	commonRouters.Register(engine)
	adminRouters.Register(engine)
	plannerRouters.Register(engine)

	//engine.NoRoute(middlewares.HandleGlobalErrors)

	err := engine.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
