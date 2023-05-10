package main

import (
	"github.com/gin-gonic/gin"
	tokenroutes "userservice-go/routes/token-routes"
	userroutes "userservice-go/routes/user-routes"
	"userservice-go/types"
)

// InitializeAndStartServer Initializes and starts the server
func InitializeAndStartServer() {
	server := gin.Default()
	initializeRoutes(*server)
	startServer(*server)
}

func initializeRoutes(server gin.Engine) {
	server.GET("/users", userroutes.GetUsersByUsersCriteria)
	server.POST("/token", tokenroutes.GetTokenWithPasswordGrant)
}

func startServer(server gin.Engine) {
	server.Run(types.USER_SERVICE_PORT)
}