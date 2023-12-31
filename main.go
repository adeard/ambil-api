package main

import (
	"ambil-api/config"
	"ambil-api/docs"
	"ambil-api/modules/category"
	"ambil-api/modules/merchant"
	"ambil-api/modules/order"
	"ambil-api/modules/user"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @Host localhost:6969
// @title API SWAGGER FOR AMBIL API SERVICE
// @version 1.0.0
// @description AMBIL API SERVICE
// @termsOfService http://swagger.io/terms/

// @contact.name ADE ARDIAN
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath

func main() {
	db := config.Connect()

	docs.SwaggerInfo.BasePath = "/ambil"

	router := gin.Default()
	router.Use(cors.AllowAll())
	router.GET("ambil/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"title":         "Ambil API Service",
			"documentation": "/swagger/index.html",
		})
	})

	router.GET("ambil/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := router.Group("ambil/api/v1")

	user.NewUserHandler(v1, user.UserRegistry(db))
	order.NewOrderHandler(v1, order.OrderRegistry(db))
	category.NewCategoryHandler(v1, category.CategoryRegistry(db))
	merchant.NewMerchantHandler(v1, merchant.MerchantRegistry(db))

	// router.Run(":86")

	// Mengatur mode GIN menjadi release
	gin.SetMode(gin.ReleaseMode)

	//Penyesuaian Port ke IIS
	port := "86"
	if os.Getenv("ASPNETCORE_PORT") != "" {
		port = os.Getenv("ASPNETCORE_PORT")
	}

	// Menampilkan log koneksi sukses
	log.Println("App Service run in port:", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		// Menampilkan log ketika koneksi gagal
		log.Fatal("Connection Fail -> port "+port+":", err)
	}
}
