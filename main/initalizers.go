package main

import (
	"github.com/cable_management/cable_be/_share/env"
	"github.com/cable_management/cable_be/app/contracts/database/repos"
	"github.com/cable_management/cable_be/app/domain/services"
	"github.com/cable_management/cable_be/app/usecases/admincase"
	"github.com/cable_management/cable_be/app/usecases/commomcase"
	"github.com/cable_management/cable_be/driven/database"
	imRepos "github.com/cable_management/cable_be/driven/database/repos"
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

// domain
var (
	// services
	passwordService services.IPasswordService
	userFac         services.IUserFactory
	tokenService    services.IAuthTokenService

	// usecases
	createUserCase admincase.ICreateUser
	signInCase     commomcase.ISignIn
)

// api
var (
	// controllers
	authContr commoncontr.IAuthController

	// routers
	commonRouters *routers.CommonRouters
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
}

func StartDb() {
	db = database.Init(environments.DbDsn)
	userRepo = imRepos.NewUserRepo(db)
}

func BuildValidator() {

	validation = validator.New()

	vlCreateUserDepd = admincase.NewVlCreateUserDepd(userRepo)

	validation.RegisterStructValidation(vlCreateUserDepd.Handle, &admincase.CreateUserReq{})
}

func BuildDomain() {

	// services
	passwordService = services.NewPasswordHash(validation)
	userFac = services.NewUserFactory(passwordService, userRepo, validation)
	tokenService = services.NewAuthTokenService(environments.JwtSecret)

	// usecases
	createUserCase = admincase.NewCreateUser(userRepo, userFac, validation)
	signInCase = commomcase.NewSignIn(userRepo, tokenService, passwordService)
}

func StartApi() {

	// controllers
	authContr = commoncontr.NewAuthController(signInCase)

	// routers
	commonRouters = routers.NewCommonRouters(authContr)

	// init
	engine := gin.Default()

	commonRouters.Register(engine)

	//engine.NoRoute(middlewares.HandleGlobalErrors)

	err := engine.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
