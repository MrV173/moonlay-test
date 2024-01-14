package models

type Sublist struct {
	ID          int          `json:"id" gorm:"primaryKey:autoIncrement"`
	Title       string       `json:"title" gorm:"size:100"`
	Description string       `json:"description" gorm:"size:1000"`
	ListID      int          `json:"list_id"`
	List        ListResponse `json:"-"`
	File        string       `json:"file" gorm:"size:300"`
}

type SublistResponse struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	ListID      int          `json:"list_id"`
	List        ListResponse `json:"-"`
	File        string       `json:"file"`
}

func (SublistResponse) TableName() string {
	return "sublists"
}
