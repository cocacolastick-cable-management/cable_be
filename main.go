package main

import (
	"github.com/cable_management/cable_be/_share/env"
	"github.com/cable_management/cable_be/app/contracts/driven/database/repos"
	"github.com/cable_management/cable_be/app/contracts/driven/email"
	"github.com/cable_management/cable_be/app/contracts/driven/sse"
	"github.com/cable_management/cable_be/app/contracts/driving/api/controllers/admincontr"
	"github.com/cable_management/cable_be/app/contracts/driving/api/controllers/commoncontr"
	"github.com/cable_management/cable_be/app/contracts/driving/api/controllers/contractorcontr"
	"github.com/cable_management/cable_be/app/contracts/driving/api/controllers/plannercontr"
	"github.com/cable_management/cable_be/app/contracts/driving/api/controllers/suppliercontr"
	"github.com/cable_management/cable_be/app/contracts/driving/api/dtos"
	"github.com/cable_management/cable_be/app/contracts/driving/api/validations"
	"github.com/cable_management/cable_be/app/domain/services"
	"github.com/cable_management/cable_be/app/usecases/admincase"
	"github.com/cable_management/cable_be/app/usecases/commomcase"
	contractorcase "github.com/cable_management/cable_be/app/usecases/contractor"
	"github.com/cable_management/cable_be/app/usecases/plannercase"
	suppliercase "github.com/cable_management/cable_be/app/usecases/supplier"
	"github.com/cable_management/cable_be/driven/database"
	imrepos "github.com/cable_management/cable_be/driven/database/repos"
	imemail "github.com/cable_management/cable_be/driven/email"
	imsse "github.com/cable_management/cable_be/driven/sse"
	imadmincontr "github.com/cable_management/cable_be/driving/api/controllers/admincontr"
	imcommoncontr "github.com/cable_management/cable_be/driving/api/controllers/commoncontr"
	imcontractorcontr "github.com/cable_management/cable_be/driving/api/controllers/contractorcontr"
	implannercontr "github.com/cable_management/cable_be/driving/api/controllers/plannercontr"
	imsuppliercontr "github.com/cable_management/cable_be/driving/api/controllers/suppliercontr"
	"github.com/cable_management/cable_be/driving/api/routers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log"
)

func main() {

	engine := gin.Default()

	BuildEnv()

	StartDb()

	StartEmail()

	BuildValidator()

	// TODO nah I'm stupid as f*ck, this is a circular dependency bug
	BuildSSEServer(engine)

	BuildDomain()

	BuildApi(engine)

	// things above will introduce circular dependency

	err := engine.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
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

	environments.DbDsn = "host=db user=postgres password=postgrespw dbname=cable_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"

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
	notificationRepo   repos.INotificationRepo
)

func StartDb() {
	db = database.Init(environments.DbDsn)
	userRepo = imrepos.NewUserRepo(db)
	contractRepo = imrepos.NewContractRepo(db)
	requestRepo = imrepos.NewRequestRepo(db)
	requestHistoryRepo = imrepos.NewRequestHistoryRepo(db)
	notificationRepo = imrepos.NewNotificationRepo(db)
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
	notificationFac   services.INotificationFactory

	// usecases
	createUserCase               admincase.ICreateUser
	signInCase                   commomcase.ISignIn
	updateUserIsActiveCase       admincase.IUpdateUserIsActive
	createRequestCase            plannercase.ICreateRequest
	updateRequestStatusCase      commomcase.IUpdateRequestStatus
	getRequestListCase           plannercase.IGetRequestList
	getContractListCase          plannercase.IGetContractList
	getUserListCase              commomcase.IGetUserList
	getSupplierRequestListCase   suppliercase.IGetSupplierRequestList
	getSupplierContractListCase  suppliercase.IGetSupplierContractList
	getContractorRequestListCase contractorcase.IGetContractorRequestList
	getNotificationListCase      commomcase.IGetNotificationList
	updateNotificationIsReadCase commomcase.IUpdateNotificationIsRead
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
	notificationFac = services.NewNotificationFactory(userRepo, requestHistoryRepo)

	// usecases
	createUserCase = admincase.NewCreateUser(userRepo, userFac, validation, authorService, emailDriven, passwordService, notificationFac, notificationRepo, sseDriven)
	signInCase = commomcase.NewSignIn(userRepo, tokenService, passwordService)
	updateUserIsActiveCase = admincase.NewUpdateUserIsActive(userRepo, authorService, emailDriven, notificationFac, notificationRepo, sseDriven)
	createRequestCase = plannercase.NewCreateRequest(authorService, validation, requestFac, requestRepo, requestHistoryFac, mailDataFac, requestHistoryRepo, emailDriven, notificationFac, notificationRepo, sseDriven)
	updateRequestStatusCase = commomcase.NewUpdateRequestStatus(authorService, requestRepo, requestHistoryRepo, userRepo, validation, requestHistoryFac, mailDataFac, emailDriven, notificationFac, notificationRepo, sseDriven)
	getRequestListCase = plannercase.NewGetRequestList(requestRepo, authorService)
	getContractListCase = plannercase.NewGetContractList(authorService, contractRepo)
	getUserListCase = commomcase.NewGetUserList(authorService, userRepo)
	getSupplierRequestListCase = suppliercase.NewGetSupplierRequestList(requestRepo, authorService)
	getSupplierContractListCase = suppliercase.NewGetSupplierContractList(contractRepo, authorService)
	getContractorRequestListCase = contractorcase.NewGetContractorRequestList(requestRepo, authorService)
	getNotificationListCase = commomcase.NewGetNotificationList(notificationRepo, authorService)
	updateNotificationIsReadCase = commomcase.NewUpdateNotificationIsRead(authorService, notificationRepo)
}

