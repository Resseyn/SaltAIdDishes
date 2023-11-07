package models

type AIResponce struct {
	UserID int
	ChatID int64
}

type Dish struct {
	ID          int
	Name        string
	Description string
	Ingredients string
	Recipe      string
	Url         string
	Params      []string
}
