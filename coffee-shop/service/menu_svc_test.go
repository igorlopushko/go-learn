package service

import (
	"testing"

	"github.com/igorlopushko/go-learn/coffee-shop/repo"
)

func TestGetItems(t *testing.T) {
	mLen := 4
	lang := "en"
	r := repo.FakeRepo{}
	menuSvc := NewMenuService(r, lang)
	m, err := menuSvc.GetItems()

	if err != nil {
		t.Error(err)
	}

	if len(m) != mLen {
		t.Errorf("Got wrong menu len(). Has to be %v, but is %v", mLen, len(m))
	}
}
