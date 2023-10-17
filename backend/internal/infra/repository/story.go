package repository

import (
	Db "blog/db"
	"errors"

	"gorm.io/gorm"
)

type Story struct {
	gorm.Model
	Author  string
	Title   string
	Content string
	IsDraft bool
	Label   string
}

func NewOrmStory() *Story {
	return &Story{}
}

func (s *Story) GetAll() ([]Story, error) {
	stories := []Story{}
	err := Db.Db.Find(&stories).Error
	return stories, err
}

func (s *Story) GetByTitle(title string) (Story, error) {
	story := Story{}
	err := Db.Db.Where("Title = ?", title).First(&story).Error
	return story, err
}

func (s *Story) Create(story *Story) error {
	if inputCheck(story.Title, story.Author, story.Content) {
		return errors.New("Title, Author, or Content cannot be empty")
	}
	err := Db.Db.Model(&Story{}).Create(map[string]interface{}{
		"Title":   story.Title,
		"Author":  story.Author,
		"Context": story.Content,
		"Label":   story.Label,
		"IsDraft": story.IsDraft,
	}).Error
	return err
}

func (s *Story) Update(title string, story *Story) error {
	if inputCheck(story.Title, story.Author, story.Content) {
		return errors.New("Title, Author, or Content cannot be empty")
	}
	err := Db.Db.Model(&Story{}).Where("Title = ?", title).Updates(map[string]interface{}{
		"Title":   story.Title,
		"Author":  story.Author,
		"Context": story.Content,
		"Label":   story.Label,
		"IsDraft": story.IsDraft,
	}).Error
	return err
}

func (s *Story) Delete(title string) error {
	err := Db.Db.Where("Title = ?", title).Delete(&Story{}).Error
	return err
}

func (s *Story) GetByFilter(filters map[string]interface{}) ([]Story, error) {
	stories := []Story{}
	err := Db.Db.Where(filters).Find(&stories).Error
	return stories, err
}

func inputCheck(title, author, context string) bool {
	if title == "" || author == "" || context == "" {
		return true
	}
	return false
}
