package main

import (
	"github.com/cable_management/cable_be/_share/env"
	"github.com/cable_management/cable_be/app/contracts/database/repos"
	"github.com/cable_management/cable_be/app/contracts/email"
	"github.com/cable_management/cable_be/app/domain/services"
	"github.com/cable_management/cable_be/app/usecases/admincase"
	"github.com/cable_management/cable_be/app/usecases/commomcase"
	"github.com/cable_management/cable_be/driven/database"
	imRepos "github.com/cable_management/cable_be/driven/database/repos"
	imEmail "github.com/cable_management/cable_be/driven/email"
	"github.com/cable_management/cable_be/driving/api/controllers/admincontr"
	"github.com/cable_management/cable_be/driving/api/controllers/commoncontr"
	"github.com/cable_management/cable_be/driving/api/routers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log"
)

var environments env.Env

var (
	db       *gorm.DB
	userRepo repos.IUserRepo
)

var (
	validation       *validator.Validate
	vlCreateUserDepd *admincase.VlCreateUserDepd
)

var (
	emailDriven email.IEmail
)

// domain
var (
	// services
	passwordService services.IPasswordService
	userFac         services.IUserFactory
	tokenService    services.IAuthTokenService
	authorService   services.IAuthorizeService

	// usecases
	createUserCase admincase.ICreateUser
	signInCase     commomcase.ISignIn
)

// api
var (
	// controllers
	authContr      commoncontr.IAuthController
	adminUserContr admincontr.IUserController

	// routers
	commonRouters routers.IRouterBase
	adminRouters  routers.IRouterBase
)

func BuildEnv() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//
	//ENV.DbDsn = os.Getenv("DB_DSN")

	environments.JwtSecret = "123467890qwertuiopasdfghjkl"

	environments.DbDsn = "host=localhost user=postgres password=postgrespw dbname=cable_db port=32768 sslmode=disable TimeZone=Asia/Shanghai"

	environments.SmtpEmail = "vuphamlethanh@gmail.com"
	environments.SmtpHost = "smtp.gmail.com"
	environments.SmtpPort = "587"
	environments.SmtpPassword = "fzjugwhesxnbpixp"
}

func StartDb() {
	db = database.Init(environments.DbDsn)
	userRepo = imRepos.NewUserRepo(db)
}

func StartEmail() {
	emailDriven = imEmail.NewEmail(imEmail.EmailConfig{
		MailHost: environments.SmtpEmail,
		Host:     environments.SmtpHost,
		Port:     environments.SmtpPort,
		Password: environments.SmtpPassword,
	})
}

func BuildValidator() {

	validation = validator.New()

	vlCreateUserDepd = admincase.NewVlCreateUserDepd(userRepo)

	validation.RegisterStructValidation(vlCreateUserDepd.Handle, admincase.CreateUserReq{})
}

func BuildDomain() {

	// services
	passwordService = services.NewPasswordHash(validation)
	userFac = services.NewUserFactory(passwordService, userRepo, validation)
	tokenService = services.NewAuthTokenService(environments.JwtSecret)
	authorService = services.NewAuthorizeService(tokenService, userRepo)

	// usecases
	createUserCase = admincase.NewCreateUser(userRepo, userFac, validation, authorService)
	signInCase = commomcase.NewSignIn(userRepo, tokenService, passwordService)
}

func StartApi() {

	// controllers
	authContr = commoncontr.NewAuthController(signInCase)
	adminUserContr = admincontr.NewUserController(createUserCase)

	// routers
	commonRouters = routers.NewCommonRouters(authContr)
	adminRouters = routers.NewAdminRouters(adminUserContr)

	// init
	engine := gin.Default()

	commonRouters.Register(engine)
	adminRouters.Register(engine)

	//engine.NoRoute(middlewares.HandleGlobalErrors)

	err := engine.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
