package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"upscape/controllers"
	"upscape/middlewares"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// load env file
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Starting the application")

	// initilise database

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// connect to DB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err.Error())
	}
	// set database
	clientDatabase := client.Database(os.Getenv("DATABASE"))
	// when disconnected

	fmt.Println("Connected to Atlas successfully")
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// instantiate router
	router := gin.Default()


	// define go logger
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("[%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())

	// define middlewares
	router.Use(middlewares.CORSMiddleware())

	// declare routes

	// if app is in production
	if os.Getenv("GIN_MODE") == "release" {
		// handle production
		fmt.Println(os.Getenv("GIN_MODE"))
		// static folder
		// Serve the frontend
		router.Use(static.Serve("/", static.LocalFile("./public", true)))
	}

	router.POST("/api/auth/login", controllers.LoginUser(clientDatabase))
	// user routes
	router.POST("/api/user", controllers.CreateUser(clientDatabase))
	router.PUT("/api/user", middlewares.IsAuthenticated, controllers.UpdateUser(clientDatabase))
	router.PUT("/api/user/password", middlewares.IsAuthenticated, controllers.UpdateUserPassword(clientDatabase))
	router.GET("/api/user", middlewares.IsAuthenticated, controllers.GetUserMeta(clientDatabase))
	// task routes
	router.GET("/api/task", middlewares.IsAuthenticated, controllers.GetTasks(clientDatabase))
	router.POST("/api/task", middlewares.IsAuthenticated, controllers.CreateTask(clientDatabase))
	router.PUT("/api/task/:id", middlewares.IsAuthenticated, controllers.UpdateTask(clientDatabase))
	router.DELETE("/api/task/:id", middlewares.IsAuthenticated, controllers.DeleteTask(clientDatabase))
	// objective routes
	router.GET("/api/objective", middlewares.IsAuthenticated, controllers.GetObjectives(clientDatabase))
	router.POST("/api/objective", middlewares.IsAuthenticated, controllers.CreateObjective(clientDatabase))
	router.PUT("/api/objective/:id", middlewares.IsAuthenticated, controllers.UpdateObjective(clientDatabase))
	router.DELETE("/api/objective/:id", middlewares.IsAuthenticated, controllers.DeleteObjective(clientDatabase))
	// tag routes
	router.GET("/api/tag", middlewares.IsAuthenticated, controllers.GetTags(clientDatabase))
	router.POST("/api/tag", middlewares.IsAuthenticated, controllers.CreateTag(clientDatabase))
	router.PUT("/api/tag/:id", middlewares.IsAuthenticated, controllers.UpdateTag(clientDatabase))
	router.DELETE("/api/tag/:id", middlewares.IsAuthenticated, controllers.DeleteTag(clientDatabase))

	router.Run()
}
