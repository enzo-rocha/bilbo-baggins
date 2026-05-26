package usecase

import (
	"context"

	"bilbo-baggins/internal/entity"
)

type HomePageRepository interface {
	Find(ctx context.Context) (entity.Page, error)
}

type ShowHomePage struct {
	homePageRepository HomePageRepository
}

func NewShowHomePage(homePageRepository HomePageRepository) ShowHomePage {
	return ShowHomePage{
		homePageRepository: homePageRepository,
	}
}

func (uc ShowHomePage) Execute(ctx context.Context) (entity.Page, error) {
	return uc.homePageRepository.Find(ctx)
}
