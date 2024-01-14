package tests

import (
	todolistdto "moonlay-test/dto/to_do_list"
	"moonlay-test/handlers"
	"moonlay-test/models"
	"moonlay-test/repository"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestFindListsByFilter(t *testing.T) {
	e := echo.New()

	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	listRepository := repository.RepositoryList(db)

	handler := handlers.ListHandler(listRepository)

	dummyList := models.List{
		Title:       "title",
		Description: "description",
	}
	_, err = listRepository.CreateList(dummyList)
	if err != nil {
		t.Fatalf("Failed to create dummy list: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/lists", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/lists")
	c.SetParamValues("title", "title")
	c.SetParamValues("description", "")
	c.SetParamValues("page", "1")
	c.SetParamValues("limit", "3")

	err = handler.FindListsByFilter(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, dummyList.Title, "title")
	assert.Equal(t, dummyList.Description, "description")
	assert.Equal(t, dummyList.File, "")

}

func TestGetList(t *testing.T) {
	e := echo.New()

	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	listRepository := repository.RepositoryList(db)

	handler := handlers.ListHandler(listRepository)

	dummyList := models.List{
		Title:       "title",
		Description: "description",
	}
	createdList, err := listRepository.CreateList(dummyList)
	if err != nil {
		t.Fatalf("Failed to create dummy list: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/list/"+strconv.Itoa(createdList.ID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/list/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(createdList.ID))

	err = handler.GetList(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, dummyList.Title, "title")
	assert.Equal(t, dummyList.Description, "description")
	assert.Equal(t, dummyList.File, "")

}

func TestCreateList(t *testing.T) {
	e := echo.New()

	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	listRepository := repository.RepositoryList(db)

	handler := handlers.ListHandler(listRepository)

	req := httptest.NewRequest(http.MethodPost, "/list", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/list")

	c.Set("dataFile", []string{"file1.txt", "file2.txt"})

	req.Form = make(url.Values)
	req.Form.Add("title", "YourTitle")
	req.Form.Add("description", "YourDescription")

	err = handler.CreateList(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

}

func TestUpdateList(t *testing.T) {
	e := echo.New()

	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	listRepository := repository.RepositoryList(db)

	handler := handlers.ListHandler(listRepository)

	createdList, err := listRepository.CreateList(models.List{
		Title:       "Title",
		Description: "Description",
	})
	if err != nil {
		t.Fatalf("Failed to create initial list entry: %v", err)
	}

	req := httptest.NewRequest(http.MethodPut, "/list/"+strconv.Itoa(createdList.ID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/list/:id")

	updatePayload := todolistdto.UpdateToDO{
		Title:       "UpdatedTitle",
		Description: "UpdatedDescription",
	}

	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(createdList.ID))
	c.Set("dataFile", []string{"file1.txt", "file2.txt"})
	c.SetRequest(req)
	req.Form = make(url.Values)
	req.Form.Add("title", updatePayload.Title)
	req.Form.Add("description", updatePayload.Description)

	err = handler.UpdateList(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, updatePayload.Title, "UpdatedTitle")
	assert.Equal(t, updatePayload.Description, "UpdatedDescription")
	assert.Equal(t, updatePayload.File, "")

}

func TestDeleteList(t *testing.T) {
	e := echo.New()

	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	listRepository := repository.RepositoryList(db)

	handler := handlers.ListHandler(listRepository)

	createdList, err := listRepository.CreateList(models.List{
		Title:       "title",
		Description: "description",
	})
	if err != nil {
		t.Fatalf("Failed to create initial list entry: %v", err)
	}

	req := httptest.NewRequest(http.MethodDelete, "/list/"+strconv.Itoa(createdList.ID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/list/:id")

	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(createdList.ID))

	err = handler.DeleteList(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, createdList.Title, "title")
	assert.Equal(t, createdList.Description, "description")
	assert.Equal(t, createdList.File, "")

}
