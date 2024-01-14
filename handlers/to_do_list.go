package handlers

import (
	"moonlay-test/dto/result"
	todolistdto "moonlay-test/dto/to_do_list"
	"moonlay-test/models"
	"moonlay-test/repository"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var path_file = "http://localhost:5000/storages/"

type HandlerList struct {
	ListRepository repository.ListRepository
}

func ListHandler(listRepository repository.ListRepository) *HandlerList {
	return &HandlerList{listRepository}
}

func (h *HandlerList) FindListsByFilter(c echo.Context) error {
	title := c.QueryParam("title")
	description := c.QueryParam("description")
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 3
	}

	paginationResult, err := h.ListRepository.FindListsByFilter(title, description, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result.SuccessResult{
		Code: http.StatusOK,
		Data: paginationResult,
	})
}

func (h *HandlerList) GetList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	list, err := h.ListRepository.GetList(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, result.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	list.File = path_file + list.File

	return c.JSON(http.StatusOK, result.SuccessResult{
		Code: http.StatusOK,
		Data: list,
	})
}

func (h *HandlerList) CreateList(c echo.Context) error {
	var err error
	files := c.Get("dataFile").([]string)
	mainFile := "-"
	if len(files) > 0 {
		mainFile = files[0]
	}
	request := todolistdto.CreateToDo{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		File:        mainFile,
	}

	validation := validator.New()
	err = validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, result.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	list := models.List{
		Title:       request.Title,
		Description: request.Description,
		File:        request.File,
	}

	response, err := h.ListRepository.CreateList(list)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, result.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result.SuccessResult{
		Code: http.StatusOK,
		Data: response,
	})
}

func (h *HandlerList) UpdateList(c echo.Context) error {
	request := new(todolistdto.UpdateToDO)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	list, err := h.ListRepository.GetList(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if request.Title != "" {
		list.Title = request.Title
	}

	if request.Description != "" {
		list.Description = request.Description
	}

	if request.File != "" {
		list.File = request.File
	}

	response, err := h.ListRepository.UpdateList(list)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result.SuccessResult{
		Code: http.StatusOK,
		Data: response,
	})
}

func (h *HandlerList) DeleteList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	list, err := h.ListRepository.GetList(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	response, err := h.ListRepository.DeleteList(list)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result.SuccessResult{
		Code: http.StatusOK,
		Data: response,
	})

}
