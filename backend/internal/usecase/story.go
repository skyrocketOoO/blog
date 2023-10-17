package usecase

import (
	"blog/internal/infra/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) {
	Orm := repository.NewOrmStory()
	stories, err := Orm.GetAll()
	status := http.StatusOK
	if err != nil {
		status = http.StatusBadRequest
	}
	ctx.JSON(status, gin.H{
		"context": stories,
		"error":   err,
	})
}

func GetByTitle(ctx *gin.Context) {
	Orm := repository.NewOrmStory()
	story, err := Orm.GetByTitle(ctx.Param("title"))
	status := http.StatusOK
	if err != nil {
		status = http.StatusBadRequest
	}
	ctx.JSON(status, gin.H{
		"context": story,
		"error":   err,
	})
}

func Create(ctx *gin.Context) {
	Orm := repository.NewOrmStory()
	story := &repository.Story{
		Title:   ctx.PostForm("title"),
		Author:  ctx.PostForm("author"),
		Content: ctx.PostForm("content"),
		Label:   ctx.PostForm("label"),
	}
	err := Orm.Create(story)
	status := http.StatusOK
	if err != nil {
		status = http.StatusBadRequest
	}
	ctx.JSON(status, gin.H{
		"error": err,
	})
}

func Update(ctx *gin.Context) {
	Orm := repository.NewOrmStory()
	story := &repository.Story{
		Title:   ctx.PostForm("title"),
		Author:  ctx.PostForm("author"),
		Content: ctx.PostForm("content"),
		Label:   ctx.PostForm("label"),
	}
	err := Orm.Update(ctx.Param("title"), story)
	status := http.StatusOK
	if err != nil {
		status = http.StatusBadRequest
	}
	ctx.JSON(status, gin.H{
		"error": err,
	})
}

func Delete(ctx *gin.Context) {
	Orm := repository.NewOrmStory()
	err := Orm.Delete(ctx.Param("title"))
	status := http.StatusOK
	if err != nil {
		status = http.StatusBadRequest
	}
	ctx.JSON(status, gin.H{
		"error": err,
	})
}

func GetByFilter(ctx *gin.Context) {
	filters := make(map[string]interface{})
	for key, val := range ctx.Request.URL.Query() {
		filters[key] = val
	}

	Orm := repository.NewOrmStory()
	stories, err := Orm.GetByFilter(filters)
	status := http.StatusOK
	if err != nil {
		status = http.StatusBadRequest
	}
	ctx.JSON(status, gin.H{
		"context": stories,
		"error":   err,
	})
}
