package models

type List struct {
	ID          int               `json:"id" gorm:"primaryKey:autoIncrement"`
	Title       string            `json:"title" gorm:"size:100"`
	Description string            `json:"description" gorm:"size:1000"`
	Sublists    []SublistResponse `json:"sublists"`
	File        string            `json:"file" gorm:"size:300"`
}

type ListResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	File        string `json:"file"`
}

func (ListResponse) TableName() string {
	return "lists"
}
