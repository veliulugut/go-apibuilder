package router

import (
	"go-apibuilder/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(apiGroup *gin.RouterGroup, userHandler *handler.UserHandler) {
	userRoutes := apiGroup.Group("/users")
	{
		userRoutes.POST("", userHandler.CreateUser)
		userRoutes.GET("/:id", userHandler.GetUserByID)
	}
}
