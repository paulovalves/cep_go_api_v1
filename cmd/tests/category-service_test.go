package tests

import (
	"database/sql"
	"errors"
	"log"
	"service"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	entity "models/entity"
)

func TestGetCategories(t *testing.T) {
	log.Printf("Starting TestGetCategories")
	db, mock := setupTestDB(t)
	defer db.Close()

	// Define expected categories and mock response
	expectedCategories := []entity.Category{
		{Id: uuid.New(), Name: "Category 1", Status: "Ativo"},
		{Id: uuid.New(), Name: "Category 2", Status: "Ativo"},
		{Id: uuid.New(), Name: "Category 3", Status: "Ativo"},
	}

	// Expect the query to return categories
	rows := sqlmock.NewRows([]string{"id", "name", "status"}).
		AddRow(expectedCategories[0].Id, expectedCategories[0].Name, expectedCategories[0].Status).
		AddRow(expectedCategories[1].Id, expectedCategories[1].Name, expectedCategories[1].Status).
		AddRow(expectedCategories[2].Id, expectedCategories[2].Name, expectedCategories[2].Status)
	mock.ExpectQuery(`SELECT \* FROM "public"."categories`).WillReturnRows(rows)

	// Call the function being tested
	response := service.GetCategories()

	// Assert the response
	assert.NotNil(t, response.Data, "Expected data not to be nil")
	assert.Nil(t, response.Error, "Expected error to be nil")
	assert.Equal(t, "success", response.Message, "Expected success message")

	// Assert the data
	categories, ok := response.Data.([]entity.Category)
	assert.True(t, ok, "Expected data to be of type []entity.Category")
	assert.Equal(
		t,
		len(expectedCategories),
		len(categories),
		"Expected number of categories to match",
	)

	// Iterate over categories and assert each one
	for i := range categories {
		assert.Equal(t, expectedCategories[i].Id, categories[i].Id, "Expected category ID to match")
		assert.Equal(
			t,
			expectedCategories[i].Name,
			categories[i].Name,
			"Expected category Name to match",
		)
		assert.Equal(
			t,
			expectedCategories[i].Status,
			categories[i].Status,
			"Expected category Status to match",
		)
	}

	// Assert no unexpected calls to the mock DB
	assert.NoError(t, mock.ExpectationsWereMet(), "Expected all expectations to be met")
}

func TestGetCategoriesError(t *testing.T) {
	log.Printf("Starting TestGetCategoriesError")
	db, mock := setupTestDB(t)
	defer db.Close()

	// Expect the query to return an error
	mock.ExpectQuery(`SELECT \* FROM "public"."categories`).
		WillReturnError(errors.New("some error occurred"))

	// Call the function being tested
	response := service.GetCategories()

	// Assert the response
	assert.Nil(t, response.Data, "Expected data to be nil")
	assert.NotNil(t, response.Error, "Expected error not to be nil")
	assert.Equal(t, "Error while fetching categories", response.Message, "Expected error message")

	// Assert the error
	assert.Equal(
		t,
		"some error occurred",
		response.Error.(error).Error(),
		"Expected error message to match",
	)
}

// func TestGetCategoryById(t *testing.T) {
// 	log.Printf("Starting TestGetCategoryById")
// 	// Create a new mock database
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("An error occurred while creating mock DB: %v", err)
// 	}
// 	defer db.Close()
//
// 	// Define test cases
// 	testCases := []struct {
// 		ID            string
// 		ExpectedData  entity.Category
// 		ExpectedError error
// 	}{
// 		{
// 			ID:           "valid-id",
// 			ExpectedData: entity.Category{
// 				// Define expected category data for the given ID
// 			},
// 			ExpectedError: nil,
// 		},
// 		{
// 			ID:            "invalid-id",
// 			ExpectedData:  entity.Category{},
// 			ExpectedError: errors.New("Invalid UUID"),
// 		},
// 		{
// 			ID:            "non-existent-id",
// 			ExpectedData:  entity.Category{},
// 			ExpectedError: errors.New("record not found"),
// 		},
// 	}
//
// 	for _, tc := range testCases {
// 		mock.ExpectQuery("SELECT (.+) FROM categories WHERE id = ?").
// 			WithArgs(tc.ID).
// 			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "status"}).
// 				AddRow(tc.ExpectedData.Id, tc.ExpectedData.Name, tc.ExpectedData.Status))
// 		response := service.GetCategoryById(tc.ID)
//
// 		assert.Equal(t, tc.ExpectedData, response.Data)
// 		assert.Equal(t, tc.ExpectedError, response.Error)
//
// 		// Assert the response message
// 		if tc.ExpectedError != nil {
// 			assert.Equal(t, "Error while fetching category", response.Message)
// 		} else {
// 			assert.Equal(t, "success", response.Message)
// 		}
// 	}
//
// 	// Assert no unexpected calls to the mock DB
// 	assert.NoError(t, mock.ExpectationsWereMet(), "Expected all expectations to be met")
// }

// Mock DB setup function
func setupTestDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred while creating mock DB: %v", err)
	}

	// Initialize gorm DB with mock DB
	gormDB, err := gorm.Open(
		"postgres", db,
	)
	if err != nil {
		t.Fatalf("An error occurred while initializing gorm DB: %v", err)
	}

	// Set the mock DB to the global DB variable
	service.SetDB(gormDB)

	return db, mock
}
