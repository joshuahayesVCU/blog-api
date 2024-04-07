package repository

import (
	"fmt"

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
	INSERT INTO posts (title, content, author_id)
	VALUES (?, ?, ?)
	`

	// Exec executes the query without returning any rows
	// The args are for any placeholder parameter in the query.
	_, err := r.db.Exec(query, post.Title, post.Content, post.AuthorID)
	if err != nil {
		return fmt.Errorf("createPost: %v", err)
	}

	return nil
}
