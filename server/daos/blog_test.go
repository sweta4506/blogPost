// database_test.go
package daos

import (
	"blogPost/dtos"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost_Success(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	// Create Database instance with the mock DB
	database := &Database{
		db: db,
	}

	// Set expectations on the mock
	mock.ExpectExec("INSERT INTO posts").
		WithArgs(sqlmock.AnyArg(), "Test Title", "Test Content", "Test Author", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Test the CreatePost method
	postID, err := database.CreatePost(dtos.Post{
		Title:   "Test Title",
		Content: "Test Content",
		Author:  "Test Author",
		Tags:    []string{"tag1", "tag2"},
	})

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, int32(0), postID)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreatePost_Failure(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	// Create Database instance with the mock DB
	database := &Database{
		db: db,
	}

	// Set expectations on the mock
	mock.ExpectExec("INSERT INTO posts").
		WithArgs(sqlmock.AnyArg(), "Test Title", "Test Content", "Test Author", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnError(errors.New("CreatePost query failed"))

	// Test the CreatePost method
	_, testerr := database.CreatePost(dtos.Post{
		Title:   "Test Title",
		Content: "Test Content",
		Author:  "Test Author",
		Tags:    []string{"tag1", "tag2"},
	})

	// Assertions
	assert.NotNil(t, testerr)
	assert.Contains(t, testerr.Error(), "CreatePost query failed")

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestReadPost_Success(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	// Create Database instance with the mock DB
	database := &Database{
		db: db,
	}

	// Set expectations on the mock
	mock.ExpectQuery("SELECT post_id, title, content, author, pub_date, tags FROM posts WHERE post_id = ?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"post_id", "title", "content", "author", "pub_date", "tags"}).
			AddRow(1, "Test Title", "Test Content", "Test Author", "2022-01-31", `["tag1", "tag2"]`))

	// Test the ReadPost method
	post, err := database.ReadPost(1)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, post)
	assert.Equal(t, "Test Title", post.Title)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestReadPost_Failure(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	// Create Database instance with the mock DB
	database := &Database{
		db: db,
	}

	// Set expectations on the mock
	mock.ExpectQuery("SELECT post_id, title, content, author, pub_date, tags FROM posts WHERE post_id = ?").
		WithArgs(1).
		WillReturnError(errors.New("Read query failed"))

	// Test the ReadPost method
	_, readErr := database.ReadPost(1)

	// Assertions
	assert.NotNil(t, readErr)
	assert.Contains(t, readErr.Error(), "Read query failed")

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdatePost_Success(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	// Create Database instance with the mock DB
	database := &Database{
		db: db,
	}

	// Set expectations on the mock
	mock.ExpectExec(regexp.QuoteMeta("UPDATE posts SET title = ?, content = ?, author = ?,tags=? WHERE post_id = ?")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Test the UpdatePost method
	err = database.UpdatePost(&dtos.Post{
		PostID:  1,
		Title:   "Updated Title",
		Content: "Updated Content",
		Author:  "Updated Author",
	})

	// Assertions
	assert.NoError(t, err)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestUpdatePost_Failure(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	// Create Database instance with the mock DB
	database := &Database{
		db: db,
	}

	// Set expectations on the mock
	mock.ExpectExec(regexp.QuoteMeta("UPDATE posts SET title = ?, content = ?, author = ?,tags=? WHERE post_id = ?")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 1).
		WillReturnError(errors.New("Update query failed"))

	// Test the UpdatePost method
	err = database.UpdatePost(&dtos.Post{
		PostID:  1,
		Title:   "Updated Title",
		Content: "Updated Content",
		Author:  "Updated Author",
	})

	// Assertions
	assert.NotNil(t, err)

	// Ensure all expectations were met
	assert.Contains(t, err.Error(), "Update query failed")
}

func TestDeletePost_Success(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	// Create Database instance with the mock DB
	database := &Database{
		db: db,
	}

	// Set expectations on the mock
	mock.ExpectExec("DELETE FROM posts WHERE post_id = ?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Test the DeletePost method
	err = database.DeletePost(1)

	// Assertions
	assert.NoError(t, err)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestDeletePost_Failure(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	// Create Database instance with the mock DB
	database := &Database{
		db: db,
	}

	// Set expectations on the mock
	mock.ExpectExec("DELETE FROM posts WHERE post_id = ?").
		WithArgs(1).
		WillReturnError(errors.New("Delete query failed"))

	// Test the DeletePost method
	err = database.DeletePost(1)

	// Assertions
	assert.NotNil(t, err)

	// Ensure all expectations were met
	assert.Contains(t, err.Error(), "Delete query failed")
}
