package repository

import (
	"fmt"
	"log"

	"github.com/joshuahayesVCU/blog-api/pkg/model"
)

type PostRepository struct {
	db *model.DB
}

func NewPostRepository(db *model.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(post *model.Post) error {
	/// SQL statement to insert a new post
	query := `
	INSERT INTO Posts (UserID, Title, Content)
	VALUES (?, ?, ?)
	`

	// Exec executes the query without returning any rows
	// The args are for any placeholder parameter in the query.
	_, err := r.db.Exec(query, post.UserID, post.Title, post.Content)
	if err != nil {
		log.Printf("Failed to create post: %v", err)
		return fmt.Errorf("createPost: %v", err)
	}

	return nil
}

func (r *PostRepository) GetAll() ([]model.Post, error) {
	// SQL statement to select all post
	query := `
	SELECT PostID, UserID, Title, Content, CreatedAt, UpdatedAt FROM Posts
	`

	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Failed to get posts: %v", err)
		return nil, fmt.Errorf("getPosts: %v", err)
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		// Scan each row's columns into the Post struct.
		// Order of arguments must match SQL query
		if err := rows.Scan(&post.PostID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt); err != nil {
			log.Printf("Error scanning post: %v", err)
			return nil, fmt.Errorf("getPosts: %v", err)
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error encountered during rows processing: %v", err)
		return nil, fmt.Errorf("getPosts: %v", err)
	}

	return posts, nil
}
