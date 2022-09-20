package storyMd

import (
	Db "blog/db"
	"errors"

	"gorm.io/gorm"
)

type Story struct {
	gorm.Model
	Author  string
	Title   string
	Context string
	Label   string
}

func NewOrmStory() *Story {
	return &Story{}
}

func (s *Story) GetAll() ([]Story, error) {
	keywords := []Story{}
	tx := Db.Db.Find(&keywords)
	return keywords, tx.Error
}

func (s *Story) GetByTitle(title string) (Story, error) {
	keyword := Story{Title: title}
	tx := Db.Db.First(&keyword)
	return keyword, tx.Error
}

func (s *Story) Create(story *Story) error {
	if inputCheck(story.Title, story.Author, story.Context) {
		return errors.New("check the input info")
	}
	tx := Db.Db.Model(&Story{}).Create(map[string]interface{}{
		"Title":   story.Title,
		"Author":  story.Author,
		"Context": story.Context,
		"Label":   story.Label,
	})
	return tx.Error
}

func (s *Story) Update(title string, story *Story) error {
	if inputCheck(story.Title, story.Author, story.Context) {
		return errors.New("check the input info")
	}
	tx := Db.Db.Model(&Story{}).Where("Title = ?", title).Updates(map[string]interface{}{
		"Title":   story.Title,
		"Author":  story.Author,
		"Context": story.Context,
		"Label":   story.Label,
	})
	return tx.Error
}

func (s *Story) Delete(title string) error {
	tx := Db.Db.Where("Title = ?", title).Delete(&Story{})
	return tx.Error
}

func inputCheck(title, author, context string) bool {
	if title == "" || author == "" || context == "" {
		return true
	}
	return false
}
