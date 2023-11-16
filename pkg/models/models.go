package models

type AIResponce struct {
	UserID int
	ChatID int64
}

type Dish struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Ingredients string   `json:"ingredients"`
	Recipe      string   `json:"recipe"`
	Url         string   `json:"url"`
	Params      []string `json:"params"`
	Link        string   `json:"link"`
	RussianName string   `json:"russianName"`
}
