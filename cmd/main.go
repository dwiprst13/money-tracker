package main

import (
    "money-tracker/config"
    "money-tracker/routes"
)

func main() {
    config.ConnectDatabase()
    r := routes.SetupRouter()
    r.Run(":8080")
}
