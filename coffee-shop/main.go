package main

import (
	"flag"
	"fmt"

	"github.com/igorlopushko/go-learn/coffee-shop/model"
	"github.com/igorlopushko/go-learn/coffee-shop/repo"
	"github.com/igorlopushko/go-learn/coffee-shop/service"
)

var (
	lang string
)

func main() {
	flag.StringVar(&lang, "l", lang, "menu language (e.g. 'en' or 'ua')")
	flag.Parse()

	r := repo.FileRepo{}
	menuSvc := service.NewMenuService(r, lang)
	m, err := menuSvc.GetItems()
	if err != nil {
		fmt.Println(err)
	}

	menuSvc.Print(m)

	orderSvc := service.NewOrderService(r)
	/*
		err = orderSvc.Create(m[:2])
		if err != nil {
			fmt.Println(err)
		}

		err = orderSvc.Create(m[1:2])
		if err != nil {
			fmt.Println(err)
		}

		err = orderSvc.Create(m[2:])
		if err != nil {
			fmt.Println(err)
		}
	*/

	buckets := [][]model.Item{
		m[:2],
		m[1:2],
		m[2:],
	}
	err = orderSvc.CreateMany(buckets)
	if err != nil {
		fmt.Println(err)
	}
}
