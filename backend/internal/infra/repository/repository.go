package repository

import "gorm.io/gorm"

type Repository struct {
	StoryOrm *StoryOrm
}

func NewRepository(db *gorm.DB) *Repository {
	storyOrm := NewStoryOrm(db)

	return &Repository{
		StoryOrm: storyOrm,
	}
}
