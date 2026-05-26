package web

import (
	"context"

	"bilbo-baggins/internal/entity"
)

type HomePageRepository struct{}

func NewHomePageRepository() HomePageRepository {
	return HomePageRepository{}
}

func (repository HomePageRepository) Find(ctx context.Context) (entity.Page, error) {
	content, err := staticFiles.ReadFile("static/index.html")
	if err != nil {
		return entity.Page{}, err
	}

	return entity.Page{
		Content:     content,
		ContentType: "text/html; charset=utf-8",
	}, nil
}
