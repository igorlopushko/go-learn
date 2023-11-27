package model

type Item struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type Order struct {
	Id    int     `json:"id"`
	Total float32 `json:"total"`
	Items []Item  `json:"items"`
}
