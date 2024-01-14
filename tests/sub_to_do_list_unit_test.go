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
	// Setup Echo
	e := echo.New()

	// Setup PostgreSQL database connection
	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	listRepository := repository.RepositoryList(db)
	// Create repository with real database connection
	sublistRepository := repository.RepositorySublist(db)

	// Inject repository into handler
	handler := handlers.SublistHandler(sublistRepository)

	// Create a sample list entry for testing
	createdList, err := listRepository.CreateList(models.List{
		Title:       "Test Title",
		Description: "Test description",
		File:        "file.txt",
	})
	if err != nil {
		t.Fatalf("Failed to create initial list entry: %v", err)
	}

	// Create a sample sublist entry for testing
	_, err = sublistRepository.CreateSublist(models.Sublist{
		ListID:      createdList.ID,
		Title:       "title",
		Description: "Test Sublist Description",
		File:        "sublist_file.txt",
	})
	if err != nil {
		t.Fatalf("Failed to create initial sublist entry: %v", err)
	}

	// Create a request to test the handler
	req := httptest.NewRequest(http.MethodGet, "/lists/"+strconv.Itoa(createdList.ID)+"/sublist/?page=1&limit=2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/lists/:id/sublist/")

	// Set the path parameter with the ID of the created list entry
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(createdList.ID))
	// Execute the handler
	err = handler.FindSublistsByFilter(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// TODO: Add more assertions based on the expected behavior
}

func TestGetSublist(t *testing.T) {
	// Setup Echo
	e := echo.New()

	// Setup PostgreSQL database connection
	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	listRepository := repository.RepositoryList(db)

	// Create repository with real database connection
	sublistRepository := repository.RepositorySublist(db)

	// Inject repository into handler
	handler := handlers.SublistHandler(sublistRepository)

	// Create a sample list entry for testing
	createdList, err := listRepository.CreateList(models.List{
		Title:       "GetSublistTest",
		Description: "Test description",
		File:        "get_sublist_file.txt",
	})
	if err != nil {
		t.Fatalf("Failed to create initial list entry: %v", err)
	}

	// Create a sample sublist entry for testing
	createdSublist, err := sublistRepository.CreateSublist(models.Sublist{
		ListID:      createdList.ID,
		Title:       "TestSublistTitle",
		Description: "TestSublistDescription",
		File:        "get_sublist_file.txt",
	})
	if err != nil {
		t.Fatalf("Failed to create initial sublist entry: %v", err)
	}

	// Create a request to test the handler
	req := httptest.NewRequest(http.MethodGet, "/sublist/"+strconv.Itoa(createdSublist.ID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/sublist/:id")

	// Set the path parameters
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(createdSublist.ID))

	// Execute the handler
	err = handler.GetSublist(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// TODO: Add more assertions based on the expected behavior
}

func TestCreateSublist(t *testing.T) {
	// Setup Echo
	e := echo.New()

	// Setup PostgreSQL database connection
	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	listRepository := repository.RepositoryList(db)

	// Create repository with real database connection
	sublistRepository := repository.RepositorySublist(db)

	// Inject repository into handler
	handler := handlers.SublistHandler(sublistRepository)

	// Create a sample list entry for testing
	createdList, err := listRepository.CreateList(models.List{
		Title:       "CreateSublistTest",
		Description: "Test description",
		File:        "create_sublist_file.txt",
	})
	if err != nil {
		t.Fatalf("Failed to create initial list entry: %v", err)
	}

	// Create a request to test the handler
	req := httptest.NewRequest(http.MethodPost, "/sublist", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/sublist")

	c.Set("dataFile", []string{"file1.txt", "file2.txt"})

	// Set the form values
	req.Form = make(url.Values)
	req.Form.Add("title", "TestSublistTitle")
	req.Form.Add("description", "TestSublistDescription")
	req.Form.Add("list_id", strconv.Itoa(createdList.ID))

	// Execute the handler
	err = handler.CreateSublist(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// TODO: Add more assertions based on the expected behavior
}

func TestUpdateSublist(t *testing.T) {
	// Setup Echo
	e := echo.New()

	// Setup PostgreSQL database connection
	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	listRepository := repository.RepositoryList(db)

	createdList, err := listRepository.CreateList(models.List{
		Title:       "CreateSublistTest",
		Description: "Test description",
		File:        "create_sublist_file.txt",
	})
	if err != nil {
		t.Fatalf("Failed to create initial list entry: %v", err)
	}

	// Create repository with real database connection
	sublistRepository := repository.RepositorySublist(db)

	// Inject repository into handler
	handler := handlers.SublistHandler(sublistRepository)

	// Create a sample sublist entry for testing
	createdSublist, err := sublistRepository.CreateSublist(models.Sublist{
		Title:       "UpdateSublistTest",
		Description: "Test description",
		File:        "update_sublist_file.txt",
		ListID:      createdList.ID, // Replace with the actual list ID
	})
	if err != nil {
		t.Fatalf("Failed to create initial sublist entry: %v", err)
	}

	// Create a request to test the handler
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/sublist/%d", createdSublist.ID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/sublist/:id")

	updatePayload := subtodolistdto.UpdateSublist{
		Title:       "UpdatedTitle",
		Description: "UpdatedDescription",
	}

	// Set the form values
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(createdSublist.ID))
	c.Set("dataFile", []string{"file1.txt", "file2.txt"})
	c.SetRequest(req)
	req.Form = make(url.Values)
	req.Form.Add("title", updatePayload.Title)
	req.Form.Add("description", updatePayload.Description)

	// Execute the handler
	err = handler.UpdateSublist(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// TODO: Add more assertions based on the expected behavior
}

func TestDeleteSublist(t *testing.T) {
	// Setup Echo
	e := echo.New()

	// Setup PostgreSQL database connection
	dsn := "user=postgres password=1234 dbname=moonlay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	listRepository := repository.RepositoryList(db)

	createdList, err := listRepository.CreateList(models.List{
		Title:       "CreateSublistTest",
		Description: "Test description",
		File:        "create_sublist_file.txt",
	})
	if err != nil {
		t.Fatalf("Failed to create initial list entry: %v", err)
	}

	// Create repository with real database connection
	sublistRepository := repository.RepositorySublist(db)

	// Inject repository into handler
	handler := handlers.SublistHandler(sublistRepository)

	// Create a sample sublist entry for testing
	createdSublist, err := sublistRepository.CreateSublist(models.Sublist{
		Title:       "DeleteSublistTest",
		Description: "Test description",
		File:        "delete_sublist_file.txt",
		ListID:      createdList.ID,
	})
	if err != nil {
		t.Fatalf("Failed to create initial sublist entry: %v", err)
	}

	// Create a request to test the handler
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/sublist/%d", createdSublist.ID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/sublist/:id")

	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(createdSublist.ID))
	// Execute the handler
	err = handler.DeleteSublist(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// TODO: Add more assertions based on the expected behavior
}
