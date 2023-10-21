package repository

import (
	"gorm.io/gorm"
)

type Story struct {
	gorm.Model
	Title   string
	Content string
	IsDraft bool
	Label   string
}

type StoryOrm struct {
	db *gorm.DB
}

func NewStoryOrm(db *gorm.DB) *StoryOrm {
	db.AutoMigrate(&Story{})
	return &StoryOrm{
		db: db,
	}
}

func (s *StoryOrm) GetAll() ([]Story, error) {
	stories := []Story{}
	err := s.db.Find(&stories).Error
	return stories, err
}

func (s *StoryOrm) Get(id uint) (Story, error) {
	var story Story
	if err := s.db.First(&story, id).Error; err != nil {
		return Story{}, err
	}
	return story, nil
}

func (s *StoryOrm) Create(title, content, label string, isDraft bool) (uint, error) {

	newStory := Story{
		Title:   title,
		Content: content,
		Label:   label,
		IsDraft: isDraft,
	}
	if err := s.db.Create(&newStory).Error; err != nil {
		return 0, err
	}

	return newStory.ID, nil
}

func (s *StoryOrm) Update(id uint, title, content, label string, isDraft bool) error {
	existingStory := Story{}
	if err := s.db.First(&existingStory, id).Error; err != nil {
		return err
	}

	if title != "" {
		existingStory.Title = title
	}
	if content != "" {
		existingStory.Content = content
	}
	if label != "" {
		existingStory.Label = label
	}
	existingStory.IsDraft = isDraft

	if err := s.db.Save(&existingStory).Error; err != nil {
		return err
	}

	return nil
}

func (s *StoryOrm) Delete(id uint) error {
	existingStory := Story{}
	if err := s.db.First(&existingStory, id).Error; err != nil {
		return err
	}

	if err := s.db.Delete(&existingStory).Error; err != nil {
		return err
	}

	return nil
}

func (s *StoryOrm) GetByFilter(filters map[string]interface{}) ([]Story, error) {
	stories := []Story{}
	err := s.db.Where(filters).Find(&stories).Error
	return stories, err
}
