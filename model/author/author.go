package author

import (
	Db "blog/db"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Author struct {
	Name            string `gorm:"primaryKey"`
	Introduce       string
	PhotoAddress    string
	PersonalWebsite string
}

func (s *Author) GetAll(ctx *gin.Context) ([]Author, error) {
	authors := []Author{}
	tx := Db.Db.Find(&authors)
	return authors, tx.Error
}

func (s *Author) GetByname(ctx *gin.Context) (Author, error) {
	author := Author{Name: ctx.Param("name")}
	tx := Db.Db.First(&author)
	return author, tx.Error
}

func (s *Author) Create(ctx *gin.Context) error {
	if inputCheck(ctx.PostForm("name"), ctx.PostForm("introduce"), ctx.PostForm("photoAddress")) {
		return errors.New("check the input info")
	}
	tx := Db.Db.Model(&Author{}).Create(map[string]interface{}{
		"Name":            ctx.PostForm("name"),
		"Introduce":       ctx.PostForm("introduce"),
		"PhotoAddress":    ctx.PostForm("photoAddress"),
		"PersonalWebsite": ctx.PostForm("personalWebsite"),
	})
	return tx.Error
}

func (s *Author) Update(ctx *gin.Context) error {
	if inputCheck(ctx.PostForm("name"), ctx.PostForm("author"), ctx.PostForm("context")) {
		return errors.New("check the input info")
	}

	model := Db.Db.First(&Author{Name: ctx.Param("name")})
	if model.Error == gorm.ErrRecordNotFound {
		return model.Error
	}
	tx := model.Updates(map[string]interface{}{
		"Name":            ctx.PostForm("name"),
		"Introduce":       ctx.PostForm("introduce"),
		"PhotoAddress":    ctx.PostForm("photoAddress"),
		"PersonalWebsite": ctx.PostForm("personalWebsite"),
	})
	return tx.Error
}

func (s *Author) Delete(ctx *gin.Context) error {
	author := Author{Name: ctx.Param("name")}
	tx := Db.Db.First(&author)
	if tx.Error == gorm.ErrRecordNotFound {
		return tx.Error
	}
	tx = Db.Db.Delete(&author)
	return tx.Error
}

func inputCheck(name, author, context string) bool {
	if name == "" || author == "" || context == "" {
		return true
	}
	return false
}
