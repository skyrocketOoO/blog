package authorOrm

import (
	Db "blog/db"
	"errors"
)

type Author struct {
	Name            string `gorm:"primaryKey"`
	Introduce       string
	PhotoAddress    string
	PersonalWebsite string
}

func NewOrmAuthor() *Author {
	return &Author{}
}

func (s *Author) GetAll() ([]Author, error) {
	stories := []Author{}
	err := Db.Db.Find(&stories).Error
	return stories, err
}

func (s *Author) GetByTitle(title string) (Author, error) {
	author := Author{}
	err := Db.Db.Where("Title = ?", title).First(&author).Error
	return author, err
}

func (s *Author) Create(author *Author) error {
	if inputCheck(author.Name, author.Introduce, author.PhotoAddress) {
		return errors.New("check the input info")
	}
	err := Db.Db.Model(&Author{}).Create(map[string]interface{}{
		"Title":   author.Name,
		"Author":  author.Introduce,
		"Context": author.PhotoAddress,
		"Label":   author.PersonalWebsite,
	}).Error
	return err
}

func (s *Author) Update(title string, author *Author) error {
	if inputCheck(author.Name, author.Introduce, author.PhotoAddress) {
		return errors.New("check the input info")
	}
	err := Db.Db.Model(&Author{}).Where("Title = ?", title).Updates(map[string]interface{}{
		"Title":   author.Name,
		"Author":  author.Introduce,
		"Context": author.PhotoAddress,
		"Label":   author.PersonalWebsite,
	}).Error
	return err
}

func (s *Author) Delete(title string) error {
	err := Db.Db.Where("Title = ?", title).Delete(&Author{}).Error
	return err
}

func inputCheck(name, author, context string) bool {
	if name == "" || author == "" || context == "" {
		return true
	}
	return false
}
