package handlers

import (
	"moonlay-test/dto/result"
	subtodolistdto "moonlay-test/dto/sub_to_do_list"
	"moonlay-test/models"
	"moonlay-test/repository"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type HandlerSublist struct {
	SublistRepository repository.SublistRepository
}

func SublistHandler(sublistRepository repository.SublistRepository) *HandlerSublist {
	return &HandlerSublist{sublistRepository}
}

func (h *HandlerSublist) FindSublistsByFilter(c echo.Context) error {
	listID, _ := strconv.Atoi(c.Param("id"))
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	title := c.QueryParam("title")
	description := c.QueryParam("description")

	if err != nil || limit < 1 {
		limit = 2
	}

	response, err := h.SublistRepository.FindSublistsByFilter(listID, page, limit, title, description)
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

func (h *HandlerSublist) GetSublist(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	sublist, err := h.SublistRepository.GetSubList(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, result.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	sublist.File = path_file + sublist.File

	return c.JSON(http.StatusOK, result.SuccessResult{
		Code: http.StatusOK,
		Data: sublist,
	})
}

func (h *HandlerSublist) CreateSublist(c echo.Context) error {
	var err error
	files := c.Get("dataFile").([]string)
	list_id, _ := strconv.Atoi(c.FormValue("list_id"))
	mainFile := "-"
	if len(files) > 0 {
		mainFile = files[0]
	}

	request := subtodolistdto.CreateSublist{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		File:        mainFile,
		ListID:      list_id,
	}

	validation := validator.New()
	err = validation.Struct(request)

	sublist := models.Sublist{
		Title:       request.Title,
		Description: request.Description,
		File:        request.File,
		ListID:      request.ListID,
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, result.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	response, err := h.SublistRepository.CreateSublist(sublist)
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

func (h *HandlerSublist) UpdateSublist(c echo.Context) error {
	request := new(subtodolistdto.UpdateSublist)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	sublist, err := h.SublistRepository.GetSubList(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if request.Title != "" {
		sublist.Title = request.Title
	}

	if request.Description != "" {
		sublist.Description = request.Description
	}

	if request.File != "" {
		sublist.File = request.File
	}

	ListID := strconv.Itoa(request.ListID)
	if ListID != "" {
		newListID, err := strconv.Atoi(ListID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, result.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		}
		sublist.ListID = newListID
	}

	response, err := h.SublistRepository.UpdateSublist(sublist)
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

func (h *HandlerSublist) DeleteSublist(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	list, err := h.SublistRepository.GetSubList(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	response, err := h.SublistRepository.DeleteSublist(list)
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
