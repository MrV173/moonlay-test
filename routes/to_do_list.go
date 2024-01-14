package routes

import (
	"moonlay-test/handlers"
	"moonlay-test/pkg/database"
	"moonlay-test/pkg/upload"
	"moonlay-test/repository"

	"github.com/labstack/echo/v4"
)

func ListRoutes(e *echo.Group) {
	r := repository.RepositoryList(database.DB)
	h := handlers.ListHandler(r)

	e.GET("/lists", h.FindListsByFilter)
	e.GET("/list/:id", h.GetList)
	e.POST("/list", upload.UploadFile(h.CreateList))
	e.PATCH("/list/:id", h.UpdateList)
	e.DELETE("/list/:id", h.DeleteList)
}
