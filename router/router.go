package router

import (
	"CRUD-GIN/controllers"
	"CRUD-GIN/middlewares"

	"github.com/gin-gonic/gin"
)

func RouterSetUp() *gin.Engine {
	router := gin.Default()

	public := router.Group("/")
	{
		public.POST("signup", controllers.SignUp)
		public.POST("login", controllers.Login)
	}

	protected := router.Group("/protected")
	protected.Use(middlewares.JWTMiddleware())
	{
		protected.GET("/users/:userid/tasks", controllers.Tasks)
		protected.GET("/users/:userid/tasks/:taskid", controllers.Task)
		protected.POST("/users/:userid/task", controllers.Create)
		protected.DELETE("/users/:userid/tasks/:taskid", controllers.Delete)
		protected.PATCH("/users/:userid/tasks/:taskid", controllers.Update)
	}

	return router
}
