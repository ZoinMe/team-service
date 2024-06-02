package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ZoinMe/team-service/model"
)

type CommentRepository struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{DB: db}
}

func (cr *CommentRepository) GetAllCommentsByTeamID(ctx context.Context, teamID int64) ([]*model.Comment, error) {
	query := "SELECT id, userid, teamid, text, parentid, created_at, updated_at FROM comments WHERE teamid = ?"
	rows, err := cr.DB.QueryContext(ctx, query, teamID)
	if err != nil {
		return nil, fmt.Errorf("failed to get comments by team ID: %v", err)
	}
	defer rows.Close()

	var comments []*model.Comment
	for rows.Next() {
		var comment model.Comment
		if err := rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.TeamID,
			&comment.Text,
			&comment.ParentID,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan comment row: %v", err)
		}
		comments = append(comments, &comment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading comment rows: %v", err)
	}

	return comments, nil
}

func (cr *CommentRepository) GetCommentByID(ctx context.Context, id int64) (*model.Comment, error) {
	query := "SELECT id, userid, teamid, text, parentid, created_at, updated_at FROM comments WHERE id = ?"
	row := cr.DB.QueryRowContext(ctx, query, id)

	var comment model.Comment
	if err := row.Scan(
		&comment.ID,
		&comment.UserID,
		&comment.TeamID,
		&comment.Text,
		&comment.ParentID,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to get comment by ID: %v", err)
	}

	return &comment, nil
}

func (cr *CommentRepository) CreateComment(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()

	query := "INSERT INTO comments (userid, teamid, text, parentid, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := cr.DB.ExecContext(ctx, query, comment.UserID, comment.TeamID, comment.Text, comment.ParentID, comment.CreatedAt, comment.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create comment: %v", err)
	}
	commentID, _ := result.LastInsertId()
	comment.ID = commentID
	return comment, nil
}

func (cr *CommentRepository) UpdateComment(ctx context.Context, updatedComment *model.Comment) (*model.Comment, error) {
	updatedComment.UpdatedAt = time.Now()

	query := "UPDATE comments SET userid=?, teamid=?, text=?, parentid=?, created_at=?, updated_at=? WHERE id=?"
	_, err := cr.DB.ExecContext(ctx, query, updatedComment.UserID, updatedComment.TeamID, updatedComment.Text, updatedComment.ParentID, updatedComment.CreatedAt, updatedComment.UpdatedAt, updatedComment.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update comment: %v", err)
	}
	return updatedComment, nil
}

func (cr *CommentRepository) DeleteComment(ctx context.Context, id int64) error {
	query := "DELETE FROM comments WHERE id = ?"
	_, err := cr.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete comment: %v", err)
	}
	return nil
}