// gin
var ()

// api
var (
	// controllers
	authContr              commoncontr.IAuthController
	adminUserContr         admincontr.IUserController
	plannerRequestContr    plannercontr.IRequestContr
	commonRequestContr     commoncontr.IRequestController
	plannerContractContr   plannercontr.IContractContr
	commonUserContr        commoncontr.IUserContr
	supplierRequestContr   suppliercontr.IRequestContr
	supplierContractContr  suppliercontr.IContractContr
	contractorRequestContr contractorcontr.IRequestContr
	notificationContr      commoncontr.INotificationContr

	// routers
	baseRouters       routers.IRouterBase
	commonRouters     routers.IRouterBase
	adminRouters      routers.IRouterBase
	plannerRouters    routers.IRouterBase
	supplierRouters   routers.IRouterBase
	contractorRouters routers.IRouterBase
)

func BuildApi(engine *gin.Engine) {

	// controllers
	authContr = imcommoncontr.NewAuthController(signInCase)
	adminUserContr = imadmincontr.NewUserController(createUserCase, updateUserIsActiveCase)
	plannerRequestContr = implannercontr.NewRequestContr(createRequestCase, getRequestListCase)
	commonRequestContr = imcommoncontr.NewRequestController(updateRequestStatusCase)
	plannerContractContr = implannercontr.NewContractContr(getContractListCase)
	commonUserContr = imcommoncontr.NewUserContr(getUserListCase)
	supplierRequestContr = imsuppliercontr.NewRequestContr(getSupplierRequestListCase)
	supplierContractContr = imsuppliercontr.NewContractContr(getSupplierContractListCase)
	contractorRequestContr = imcontractorcontr.NewRequestContr(getContractorRequestListCase)
	notificationContr = imcommoncontr.NewNotificationContr(getNotificationListCase, updateNotificationIsReadCase)

	// routers
	baseRouters = routers.NewRouterBase()
	commonRouters = routers.NewCommonRouters(authContr, commonRequestContr, commonUserContr, notificationContr)
	adminRouters = routers.NewAdminRouters(adminUserContr)
	plannerRouters = routers.NewPlannerRouters(plannerRequestContr, plannerContractContr)
	supplierRouters = routers.NewSupplierRouters(supplierRequestContr, supplierContractContr)
	contractorRouters = routers.NewContractorRouters(contractorRequestContr)

	apiRouters := engine.Group("/api")

	baseRouters.Register(apiRouters)
	commonRouters.Register(apiRouters)
	adminRouters.Register(apiRouters)
	plannerRouters.Register(apiRouters)
	supplierRouters.Register(apiRouters)
	contractorRouters.Register(apiRouters)

	//engine.NoRoute(middlewares.HandleGlobalErrors)
}

// sse
var (
	sseDriven sse.ISSEDriven
	sseServer imsse.ISSEServer
)

func BuildSSEServer(engine *gin.Engine) {
	sseServer = imsse.NewSSEServer(environments)
	sseDriven = imsse.NewSSEDriven(sseServer)
	sseServer.Register(engine)
}
