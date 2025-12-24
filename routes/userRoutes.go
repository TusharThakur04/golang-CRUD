package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tusharthakur04/golang-CRUD/handlers"
)

type UserRoutes struct{}



func (u *UserRoutes) Register(router *gin.RouterGroup){

	h := &handlers.UserHandler{}


	router.GET("/" ,h.GetUsers)
	router.GET("/:id", h.GetUser)
	router.POST("/", h.CreateUser)
	router.PUT("/:id", h.UpdateUser)
	router.DELETE("/:id", h.DeleteUser)

}




