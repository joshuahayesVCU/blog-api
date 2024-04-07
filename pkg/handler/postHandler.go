package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshuahayesVCU/blog-api/pkg/model"
	"github.com/joshuahayesVCU/blog-api/pkg/service"
)

type PostHandler struct {
	service *service.PostService
}

type createPostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  int64  `json:"user_id" binding:"required"`
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var req createPostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// create a model.Post object from the request
	post := model.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserID,
	}

	// call the service layer to create the post
	err := h.service.CreatePost(&post)
	if err != nil {
		// Handle potential errors from the service layer
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	// If successful, return a success message and a 201
	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
}

func (h *PostHandler) GetAllPosts(c *gin.Context) {
	posts, err := h.service.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, posts)
}
