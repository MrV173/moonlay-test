package todolistdto

type CreateToDo struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	File        string `json:"file" form:"file"`
}

type UpdateToDO struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	File        string `json:"file" form:"file"`
}

type ToDoResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	File        string `json:"file"`
}
