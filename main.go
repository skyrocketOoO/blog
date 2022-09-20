package main

import (
	Story "blog/app/story"
	Config "blog/config"
	Db "blog/db"
	Middleware "blog/middleware"
	StoryMd "blog/model/story"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	var err error

	// init db
	Db.Db, err = Db.InitDb(Config.Dsn)
	if err != nil {
		fmt.Println("database connect fail")
	}
	Db.Db.AutoMigrate(&StoryMd.Story{})

	server := gin.Default()

	// cors
	server.Use(Middleware.CORSMiddleware())

	//auth
	authorized := server.Group("/", gin.BasicAuth(gin.Accounts{
		"abc": "123",
	}))

	authorized.GET("/info", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "user.info"})
	})

	story := authorized.Group("/story")
	{
		story.GET("/", Story.GetAll)
		story.GET("/:title", Story.GetByTitle)
		story.POST("/", Story.Create)
		story.DELETE("/:title", Story.Delete)
		story.PUT("/:title", Story.Update)
	}

	server.Run(":8080")
}
