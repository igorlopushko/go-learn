package service

import (
	"fmt"

	"github.com/igorlopushko/go-learn/coffee-shop/model"
	"github.com/igorlopushko/go-learn/coffee-shop/repo"
)

const (
	en = "en"
	ua = "ua"
)

type MenuService struct {
	repo     repo.Repository
	language string
}

func NewMenuService(repo repo.Repository, lang string) MenuService {
	return MenuService{
		repo:     repo,
		language: lang,
	}
}

func (s *MenuService) GetItems() ([]model.Item, error) {
	items, err := s.repo.Load(s.language)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *MenuService) Print(items []model.Item) {
	for _, v := range items {
		if s.language == en {
			fmt.Printf("%v.%v - %.2fUAH\n", v.Id, v.Name, v.Price)
		} else if s.language == ua {
			fmt.Printf("%v.%v - %.2fгрн.\n", v.Id, v.Name, v.Price)
		}
	}
}
