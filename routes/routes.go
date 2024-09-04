package routes

import (
    "github.com/reihtw/crypto-wallet-gateway/handlers"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()
    
    router.POST("/user/register", handlers.RegisterUser)
    router.POST("/user/login", handlers.LoginUser)

    router.GET("/balance/:userId", handlers.GetBalance)

    router.POST("/transaction", handlers.CreateTransaction)

    return router
}

