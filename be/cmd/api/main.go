package main

import (
	"log"
	"performance_tracker_v2_be/config"
	"performance_tracker_v2_be/db"
	main_db "performance_tracker_v2_be/db/main-db"
	"performance_tracker_v2_be/docs"
	"performance_tracker_v2_be/middlewares"
	"performance_tracker_v2_be/modules"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Performance Tracker V2 API
// @version		1.0
// @description	Performance Tracker, advanced version. Pet project
// @BasePath		/api/v1
func main() {
	cfg := config.Load()

	databaseConnections, err := db.InitializeDatabases(cfg)
	if err != nil {
		log.Fatalf("failed to initialize databases: %v", err)
	}

	println("Databases initialized successfully")

	defer databaseConnections.Close()

	if err := main_db.RunMigrations(databaseConnections.MainDatabase); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	println("Database migrations ran successfully")

	// swagger to set correct host, not hardcoded
	docs.SwaggerInfo.Host = cfg.HOST

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	app := gin.Default()
	app.Use(middlewares.CORSMiddleware(cfg))

	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	modules.RootRouter(cfg, databaseConnections.MainDatabase, app)

	println("Initialized routers successfully")

	serverAddress := cfg.GetServerAddr()

	runErr := app.Run(serverAddress)

	if runErr != nil {
		panic(runErr)
	}

	println("Server running on " + serverAddress)
}
