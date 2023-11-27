package repo

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/igorlopushko/go-learn/coffee-shop/model"
)

type FileRepo struct{}

func (r FileRepo) Load(lang string) ([]model.Item, error) {
	jsonFile, err := os.Open(fmt.Sprintf("data/data-%v.json", lang))

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var items []model.Item

	err = json.Unmarshal(byteValue, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r FileRepo) Save(order model.Order) error {
	file, err := os.Create(fmt.Sprintf("data/order-%v.json", order.Id))
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(order)
	if err != nil {
		return err
	}

	return nil
}
