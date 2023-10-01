package router

import (
	photocontroller "github.com/davethio/task-5-pbi-btpns-DaveChristianThio/controllers/photo_controller"
	usercontroller "github.com/davethio/task-5-pbi-btpns-DaveChristianThio/controllers/user_controller"
	"github.com/davethio/task-5-pbi-btpns-DaveChristianThio/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	router := r

	// User route
	userGroup := router.Group("/users")
	userGroup.POST("/register", usercontroller.RegisterUser)
	userGroup.POST("/login", usercontroller.LoginUser)
	userGroup.GET("/logout", usercontroller.Logout)
	userGroup.DELETE("/:id", usercontroller.DeleteUser)
	userGroup.PUT("/:id", usercontroller.UpdateUser)

	// Photo route
	photoGroup := router.Group("/photos")
	photoGroup.Use(middlewares.JWTMiddleware())
	{
		photoGroup.GET("/", photocontroller.GetPhotos)
		photoGroup.POST("/", photocontroller.CreatePhoto)
		photoGroup.PUT("/:photoId", photocontroller.UpdatePhoto)
		photoGroup.DELETE("/:photoId", photocontroller.DeletePhoto)
	}
}
