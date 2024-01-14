package routes

import (
	"moonlay-test/handlers"
	"moonlay-test/pkg/database"
	"moonlay-test/pkg/upload"
	"moonlay-test/repository"

	"github.com/labstack/echo/v4"
)

func SublistRoutes(e *echo.Group) {
	r := repository.RepositorySublist(database.DB)
	h := handlers.SublistHandler(r)

	e.GET("/lists/:id/sublist", h.FindSublistsByFilter)
	e.GET("/sublist/:id", h.GetSublist)
	e.POST("/sublist", upload.UploadFile(h.CreateSublist))
	e.PATCH("/sublist/:id", h.UpdateSublist)
	e.DELETE("/sublist/:id", h.DeleteSublist)
}
