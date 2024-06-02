package service

import (
	"context"
	"fmt"

	"github.com/ZoinMe/team-service/model"
	"github.com/ZoinMe/team-service/repository"
)

type CommentService struct {
	commentRepository *repository.CommentRepository
}

func NewCommentService(commentRepository *repository.CommentRepository) *CommentService {
	return &CommentService{commentRepository}
}

func (cs *CommentService) GetAllCommentsByTeamID(ctx context.Context, teamID int64) ([]*model.Comment, error) {
	comments, err := cs.commentRepository.GetAllCommentsByTeamID(ctx, teamID)
	if err != nil {
		return nil, fmt.Errorf("failed to get comments by team ID: %v", err)
	}
	return comments, nil
}

func (cs *CommentService) GetCommentByID(ctx context.Context, id int64) (*model.Comment, error) {
	comment, err := cs.commentRepository.GetCommentByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get comment by ID: %v", err)
	}
	return comment, nil
}

func (cs *CommentService) CreateComment(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	createdComment, err := cs.commentRepository.CreateComment(ctx, comment)
	if err != nil {
		return nil, fmt.Errorf("failed to create comment: %v", err)
	}
	return createdComment, nil
}

func (cs *CommentService) UpdateComment(ctx context.Context, updatedComment *model.Comment) (*model.Comment, error) {
	updatedComment, err := cs.commentRepository.UpdateComment(ctx, updatedComment)
	if err != nil {
		return nil, fmt.Errorf("failed to update comment: %v", err)
	}
	return updatedComment, nil
}

func (cs *CommentService) DeleteComment(ctx context.Context, id int64) error {
	err := cs.commentRepository.DeleteComment(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete comment: %v", err)
	}
	return nil
}
