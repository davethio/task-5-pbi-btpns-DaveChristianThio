package router

import (
	photocontroller "github.com/davethio/task-5-pbi-btpns-DaveChristianThio/controllers/photo_controller"
	usercontroller "github.com/davethio/task-5-pbi-btpns-DaveChristianThio/controllers/user_controller"

	"github.com/davethio/task-5-pbi-btpns-DaveChristianThio/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	router := r

	userGroup := router.Group("/users")
	userGroup.GET("/index", usercontroller.IndexUser)
	userGroup.GET("/:id", usercontroller.ShowUser)
	userGroup.POST("/register", usercontroller.RegisterUser)
	userGroup.POST("/login", usercontroller.LoginUser)
	userGroup.GET("/logout", usercontroller.Logout)
	userGroup.DELETE("/:id", usercontroller.DeleteUser)
	userGroup.PUT("/:id", usercontroller.UpdateUser)

	// Group the photo routes

	photoGroup := router.Group("/photos")
	photoGroup.Use(middlewares.JWTMiddleware())
	{
		photoGroup.GET("/", photocontroller.GetPhotos)
		photoGroup.POST("/create", photocontroller.CreatePhoto)
		photoGroup.PUT("/:photoId", photocontroller.UpdatePhoto)
		photoGroup.DELETE("/:photoId", photocontroller.DeletePhoto)
	}
}
