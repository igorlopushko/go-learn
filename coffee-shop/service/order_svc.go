package service

import (
	"fmt"
	"sync"

	"github.com/igorlopushko/go-learn/coffee-shop/model"
	"github.com/igorlopushko/go-learn/coffee-shop/repo"
)

type OrderService struct {
	repo        repo.Repository
	orderNumber int
	mu          sync.Mutex
}

func NewOrderService(repo repo.Repository) OrderService {
	return OrderService{
		repo:        repo,
		orderNumber: 1,
	}
}

func (s *OrderService) Create(items []model.Item) error {
	var total float32
	for _, item := range items {
		total += item.Price
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	order := model.Order{
		Id:    s.orderNumber,
		Items: items,
		Total: total,
	}

	err := s.repo.Save(order)
	if err != nil {
		return err
	}

	s.orderNumber++

	return nil
}

func (s *OrderService) CreateMany(buckets [][]model.Item) error {
	var wg sync.WaitGroup

	for _, items := range buckets {
		wg.Add(1)

		go func() {
			defer wg.Done()
			err := s.Create(items)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	wg.Wait()

	return nil
}
