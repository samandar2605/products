package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/samandar2605/products/api/v1"
	"github.com/samandar2605/products/config"
	"github.com/samandar2605/products/storage"

	"github.com/gin-contrib/cors"

	_ "github.com/samandar2605/products/api/docs" // for swagger
	swaggerFiles "github.com/swaggo/files"        // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"    // gin-swagger middleware
)

type RouterOptions struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

// @title           Swagger for product api
// @version         1.0
// @description     This is a product service api.
// @BasePath  		/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "*")
	router.Use(cors.New(corsConfig))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Cfg:     opt.Cfg,
		Storage: opt.Storage,
	})
	router.Static("/media", "./media")
	apiV1 := router.Group("/v1")

	// // User
	// apiV1.POST("/users", handlerV1.CreateUser)
	apiV1.GET("/products", handlerV1.GetAllProducts)
	// apiV1.GET("/users/:id", handlerV1.GetUser)
	// apiV1.PUT("/users/:id", handlerV1.UpdateUser)
	// apiV1.DELETE("/users/:id", handlerV1.DeleteUser)

	// file upload
	apiV1.POST("/file-upload", handlerV1.UploadFile)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
