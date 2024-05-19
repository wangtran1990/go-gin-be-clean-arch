package main

import (
	"time"

	route "go-gin-be-clean-arch/api/route"
	"go-gin-be-clean-arch/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	cache := app.Redis
	defer app.CloseAllConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, cache, gin)

	gin.Run(env.ServerAddress)
}
