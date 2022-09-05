package main

import (
	"auth-service/config"
	"auth-service/migrations"
	"auth-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitViperConfig()
	config.InitDatabase()
	migrations.Migrate()
	r := gin.Default()
	routes.Serve(r)
	// r.Run(":" + viper.GetString("app.port"))
	r.Run()

}
