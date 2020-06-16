package routers

import (
	"auuuth/controllers"

	"github.com/gin-gonic/gin"
)

// NewServer create a new gin instance with routers
func NewServer() *gin.Engine {
	e := gin.New()

	// -----define routres here------
	v1 := e.Group("/v1")
	{
		userGroup := v1.Group("/user")
		{
			userGroup.POST("/add", controllers.AddUser)
			userGroup.POST("/delete", controllers.DeleteUSer)
		}

		roleGroup := v1.Group("/role")
		{
			roleGroup.POST("/add", controllers.AddRole)
			roleGroup.POST("/delete", controllers.DeleteRole)
		}

		resourceGroup := v1.Group("/resource")
		{
			resourceGroup.POST("/add", controllers.AddResource)
			resourceGroup.POST("/delete", controllers.DeleteResource)
		}

		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/start", controllers.StartAuth)
			authGroup.POST("/verify", controllers.VerifyAuth)
			// authGroup.POST("/end")
		}
	}

	// -----end routers define-------

	return e
}
