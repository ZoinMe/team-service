package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ZoinMe/team-service/model"
	"github.com/ZoinMe/team-service/service"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentService *service.CommentService
}

func NewCommentHandler(commentService *service.CommentService) *CommentHandler {
	return &CommentHandler{commentService}
}

func (h *CommentHandler) GetCommentsByTeamID(c *gin.Context) {
	teamID, err := strconv.Atoi(c.Param("team_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}
	comments, err := h.commentService.GetAllCommentsByTeamID(c.Request.Context(), int64(teamID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) GetCommentByID(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}
	comment, err := h.commentService.GetCommentByID(c.Request.Context(), int64(commentID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Comment with ID %d not found", commentID)})
		return
	}
	c.JSON(http.StatusOK, comment)
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newComment, err := h.commentService.CreateComment(c.Request.Context(), &comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newComment)
}

func (h *CommentHandler) UpdateComment(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}
	var updatedComment model.Comment
	if err := c.ShouldBindJSON(&updatedComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedComment.ID = int64(commentID)
	comment, err := h.commentService.UpdateComment(c.Request.Context(), &updatedComment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comment)
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}
	err = h.commentService.DeleteComment(c.Request.Context(), int64(commentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
