package bootstrap

import (
	"go-gin-be-clean-arch/mongo"
	"go-gin-be-clean-arch/redis"
)

type Application struct {
	Env   *Env
	Mongo mongo.Client
	Redis redis.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Redis = NewRedisCache(app.Env)
	app.Mongo = NewMongoDatabase(app.Env)
	return *app
}

func (app *Application) CloseAllConnection() {
	CloseMongoDBConnection(app.Mongo)
	CloseRedisConnection(app.Redis)
}
