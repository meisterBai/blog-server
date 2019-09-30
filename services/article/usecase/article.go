package usecase

import (
	"server/services/article"
	"time"
)

type articleUsecase struct {
	articleRepo article.Repository
	contextTimeout time.Duration
}


func NewArticleUsecase(a article.Repository, timeout time.Duration) article.Usecase {
	return &articleUsecase{
		articleRepo:a,
		contextTimeout:timeout,
	}
}
