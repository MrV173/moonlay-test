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
	// Setup Echo
	e := echo.New()

	// Setup PostgreSQL database connection
	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Create repository with real database connection
	listRepository := repository.RepositoryList(db)

	// Inject repository into handler
	handler := handlers.ListHandler(listRepository)

	// Insert dummy data for testing
	// (Assuming you have a function in your repository to create a List)
	dummyList := models.List{
		Title:       "title",
		Description: "deskripsi",
	}
	_, err = listRepository.CreateList(dummyList)
	if err != nil {
		t.Fatalf("Failed to create dummy list: %v", err)
	}

	// Create a request to test the handler
	req := httptest.NewRequest(http.MethodGet, "/lists", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/lists")
	c.SetParamValues("title", "title")  // Set your query parameters
	c.SetParamValues("description", "") // Set your query parameters
	c.SetParamValues("page", "1")       // Set your query parameters
	c.SetParamValues("limit", "3")      // Set your query parameters

	// Execute the handler
	err = handler.FindListsByFilter(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// TODO: Add more assertions based on the expected behavior
}

func TestGetList(t *testing.T) {
	// Setup Echo
	e := echo.New()

	// Setup PostgreSQL database connection
	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Create repository with real database connection
	listRepository := repository.RepositoryList(db)

	// Inject repository into handler
	handler := handlers.ListHandler(listRepository)

	// Insert dummy data for testing
	// (Assuming you have a function in your repository to create a List)
	dummyList := models.List{
		Title:       "title",
		Description: "deskripsi",
	}
	createdList, err := listRepository.CreateList(dummyList)
	if err != nil {
		t.Fatalf("Failed to create dummy list: %v", err)
	}

	// Create a request to test the handler
	req := httptest.NewRequest(http.MethodGet, "/list/"+strconv.Itoa(createdList.ID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/list/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(createdList.ID))

	// Execute the handler
	err = handler.GetList(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// TODO: Add more assertions based on the expected behavior
}

func TestCreateList(t *testing.T) {
	// Setup Echo
	e := echo.New()

	// Setup PostgreSQL database connection
	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Create repository with real database connection
	listRepository := repository.RepositoryList(db)

	// Inject repository into handler
	handler := handlers.ListHandler(listRepository)

	// Create a request to test the handler
	req := httptest.NewRequest(http.MethodPost, "/list", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/list")

	// Mock the upload middleware behavior
	c.Set("dataFile", []string{"file1.txt", "file2.txt"})

	// Set the required form values
	req.Form = make(url.Values)
	req.Form.Add("title", "YourTitle")             // Set your form values
	req.Form.Add("description", "YourDescription") // Set your form values

	// Execute the handler
	err = handler.CreateList(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// TODO: Add more assertions based on the expected behavior
}

func TestUpdateList(t *testing.T) {
	// Setup Echo
	e := echo.New()

	// Setup PostgreSQL database connection
	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Create repository with real database connection
	listRepository := repository.RepositoryList(db)

	// Inject repository into handler
	handler := handlers.ListHandler(listRepository)

	// Create a sample list entry for testing
	createdList, err := listRepository.CreateList(models.List{
		Title:       "InitialTitle",
		Description: "InitialDescription",
	})
	if err != nil {
		t.Fatalf("Failed to create initial list entry: %v", err)
	}

	// Create a request to test the handler
	req := httptest.NewRequest(http.MethodPut, "/list/"+strconv.Itoa(createdList.ID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/list/:id")

	// Mock the update payload
	updatePayload := todolistdto.UpdateToDO{
		Title:       "UpdatedTitle",
		Description: "UpdatedDescription",
	}

	// Bind the update payload to the request context
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(createdList.ID))
	c.Set("dataFile", []string{"file1.txt", "file2.txt"})
	c.SetRequest(req)
	req.Form = make(url.Values)
	req.Form.Add("title", updatePayload.Title)
	req.Form.Add("description", updatePayload.Description)

	// Execute the handler
	err = handler.UpdateList(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// TODO: Add more assertions based on the expected behavior
}

func TestDeleteList(t *testing.T) {
	// Setup Echo
	e := echo.New()

	// Setup PostgreSQL database connection
	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Create repository with real database connection
	listRepository := repository.RepositoryList(db)

	// Inject repository into handler
	handler := handlers.ListHandler(listRepository)

	// Create a sample list entry for testing
	createdList, err := listRepository.CreateList(models.List{
		Title:       "ToDeleteTitle",
		Description: "ToDeleteDescription",
		File:        "to_delete_file.txt",
	})
	if err != nil {
		t.Fatalf("Failed to create initial list entry: %v", err)
	}

	// Create a request to test the handler
	req := httptest.NewRequest(http.MethodDelete, "/list/"+strconv.Itoa(createdList.ID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/list/:id")

	// Set the path parameter with the ID of the created list entry
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(createdList.ID))

	// Execute the handler
	err = handler.DeleteList(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// TODO: Add more assertions based on the expected behavior
}
