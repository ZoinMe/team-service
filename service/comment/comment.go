package comment

import (
	"context"
	"fmt"

	"github.com/ZoinMe/team-service/stores"

	"github.com/ZoinMe/team-service/model"
)

type CommentService struct {
	commentRepository stores.Comment
}

func NewCommentService(commentRepository stores.Comment) *CommentService {
	return &CommentService{commentRepository}
}

func (cs *CommentService) GetAllCommentsByTeamID(ctx context.Context, teamID int64) ([]*model.Comment, error) {
	comments, err := cs.commentRepository.GetAllCommentsByTeamID(ctx, teamID)
	if err != nil {
		return nil, fmt.Errorf("failed to get comments by team ID: %v", err)
	}
	return comments, nil
}

func (cs *CommentService) GetByID(ctx context.Context, id int64) (*model.Comment, error) {
	comment, err := cs.commentRepository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get comment by ID: %v", err)
	}

	return comment, nil
}

func (cs *CommentService) Create(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	createdComment, err := cs.commentRepository.Create(ctx, comment)
	if err != nil {
		return nil, fmt.Errorf("failed to create comment: %v", err)
	}

	return createdComment, nil
}

func (cs *CommentService) Update(ctx context.Context, updatedComment *model.Comment) (*model.Comment, error) {
	updatedComment, err := cs.commentRepository.Update(ctx, updatedComment)
	if err != nil {
		return nil, fmt.Errorf("failed to update comment: %v", err)
	}

	return updatedComment, nil
}

func (cs *CommentService) Delete(ctx context.Context, id int64) error {
	err := cs.commentRepository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete comment: %v", err)
	}

	return nil
}
