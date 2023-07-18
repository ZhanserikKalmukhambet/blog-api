package model

type Category struct {
	ID   int64  `json:"ID" db:"id"`
	Name string `json:"Name"`
}
