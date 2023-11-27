package repo

import "github.com/igorlopushko/go-learn/coffee-shop/model"

type Repository interface {
	Load(lang string) ([]model.Item, error)
	Save(order model.Order) error
}
