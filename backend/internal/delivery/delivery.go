package delivery

import (
	"blog/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewHandler(server *gin.Engine, usecase *usecase.Usecase) {
	NewStoryHandler(server.Group("/story"), usecase.StoryUsecase)
}
