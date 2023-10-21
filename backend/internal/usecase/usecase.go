package usecase

import (
	"blog/internal/infra/repository"
)

type Usecase struct {
	StoryUsecase *StoryUsecase
}

func NewUsecase(repo *repository.Repository) *Usecase {
	return &Usecase{
		StoryUsecase: NewStoryUsecase(repo.StoryOrm),
	}
}
