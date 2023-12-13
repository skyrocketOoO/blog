package delivery

import (
	"blog/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StoryHandler struct {
	storyUsecase *usecase.StoryUsecase
}

func NewStoryHandler(router *gin.RouterGroup, storyUsecase *usecase.StoryUsecase) {
	handler := &StoryHandler{
		storyUsecase: storyUsecase,
	}
	router.GET("/", handler.GetAll)
	router.GET("/:title", handler.Get)
	router.POST("/", handler.Create)
	router.DELETE("/:title", handler.Delete)
	router.PUT("/:title", handler.Update)
}

func (s *StoryHandler) GetAll(ctx *gin.Context) {
	reqContext := ctx.Request.Context()
	if stories, err := s.storyUsecase.GetAll(&reqContext); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"context": stories})
	}
}

func (s *StoryHandler) Get(ctx *gin.Context) {
	idStr := ctx.Query("id")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format: " + err.Error()})
		return
	}
	reqContext := ctx.Request.Context()
	story, err := s.storyUsecase.Get(&reqContext, uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to retrieve story: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"context": story})
}

func (s *StoryHandler) Create(ctx *gin.Context) {
	var reqBody struct {
		title   string
		content string
		label   string
		isDraft bool
	}
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reqContext := ctx.Request.Context()
	if id, err := s.storyUsecase.Create(&reqContext, reqBody.title, reqBody.content, reqBody.label, reqBody.isDraft); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	}
}

func (s *StoryHandler) Update(ctx *gin.Context) {
	idStr := ctx.Query("id")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format: " + err.Error()})
		return
	}
	var reqBody struct {
		title   string
		content string
		label   string
		isDraft bool
	}
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reqContext := ctx.Request.Context()
	if err = s.storyUsecase.Update(&reqContext, uint(id), reqBody.title, reqBody.content, reqBody.label, reqBody.isDraft); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

func (s *StoryHandler) Delete(ctx *gin.Context) {
	idStr := ctx.Query("id")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format: " + err.Error()})
		return
	}
	reqContext := ctx.Request.Context()
	if err = s.storyUsecase.Delete(&reqContext, uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

// func (s *StoryHandler) GetByFilter(ctx *gin.Context) {
// 	filters := make(map[string]interface{})
// 	for key, val := range ctx.Request.URL.Query() {
// 		filters[key] = val
// 	}

// 	stories, err := s.storyUsecase.GetByFilter(filters)
// 	status := http.StatusOK
// 	if err != nil {
// 		status = http.StatusBadRequest
// 	}
// 	ctx.JSON(status, gin.H{
// 		"context": stories,
// 		"error":   err,
// 	})
// }
