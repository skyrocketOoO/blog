package story

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
