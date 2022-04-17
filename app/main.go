package main

import (
	_middleware "rest-go-stock/app/middlewares" 
	"rest-go-stock/app/routes"
	"rest-go-stock/drivers/mysql"
	"log"
	"time"

	userUseCase "rest-go-stock/businesses/users"
	userController "rest-go-stock/controllers/users"
	userRepo "rest-go-stock/drivers/repository/users"

	product_typeUseCase "rest-go-stock/businesses/product_types"
	product_typeController "rest-go-stock/controllers/product_types"
	product_typeRepo "rest-go-stock/drivers/repository/product_types"

	productUseCase "rest-go-stock/businesses/products"
	productController "rest-go-stock/controllers/products"
	productRepo "rest-go-stock/drivers/repository/products"

	
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service Run on Debug mode")
	}
}

func DBMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&userRepo.User{})
	DB.AutoMigrate(&product_typeRepo.ProductType{})
	DB.AutoMigrate(&productRepo.Product{})
}



func main() {
	ConfigDB := mysql.ConfigDB{
		DB_Username: viper.GetString("database.user"),
		DB_Password: viper.GetString("database.pass"),
		DB_Host:     viper.GetString("database.host"),
		DB_Port:     viper.GetString("database.port"),
		DB_Database: viper.GetString("database.name"),
	}

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString("JWT.secretKey"),
		ExpiresDuration: viper.GetInt("JWT.expired_time"),
	}

	DB := ConfigDB.InitialDB()
	DBMigrate(DB)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	e := echo.New()
	_middleware.Log(e)
	e.Use(middleware.CORS())
	userRepoInterface := userRepo.NewUserRepo(DB)
	userUseCaseInterface := userUseCase.NewUseCase(userRepoInterface, timeoutContext, &configJWT)
	userUseControllerInterface := userController.NewUserController(userUseCaseInterface)

	product_typeRepoInterface := product_typeRepo.NewProductTypeRepo(DB)
	product_typeUseCaseInterface := product_typeUseCase.NewUseCase(product_typeRepoInterface, timeoutContext)
	product_typeUseControllerInterface := product_typeController.NewProductTypeController(product_typeUseCaseInterface)

	productRepoInterface := productRepo.NewProductRepo(DB)
	productUseCaseInterface := productUseCase.NewUseCase(productRepoInterface, timeoutContext)
	productUseControllerInterface := productController.NewProductController(productUseCaseInterface)



	routesInit := routes.RouteControllerList{
		UserController: *userUseControllerInterface,
		ProductTypeController: *product_typeUseControllerInterface,
		ProductController: *productUseControllerInterface,
		JWTMiddleware: configJWT.Init(),
	}
	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))

}