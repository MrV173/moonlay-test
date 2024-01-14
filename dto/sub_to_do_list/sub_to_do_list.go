package subtodolistdto

type CreateSublist struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	File        string `json:"file" form:"file"`
	ListID      int    `json:"list_id" form:"list_id"`
}

type UpdateSublist struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	File        string `json:"file" form:"file"`
	ListID      int    `json:"list_id" form:"list_id"`
}

type SublistResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	File        string `json:"file"`
	ListID      int    `json:"-"`
}
