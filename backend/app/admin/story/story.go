package adminStory

import (
	StoryOrm "blog/model/story"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) {
	Orm := StoryOrm.NewOrmStory()
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
	Orm := StoryOrm.NewOrmStory()
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
	Orm := StoryOrm.NewOrmStory()
	story := &StoryOrm.Story{
		Title:   ctx.PostForm("title"),
		Author:  ctx.PostForm("author"),
		Context: ctx.PostForm("context"),
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
	Orm := StoryOrm.NewOrmStory()
	story := &StoryOrm.Story{
		Title:   ctx.PostForm("title"),
		Author:  ctx.PostForm("author"),
		Context: ctx.PostForm("context"),
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
	Orm := StoryOrm.NewOrmStory()
	err := Orm.Delete(ctx.Param("title"))
	status := http.StatusOK
	if err != nil {
		status = http.StatusBadRequest
	}
	ctx.JSON(status, gin.H{
		"error": err,
	})
}

func GetByKey(ctx *gin.Context) {
	keyValPair := map[string]interface{}{}
	queries := ctx.Request.URL.Query()
	for key, val := range queries {
		keyValPair[key] = val
	}
	Orm := StoryOrm.NewOrmStory()
	stories, err := Orm.GetByKey(keyValPair)
	status := http.StatusOK
	if err != nil {
		status = http.StatusBadRequest
	}
	ctx.JSON(status, gin.H{
		"context": stories,
		"error":   err,
	})
}