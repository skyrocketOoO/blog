package main

import (
	AdminStory "blog/app/admin/story"
	Story "blog/app/story"
	Config "blog/config"
	Db "blog/db"
	Middleware "blog/middleware"
	StoryOrm "blog/model/story"
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
	Db.Db.AutoMigrate(&StoryOrm.Story{})

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

	server.GET("/")

	story := server.Group("/story")
	{
		story.GET("/", Story.GetAll)
		story.GET("/:title", Story.GetByTitle)
	}

	admin := authorized.Group("/admin")
	{
		admin.GET("/")

		story := admin.Group("/story")
		{
			story.GET("/", AdminStory.GetAll)
			story.GET("/:title", AdminStory.GetByTitle)
			story.POST("/", AdminStory.Create)
			story.DELETE("/:title", AdminStory.Delete)
			story.PUT("/:title", AdminStory.Update)
			story.GET("/key", AdminStory.GetByKey)
		}

		author := admin.Group("/author")
		{
			author.GET("/")
			author.GET("/:name")
			author.POST("/")
			author.PUT("/:name")
			author.DELETE("/:name")
		}

	}

	server.Run(":8080")
}
