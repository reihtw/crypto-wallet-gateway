package handlers

import (
    "github.com/reihtw/crypto-wallet-gateway/config"
    "github.com/gin-gonic/gin"
    "net/http"
)

func RegisterUser(c *gin.Context) {
    userServiceURL := config.GetEnv("USER_SERVICE_URL") + "/register"
    resp, err := http.Post(userServiceURL, "application/json", c.Request.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer resp.Body.Close()

    c.JSON(resp.StatusCode, gin.H{"message": "User registered successfully"})
}

func LoginUser(c *gin.Context) {
    userServiceURL := config.GetEnv("USER_SERVICE_URL") + "/login"
    resp, err := http.Post(userServiceURL, "application/json", c.Request.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer resp.Body.Close()

    c.JSON(resp.StatusCode, gin.H{"message": "User logged in successfully"})
}

func GetBalance(c *gin.Context) {
    userId := c.Param("userId")
    balanceServiceURL = config.GetEnv("BALANCE_SERVICE_URL") + "/balance/" + userId

    resp, err := http.Get(balanceServiceURL)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error", err.Error()})
        return
    }
    defer resp.Body.Close()

    c.JSON(resp.StatusCode, gin.H{"message": "Balance retrieved successfully"})
}

func CreateTransaction(c *gin.Context) {
    transactionServiceURL := config.GetEnv("TRANSACTION_SERVICE_URL") + "/transaction"
    resp, err := http.Post(transactionServiceURL, "application/json", c.Request.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer resp.Body.Close()

    c.JSON(resp.StatusCode, gin.H{"message": "Transaction created successfully"})
}

