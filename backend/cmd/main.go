package main

import (
	"blog/config"
	"blog/internal/delivery"
	"blog/internal/delivery/middleware"
	"blog/internal/infra/db"
	"blog/internal/infra/repository"
	"blog/internal/usecase"

	"flag"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	config.Init()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()
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
	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	delivery.NewHandler(server, uc)

	// authorized := server.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"abc": "123",
	// }))

	server.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	server.Run(":8080")
}
