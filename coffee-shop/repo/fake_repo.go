package repo

import (
	"errors"

	"github.com/igorlopushko/go-learn/coffee-shop/model"
)

type FakeRepo struct{}

const (
	en = "en"
	ua = "ua"
)

func (r FakeRepo) Load(lang string) ([]model.Item, error) {
	switch lang {
	case en:
		return []model.Item{
			{Id: 1, Name: "Americano", Price: 30.00},
			{Id: 2, Name: "Cappuccino", Price: 35.00},
			{Id: 3, Name: "Espresso", Price: 25.00},
			{Id: 4, Name: "Latte", Price: 40.00},
		}, nil
	case ua:
		return []model.Item{
			{Id: 1, Name: "Американо", Price: 30.00},
			{Id: 2, Name: "Капучіно", Price: 35.00},
			{Id: 3, Name: "Еспресо", Price: 25.00},
			{Id: 4, Name: "Латте", Price: 40.00},
		}, nil

	}

	return nil, errors.New("Not supported language")
}

func (r FakeRepo) Save(order *model.Order) error {
	return nil
}
