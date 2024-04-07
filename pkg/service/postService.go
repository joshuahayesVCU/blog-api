package service

import (
	"github.com/joshuahayesVCU/blog-api/pkg/model"
	"github.com/joshuahayesVCU/blog-api/pkg/repository"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post *model.Post) error {
	// Business logic like validation would be done here
	return s.repo.Create(post)
}

func (s *PostService) GetAllPosts() ([]model.Post, error) {
	// Business logic like validation would be done here
	return s.repo.GetAll()
}
