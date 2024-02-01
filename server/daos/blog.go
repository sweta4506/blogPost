// database.go
package daos

import (
	"blogPost/dtos"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
)

//go:generate $GOPATH/bin/mockgen -source=./blog.go -destination=../../mock/mock_daos/mock_blog.go -package=mock_blog IBlogPost

type Database struct {
	db         *sql.DB
	nextPostID int32
	mu         sync.Mutex
}

func NewDatabase() (IBlogPost, error) {
	// Set up MySQL connection parameters
	dbConfig := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "blog",
		AllowNativePasswords: true,
	}

	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// // Create posts table if not exists
	// _, err = db.Exec(`
	// 	CREATE TABLE IF NOT EXISTS posts (
	// 		post_id INT AUTO_INCREMENT PRIMARY KEY,
	// 		title TEXT,
	// 		content TEXT,
	// 		author TEXT,
	// 		pub_date TEXT,
	// 		tags JSON
	// 	)
	// `)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to create table: %v", err)
	// }

	return &Database{
		db: db,
	}, nil
}

type IBlogPost interface {
	CreatePost(req dtos.Post) (int32, error)
	ReadPost(postID int32) (*dtos.Post, error)
	UpdatePost(req *dtos.Post) error
	DeletePost(postID int32) error
}

func (d *Database) CreatePost(req dtos.Post) (int32, error) {
	//fmt.Println("***********", req)
	d.mu.Lock()
	defer d.mu.Unlock()

	// Generate a unique PostID
	postID := d.nextPostID
	d.nextPostID++
	// Get the current date
	currentDate := time.Now().Format("2006-01-02")
	tagsJSON, _ := json.Marshal(req.Tags)
	// Store the post in the database
	_, err := d.db.Exec("INSERT INTO posts ( post_id,title, content, author, pub_date,tags) VALUES ( ?,?, ?, ?, ?,?)",
		postID, req.Title, req.Content, req.Author, currentDate, tagsJSON)
	if err != nil {
		fmt.Println("failed to insert post into database: ", err)
		return 0, fmt.Errorf("failed to insert post into database: %v", err)
	}

	return postID, nil
}

func (d *Database) ReadPost(postID int32) (*dtos.Post, error) {
	var post dtos.Post
	var tag []byte
	// Query the post from the database
	err := d.db.QueryRow("SELECT post_id, title, content, author, pub_date, tags FROM posts WHERE post_id = ?", postID).
		Scan(&post.PostID, &post.Title, &post.Content, &post.Author, &post.PublicationDate, &tag)

	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case when no rows are returned (post not found)
			return nil, fmt.Errorf("post not found in the database")
		}
		return nil, fmt.Errorf("failed to read post from database: %v", err)
	}

	// If the 'tags' field is NULL in the database, set it to an empty slice
	//var tagsSlice []string
	if err := json.Unmarshal([]byte(tag), &post.Tags); err != nil {
		log.Printf("Failed to unmarshal tags: %v", err)
		return nil, fmt.Errorf("failed to read post from database: %v", err)
	}

	// Update the 'Tags' field with the unmarshaled slice
	//post.Tags = tagsSlice

	return &post, nil
}

func (d *Database) UpdatePost(req *dtos.Post) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	// Update the post in the database
	tagsJSON, _ := json.Marshal(req.Tags)
	_, err := d.db.Exec("UPDATE posts SET title = ?, content = ?, author = ?,tags=? WHERE post_id = ?",
		req.Title, req.Content, req.Author, tagsJSON, req.PostID)
	if err != nil {
		return fmt.Errorf("failed to update post in database: %v", err)
	}

	return nil
}

func (d *Database) DeletePost(postID int32) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	// Delete the post from the database
	_, err := d.db.Exec("DELETE FROM posts WHERE post_id = ?", postID)
	if err != nil {
		return fmt.Errorf("failed to delete post from database: %v", err)
	}

	return nil
}
