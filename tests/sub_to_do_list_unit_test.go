package tests

import (
	"fmt"
	subtodolistdto "moonlay-test/dto/sub_to_do_list"
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

func TestFindSublistsByFilter(t *testing.T) {
	e := echo.New()

	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	listRepository := repository.RepositoryList(db)
	sublistRepository := repository.RepositorySublist(db)

	handler := handlers.SublistHandler(sublistRepository)

	createdList, err := listRepository.CreateList(models.List{
		Title:       "Test Title",
		Description: "Test description",
	})
	if err != nil {
		t.Fatalf("Failed to create initial list entry: %v", err)
	}

	sublist, err := sublistRepository.CreateSublist(models.Sublist{
		ListID:      createdList.ID,
		Title:       "sublist title",
		Description: "sublist description",
	})
	if err != nil {
		t.Fatalf("Failed to create initial sublist entry: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/lists/"+strconv.Itoa(createdList.ID)+"/sublist/?page=1&limit=2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/lists/:id/sublist/")

	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(createdList.ID))
	err = handler.FindSublistsByFilter(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, sublist.Title, "sublist title")
	assert.Equal(t, sublist.Description, "sublist description")

}

func TestGetSublist(t *testing.T) {
	e := echo.New()

	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	listRepository := repository.RepositoryList(db)

	sublistRepository := repository.RepositorySublist(db)

	handler := handlers.SublistHandler(sublistRepository)

	createdList, err := listRepository.CreateList(models.List{
		Title:       "GetSublistTest",
		Description: "Test description",
		File:        "get_sublist_file.txt",
	})
	if err != nil {
		t.Fatalf("Failed to create initial list entry: %v", err)
	}

	createdSublist, err := sublistRepository.CreateSublist(models.Sublist{
		ListID:      createdList.ID,
		Title:       "sublist title",
		Description: "sublist description",
	})
	if err != nil {
		t.Fatalf("Failed to create initial sublist entry: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/sublist/"+strconv.Itoa(createdSublist.ID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/sublist/:id")

	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(createdSublist.ID))

	err = handler.GetSublist(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, createdSublist.Title, "sublist title")
	assert.Equal(t, createdSublist.Description, "sublist description")

}

func TestCreateSublist(t *testing.T) {
	e := echo.New()

	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	listRepository := repository.RepositoryList(db)

	sublistRepository := repository.RepositorySublist(db)

	handler := handlers.SublistHandler(sublistRepository)

	createdList, err := listRepository.CreateList(models.List{
		Title:       "CreateSublistTest",
		Description: "Test description",
	})
	if err != nil {
		t.Fatalf("Failed to create initial list entry: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/sublist", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/sublist")

	c.Set("dataFile", []string{"file1.txt", "file2.txt"})

	req.Form = make(url.Values)
	req.Form.Add("title", "TestSublistTitle")
	req.Form.Add("description", "TestSublistDescription")
	req.Form.Add("list_id", strconv.Itoa(createdList.ID))

	err = handler.CreateSublist(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

}

func TestUpdateSublist(t *testing.T) {
	e := echo.New()

	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	listRepository := repository.RepositoryList(db)

	createdList, err := listRepository.CreateList(models.List{
		Title:       "CreateSublistTest",
		Description: "Test description",
	})
	if err != nil {
		t.Fatalf("Failed to create initial list entry: %v", err)
	}

	sublistRepository := repository.RepositorySublist(db)

	handler := handlers.SublistHandler(sublistRepository)

	createdSublist, err := sublistRepository.CreateSublist(models.Sublist{
		Title:       "UpdateSublistTest",
		Description: "Test description",
		ListID:      createdList.ID,
	})
	if err != nil {
		t.Fatalf("Failed to create initial sublist entry: %v", err)
	}

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/sublist/%d", createdSublist.ID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/sublist/:id")

	updatePayload := subtodolistdto.UpdateSublist{
		Title:       "Updated Title",
		Description: "Updated Description",
	}

	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(createdSublist.ID))
	c.Set("dataFile", []string{"file1.txt", "file2.txt"})
	c.SetRequest(req)
	req.Form = make(url.Values)
	req.Form.Add("title", updatePayload.Title)
	req.Form.Add("description", updatePayload.Description)

	err = handler.UpdateSublist(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, updatePayload.Title, "Updated Title")
	assert.Equal(t, updatePayload.Description, "Updated Description")

}

func TestDeleteSublist(t *testing.T) {
	e := echo.New()

	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	listRepository := repository.RepositoryList(db)

	createdList, err := listRepository.CreateList(models.List{
		Title:       "CreateSublistTest",
		Description: "Test description",
	})
	if err != nil {
		t.Fatalf("Failed to create initial list entry: %v", err)
	}

	sublistRepository := repository.RepositorySublist(db)

	handler := handlers.SublistHandler(sublistRepository)

	createdSublist, err := sublistRepository.CreateSublist(models.Sublist{
		Title:       "Deleted Sub Title",
		Description: "Deleted Sub Description",
		ListID:      createdList.ID,
	})
	if err != nil {
		t.Fatalf("Failed to create initial sublist entry: %v", err)
	}

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/sublist/%d", createdSublist.ID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/sublist/:id")

	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(createdSublist.ID))
	err = handler.DeleteSublist(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, createdSublist.Title, "Deleted Sub Title")
	assert.Equal(t, createdSublist.Description, "Deleted Sub Description")

}
