package main

import (
	config "blog/config"
	"blog/internal/delivery"
	"blog/internal/infra/db"
	"blog/internal/infra/repository"
	"blog/internal/usecase"
	Middleware "blog/middleware"

	"flag"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()

	config.Init()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Debug().Msg("This message appears only when log level set to Debug")
	log.Info().Msg("This message appears when log level set to Debug or Info")

	var err error

	db, err := db.NewDb()
	if err != nil {
		log.Fatal().Msg("database connect fail")
	}

	repo := repository.NewRepository(db)

	uc := usecase.NewUsecase(repo)

	handler := delivery.NewHandler(uc)

	server := gin.Default()
	server.Use(Middleware.CORSMiddleware())

	// authorized := server.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"abc": "123",
	// }))

	server.GET("/info", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "user.info"})
	})

	server.GET("/")

	story := server.Group("/story")
	{
		story.GET("/", delivery.GetAll)
		story.GET("/:title", Story.GetByTitle)
		story.POST("/", Story.Create)
		story.DELETE("/:title", Story.Delete)
		story.PUT("/:title", Story.Update)
		story.GET("/filter", Story.GetByFilter)
	}

	server.Run(":8080")
}
