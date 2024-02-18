package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	//Initialize the route using the GIN default settings
	router := gin.Default()
	//Initialize routes
	initializeRoutes(router)
	//Run the server
	router.Run(":3000")
}
