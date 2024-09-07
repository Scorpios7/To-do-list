package models

type Todo struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}
