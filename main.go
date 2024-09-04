package main

import (
    "crypto-wallet-gateway/config"
    "crypto-wallet-gateway/routes"
    "log"
)

func main() {
    config.LoadConfig()

    router := routes.SetupRouter()

    log.Println("Starting the gateway service on port 8080...")
    log.Fatal(router.Run(":8080"))
}

