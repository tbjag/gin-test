package main

import (
	"example/web-service-gin/controller"
	"example/web-service-gin/middleware"
	"example/web-service-gin/repository"
	"example/web-service-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	recipeRepository repository.RecipeRepository = repository.NewRecipeRepository()
	recipeService    service.RecipeService       = service.New(recipeRepository)
	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()

	recipeController controller.RecipeController = controller.New(recipeService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func main() {
	defer recipeRepository.CloseDB()

	server := gin.New()

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), gin.Logger())

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

// JWT Authorization Middleware applies to "/api" only.
apiRoutes := server.Group("/api", middleware.AuthorizeJWT())
{
	apiRoutes.GET("/recipes", func(ctx *gin.Context) {
		ctx.JSON(200, recipeController.FindAll())
	})

	apiRoutes.POST("/recipes", func(ctx *gin.Context) {
		err := recipeController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
		}

	})

	apiRoutes.PUT("/recipes/:id", func(ctx *gin.Context) {
		err := recipeController.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
		}

	})

	apiRoutes.DELETE("/recipes/:id", func(ctx *gin.Context) {
		err := recipeController.Delete(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
		}

	})

}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/recipes", recipeController.ShowAll)
	}

	server.Run(":8080")
}
