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
}

func NewOrderService(repo repo.Repository) OrderService {
	return OrderService{
		repo:        repo,
		orderNumber: 1,
	}
}

func (s *OrderService) Create(items []model.Item) (*model.Order, error) {
	order, err := s.createOrder(items)
	if err != nil {
		return nil, err
	}

	s.saveOrder(order)

	return order, nil
}

func (s *OrderService) CreateMany(buckets [][]model.Item) error {
	ordersChannel := make(chan *model.Order)

	go func() {
		for _, v := range buckets {
			order, err := s.createOrder(v)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}

			ordersChannel <- order
		}

		close(ordersChannel)
	}()

	var wg sync.WaitGroup

	// for {
	// 	msg, ok := <-ordersChannel
	// 	if !ok {
	// 		break
	// 	}

	// 	wg.Add(1)
	// 	go func(order *model.Order) {
	// 		s.saveOrder(order)
	// 		wg.Done()
	// 	}(msg)
	// }

	for msg := range ordersChannel {
		wg.Add(1)
		go func(order *model.Order) {
			s.saveOrder(order)
			wg.Done()
		}(msg)
	}

	wg.Wait()

	return nil
}

func (s *OrderService) createOrder(items []model.Item) (*model.Order, error) {
	var total float32
	for _, item := range items {
		total += item.Price
	}

	order := model.Order{
		Id:    s.orderNumber,
		Items: items,
		Total: total,
	}

	s.orderNumber++

	return &order, nil
}

func (s *OrderService) saveOrder(order *model.Order) {
	err := s.repo.Save(order)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
