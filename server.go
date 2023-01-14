package todo

import (
	"net/http"

	"github.com/Gavall/todo-app/pkg/handlers"
	"github.com/gin-gonic/gin"
)

// var hu = handlers.SignUp()
func Router(port string) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	auth := r.Group("/auth")
	{
		auth.POST("/signUp", handlers.SignUp())
		auth.POST("/signIn")
	}
	api := r.Group("/api")
	{
		list := api.Group("/list")
		{
			list.GET("/:id")
			list.GET("/")
			list.POST("/")
			list.PUT("/:id")
			list.DELETE("/:id")
		}
		item := api.Group("/item")
		{
			item.GET("/:id")
			item.GET("/")
			item.POST("/")
			item.PUT("/:id")
			item.DELETE("/:id")
		}
	}

	r.Run(":" + port)
}
