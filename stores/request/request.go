package request

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ZoinMe/team-service/stores"
	"time"

	"github.com/ZoinMe/team-service/model"
)

type RequestRepository struct {
	DB *sql.DB
}

func NewRequestRepository(db *sql.DB) stores.Request {
	return &RequestRepository{DB: db}
}

func (rr *RequestRepository) GetAll(ctx context.Context) ([]*model.Request, error) {
	query := "SELECT id, user_id, team_id, status, sent_at FROM requests"
	rows, err := rr.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all requests: %v", err)
	}
	defer rows.Close()

	var requests []*model.Request
	for rows.Next() {
		var req model.Request
		if err := rows.Scan(
			&req.ID,
			&req.UserID,
			&req.TeamID,
			&req.Status,
			&req.SentAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan request row: %v", err)
		}
		requests = append(requests, &req)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading request rows: %v", err)
	}

	return requests, nil
}

func (rr *RequestRepository) GetByID(ctx context.Context, id uint) (*model.Request, error) {
	query := "SELECT id, user_id, team_id, status, sent_at FROM requests WHERE id = ?"
	row := rr.DB.QueryRowContext(ctx, query, id)

	var req model.Request
	if err := row.Scan(
		&req.ID,
		&req.UserID,
		&req.TeamID,
		&req.Status,
		&req.SentAt,
	); err != nil {
		return nil, fmt.Errorf("failed to get request by ID: %v", err)
	}

	return &req, nil
}

func (rr *RequestRepository) Create(ctx context.Context, req *model.Request) (*model.Request, error) {
	req.SentAt = time.Now()

	query := "INSERT INTO requests (user_id, team_id, status, sent_at) VALUES (?, ?, ?, ?)"
	result, err := rr.DB.ExecContext(ctx, query, req.UserID, req.TeamID, req.Status, req.SentAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	reqID, _ := result.LastInsertId()
	req.ID = reqID
	return req, nil
}

func (rr *RequestRepository) Update(ctx context.Context, updatedReq *model.Request) (*model.Request, error) {
	query := "UPDATE requests SET user_id=?, team_id=?, status=?, sent_at=? WHERE id=?"
	_, err := rr.DB.ExecContext(ctx, query, updatedReq.UserID, updatedReq.TeamID, updatedReq.Status, updatedReq.SentAt, updatedReq.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update request: %v", err)
	}
	return updatedReq, nil
}

func (rr *RequestRepository) Delete(ctx context.Context, id uint) error {
	query := "DELETE FROM requests WHERE id = ?"
	_, err := rr.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete request: %v", err)
	}
	return nil
}

func (rr *RequestRepository) GetByTeamID(ctx context.Context, teamID int64) ([]*model.Request, error) {
	query := "SELECT id, user_id, team_id, status, sent_at FROM requests WHERE team_id = ?"
	rows, err := rr.DB.QueryContext(ctx, query, teamID)
	if err != nil {
		return nil, fmt.Errorf("failed to get requests by team ID: %v", err)
	}
	defer rows.Close()

	var requests []*model.Request
	for rows.Next() {
		var request model.Request
		if err := rows.Scan(&request.ID, &request.UserID, &request.TeamID, &request.Status, &request.SentAt); err != nil {
			return nil, fmt.Errorf("failed to scan request row: %v", err)
		}
		requests = append(requests, &request)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading request rows: %v", err)
	}

	return requests, nil
}
