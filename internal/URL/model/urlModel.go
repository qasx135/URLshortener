package model

type URLModel struct {
	ID    int64  `json:"id"`
	Alias string `json:"alias"`
	Url   string `json:"url"`
}
