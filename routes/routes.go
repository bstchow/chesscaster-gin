package routes

import (
	"chesscaster-gin/controllers"
	"chesscaster-gin/helper"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the router.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	accounts := gin.Accounts{
		"admin": helper.GetEnv("ADMIN_PASSWORD", "123"),
	}

	games := r.Group("/games", gin.BasicAuth(accounts))

	{
		games.GET("/active", controllers.ActiveGamesOfUser)
		games.POST("/", controllers.CreateGame)
		games.PATCH("/:id", controllers.PatchGame)
	}

	return r
}
