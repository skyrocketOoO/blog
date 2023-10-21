package delivery

import "blog/internal/usecase"

type Handler struct {
	StoryHandler StoryHandler
}

func NewHandler(usecase *usecase.Usecase) *Handler {
	return &Handler{
		StoryHandler: *NewStoryHandler(usecase.StoryUsecase),
	}
}
