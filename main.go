package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sachinadate/learnginframeworkp1/controller"
	"github.com/sachinadate/learnginframeworkp1/middleware"
	"github.com/sachinadate/learnginframeworkp1/repository"
	"github.com/sachinadate/learnginframeworkp1/service"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/swag/example/basic/docs"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	videoController controller.VideoController = controller.New(videoService)
	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

//@securityDefinitions.apikey bearer auth
//@in header
//@name Authorization

func main() {

	//Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Pragmatic Reviews -Vedio API"
	docs.SwaggerInfo.Description = "Pragmatic Reviews -Youtube Vedio API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	defer videoRepository.CloseDB()
	setupLogOutput()
	// server := gin.Default()

	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger()) // middleware.BasicAuth(), gindump.Dump()
	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("./templates/*.html")

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			fmt.Println("post method called")
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	// server.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"Message": "OK!!",
	// 	})
	// })

	apiRoutes := server.Group("/api", middleware.AuthorizeJWT())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "video input is valid",
				})
			}
			ctx.JSON(200, videoController.Save(ctx))
		})

		apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "video input is valid",
				})
			}
		})

		apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": fmt.Errorf("error in server.go"),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "success",
				})
			}
		})

	}

	server.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(":5000")
}

//    curl "localhost:8081/login" -x POST -d '{"username":"pragmatic","password":"reviews"}'
