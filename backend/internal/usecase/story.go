package usecase

import (
	"blog/internal/infra/repository"
	"context"
)

type StoryUsecase struct {
	storyOrm *repository.StoryOrm
}

func NewStoryUsecase(storyOrm *repository.StoryOrm) *StoryUsecase {
	return &StoryUsecase{
		storyOrm: storyOrm,
	}
}

func (s *StoryUsecase) GetAll(ctx *context.Context) ([]repository.Story, error) {
	return s.storyOrm.GetAll()
}

func (s *StoryUsecase) Get(ctx *context.Context, id uint) (repository.Story, error) {
	return s.storyOrm.Get(id)
}

func (s *StoryUsecase) Create(ctx *context.Context, title, content, label string, isDraft bool) (uint, error) {
	return s.storyOrm.Create(title, content, label, isDraft)
}

func (s *StoryUsecase) Update(ctx *context.Context, id uint, title string, content string, label string, isDraft bool) error {
	return s.storyOrm.Update(id, title, content, label, isDraft)
}

func (s *StoryUsecase) Delete(ctx *context.Context, id uint) error {
	return s.storyOrm.Delete(id)
}

// func (s *StoryUsecase) GetByFilter(ctx *context.Context) {
// 	filters := make(map[string]interface{})
// 	for key, val := range ctx.Request.URL.Query() {
// 		filters[key] = val
// 	}

// 	stories, err := s.storyOrm.GetByFilter(filters)
// 	status := http.StatusOK
// 	if err != nil {
// 		status = http.StatusBadRequest
// 	}
// 	ctx.JSON(status, gin.H{
// 		"context": stories,
// 		"error":   err,
// 	})
// }
