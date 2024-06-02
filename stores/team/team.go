package team

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ZoinMe/team-service/stores"
	"time"

	"github.com/ZoinMe/team-service/model"
)

type TeamRepository struct {
	DB *sql.DB
}

func NewTeamRepository(db *sql.DB) stores.Team {
	return &TeamRepository{DB: db}
}

func (tr *TeamRepository) GetAll(ctx context.Context) ([]*model.Team, error) {
	query := "SELECT id, name, bio, profile_image_url, description, created_at, updated_at FROM teams"

	rows, err := tr.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all teams: %v", err)
	}

	defer rows.Close()

	var teams []*model.Team

	for rows.Next() {
		var team model.Team
		if err := rows.Scan(
			&team.ID,
			&team.Name,
			&team.Bio,
			&team.ProfileImageURL,
			&team.Description,
			&team.CreatedAt,
			&team.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan team row: %v", err)
		}
		teams = append(teams, &team)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading team rows: %v", err)
	}

	return teams, nil
}

func (tr *TeamRepository) GetByID(ctx context.Context, id int64) (*model.Team, error) {
	query := "SELECT id, name, bio, profile_image_url, description, created_at, updated_at FROM teams WHERE id = ?"
	row := tr.DB.QueryRowContext(ctx, query, id)

	var team model.Team

	if err := row.Scan(
		&team.ID,
		&team.Name,
		&team.Bio,
		&team.ProfileImageURL,
		&team.Description,
		&team.CreatedAt,
		&team.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to get team by ID: %v", err)
	}

	return &team, nil
}

func (tr *TeamRepository) Create(ctx context.Context, team *model.Team) (*model.Team, error) {
	team.CreatedAt = time.Now()
	team.UpdatedAt = time.Now()

	query := "INSERT INTO teams (name, bio, profile_image_url, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"

	result, err := tr.DB.ExecContext(ctx, query, team.Name, team.Bio, team.ProfileImageURL, team.Description, team.CreatedAt, team.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create team: %v", err)
	}

	teamID, _ := result.LastInsertId()
	team.ID = teamID

	return team, nil
}

func (tr *TeamRepository) Update(ctx context.Context, updatedTeam *model.Team) (*model.Team, error) {
	updatedTeam.UpdatedAt = time.Now()

	query := "UPDATE teams SET name=?, bio=?, profile_image_url=?, description=?, created_at=?, updated_at=? WHERE id=?"

	_, err := tr.DB.ExecContext(ctx, query, updatedTeam.Name, updatedTeam.Bio, updatedTeam.ProfileImageURL, updatedTeam.Description, updatedTeam.CreatedAt, updatedTeam.UpdatedAt, updatedTeam.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update team: %v", err)
	}

	return updatedTeam, nil
}

func (tr *TeamRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM teams WHERE id = ?"

	_, err := tr.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete team: %v", err)
	}

	return nil
}
