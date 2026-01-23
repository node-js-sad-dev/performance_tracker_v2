package main

import (
	"log"
	"performance_tracker_v2_be/config"
	"performance_tracker_v2_be/db"
	"performance_tracker_v2_be/modules"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Performance Tracker V2 API
// @version			1.0
// @description		Performance Tracker, advanced version. Pet project
// @BasePath		/api/v1
func main() {
	cfg := config.Load()

	// investigate what it is, and what are other modes
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	app := gin.Default()

	databaseConnections, err := db.InitializeDatabases(cfg)
	if err != nil {
		log.Fatalf("failed to initialize databases: %v", err)
	}

	defer databaseConnections.Close()

	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	modules.RootRouter(app)

	runErr := app.Run(cfg.GetServerAddr())

	if runErr != nil {
		panic(runErr)
	}
}
