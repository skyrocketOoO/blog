package adminAuthor

import (
	AuthorOrm "blog/model/author"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) {
	Orm := AuthorOrm.NewOrmAuthor()
	authors, err := Orm.GetAll()
	status := http.StatusOK
	if err != nil {
		status = http.StatusBadRequest
	}
	ctx.JSON(status, gin.H{
		"context": authors,
		"error":   err,
	})
}

func GetByTitle(ctx *gin.Context) {
	Orm := AuthorOrm.NewOrmAuthor()
	author, err := Orm.GetByTitle(ctx.Param("name"))
	status := http.StatusOK
	if err != nil {
		status = http.StatusBadRequest
	}
	ctx.JSON(status, gin.H{
		"context": author,
		"error":   err,
	})
}

func Create(ctx *gin.Context) {
	Orm := AuthorOrm.NewOrmAuthor()
	author := &AuthorOrm.Author{
		Name:            ctx.PostForm("name"),
		Introduce:       ctx.PostForm("introduce"),
		PhotoAddress:    ctx.PostForm("photoaddress"),
		PersonalWebsite: ctx.PostForm("personalwebsite"),
	}
	err := Orm.Create(author)
	status := http.StatusOK
	if err != nil {
		status = http.StatusBadRequest
	}
	ctx.JSON(status, gin.H{
		"error": err,
	})
}

func Update(ctx *gin.Context) {
	Orm := AuthorOrm.NewOrmAuthor()
	author := &AuthorOrm.Author{
		Name:            ctx.PostForm("name"),
		Introduce:       ctx.PostForm("introduce"),
		PhotoAddress:    ctx.PostForm("photoaddress"),
		PersonalWebsite: ctx.PostForm("personalwebsite"),
	}
	err := Orm.Update(ctx.Param("name"), author)
	status := http.StatusOK
	if err != nil {
		status = http.StatusBadRequest
	}
	ctx.JSON(status, gin.H{
		"error": err,
	})
}

func Delete(ctx *gin.Context) {
	Orm := AuthorOrm.NewOrmAuthor()
	err := Orm.Delete(ctx.Param("name"))
	status := http.StatusOK
	if err != nil {
		status = http.StatusBadRequest
	}
	ctx.JSON(status, gin.H{
		"error": err,
	})
}
