package teamUser

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ZoinMe/team-service/stores"

	"github.com/ZoinMe/team-service/model"
)

type TeamUserRepository struct {
	DB *sql.DB
}

func NewTeamUserRepository(db *sql.DB) stores.TeamUser {
	return &TeamUserRepository{DB: db}
}

func (tur *TeamUserRepository) GetAll(ctx context.Context) ([]*model.TeamUser, error) {
	query := "SELECT id, team_id, user_id, join_date, role FROM team_users"

	rows, err := tur.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all team users: %v", err)
	}

	defer rows.Close()

	var teamUsers []*model.TeamUser

	for rows.Next() {
		var teamUser model.TeamUser
		if err := rows.Scan(
			&teamUser.ID,
			&teamUser.TeamID,
			&teamUser.UserID,
			&teamUser.JoinDate,
			&teamUser.Role,
		); err != nil {
			return nil, fmt.Errorf("failed to scan team user row: %v", err)
		}
		teamUsers = append(teamUsers, &teamUser)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading team user rows: %v", err)
	}

	return teamUsers, nil
}

func (tur *TeamUserRepository) GetByID(ctx context.Context, id uint) (*model.TeamUser, error) {
	query := "SELECT id, team_id, user_id, join_date, role FROM team_users WHERE id = ?"
	row := tur.DB.QueryRowContext(ctx, query, id)

	var teamUser model.TeamUser

	if err := row.Scan(
		&teamUser.ID,
		&teamUser.TeamID,
		&teamUser.UserID,
		&teamUser.JoinDate,
		&teamUser.Role,
	); err != nil {
		return nil, fmt.Errorf("failed to get team user by ID: %v", err)
	}

	return &teamUser, nil
}

func (tur *TeamUserRepository) Create(ctx context.Context, teamUser *model.TeamUser) (*model.TeamUser, error) {
	query := "INSERT INTO team_users (team_id, user_id, join_date, role) VALUES (?, ?, ?, ?)"
	result, err := tur.DB.ExecContext(ctx, query, teamUser.TeamID, teamUser.UserID, teamUser.JoinDate, teamUser.Role)
	if err != nil {
		return nil, fmt.Errorf("failed to create team user: %v", err)
	}

	teamUserID, _ := result.LastInsertId()
	teamUser.ID = teamUserID

	return teamUser, nil
}

func (tur *TeamUserRepository) Delete(ctx context.Context, id uint) error {
	query := "DELETE FROM team_users WHERE id = ?"

	_, err := tur.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete team user: %v", err)
	}

	return nil
}

func (tur *TeamUserRepository) GetUsersByTeamID(ctx context.Context, teamID int64) ([]*model.TeamUser, error) {
	query := "SELECT id, user_id, team_id, join_date, role FROM team_users WHERE team_id = ?"

	rows, err := tur.DB.QueryContext(ctx, query, teamID)
	if err != nil {
		return nil, fmt.Errorf("failed to get users by team ID: %v", err)
	}

	defer rows.Close()

	var teamUsers []*model.TeamUser

	for rows.Next() {
		var teamUser model.TeamUser
		if err := rows.Scan(&teamUser.ID, &teamUser.UserID, &teamUser.TeamID, &teamUser.JoinDate, &teamUser.Role); err != nil {
			return nil, fmt.Errorf("failed to scan team user row: %v", err)
		}
		teamUsers = append(teamUsers, &teamUser)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading team user rows: %v", err)
	}

	return teamUsers, nil
}

func (tur *TeamUserRepository) GetTeamsByUserID(ctx context.Context, userID int64) ([]*model.Team, error) {
	query := `
        SELECT t.id, t.name, t.description, t.bio, t.profile_image_url, t.created_at, t.updated_at
        FROM teams t
        JOIN team_users tu ON t.id = tu.team_id
        WHERE tu.user_id = ?
    `

	rows, err := tur.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get teams by user ID: %v", err)
	}

	defer rows.Close()

	var teams []*model.Team

	for rows.Next() {
		var team model.Team
		if err := rows.Scan(&team.ID, &team.Name, &team.Description, &team.Bio, &team.ProfileImageURL, &team.CreatedAt, &team.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan team row: %v", err)
		}
		teams = append(teams, &team)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading team rows: %v", err)
	}

	return teams, nil
}
