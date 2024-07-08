package techstack

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ZoinMe/team-service/stores"

	"github.com/ZoinMe/team-service/model"
)

type TechStackRepository struct {
	DB *sql.DB
}

func NewTechStackRepository(db *sql.DB) stores.Techstack {
	return &TechStackRepository{DB: db}
}

func (tr *TechStackRepository) GetAll(ctx context.Context) ([]*model.TechStack, error) {
	query := "SELECT id, technology, team_id FROM tech_stacks"

	rows, err := tr.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all tech stacks: %v", err)
	}

	defer rows.Close()

	var techStacks []*model.TechStack

	for rows.Next() {
		var techStack model.TechStack
		if err := rows.Scan(
			&techStack.ID,
			&techStack.Technology,
			&techStack.TeamID,
		); err != nil {
			return nil, fmt.Errorf("failed to scan tech stack row: %v", err)
		}
		techStacks = append(techStacks, &techStack)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading tech stack rows: %v", err)
	}

	return techStacks, nil
}

func (tr *TechStackRepository) GetByID(ctx context.Context, id int64) (*model.TechStack, error) {
	query := "SELECT id, technology, team_id FROM tech_stacks WHERE id = ?"

	row := tr.DB.QueryRowContext(ctx, query, id)

	var techStack model.TechStack

	if err := row.Scan(
		&techStack.ID,
		&techStack.Technology,
		&techStack.TeamID,
	); err != nil {
		return nil, fmt.Errorf("failed to get tech stack by ID: %v", err)
	}

	return &techStack, nil
}

func (tr *TechStackRepository) Create(ctx context.Context, techStack *model.TechStack) (*model.TechStack, error) {
	query := "INSERT INTO tech_stacks (technology, team_id) VALUES (?, ?)"

	_, err := tr.DB.ExecContext(ctx, query, techStack.Technology, techStack.TeamID)
	if err != nil {
		return nil, fmt.Errorf("failed to create tech stack: %v", err)
	}

	return techStack, nil
}

func (tsr *TechStackRepository) Update(ctx context.Context, updatedTechStack *model.TechStack) (*model.TechStack, error) {
	query := "UPDATE tech_stacks SET technology=?, team_id=? WHERE id=?"

	_, err := tsr.DB.ExecContext(ctx, query, updatedTechStack.Technology, updatedTechStack.TeamID, updatedTechStack.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update tech stack: %v", err)
	}

	return updatedTechStack, nil
}

func (tr *TechStackRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM tech_stacks WHERE id = ?"

	_, err := tr.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete tech stack: %v", err)
	}

	return nil
}

func (tsr *TechStackRepository) GetTechStacksByTeamID(ctx context.Context, teamID int64) ([]*model.TechStack, error) {
	query := "SELECT id, technology, team_id FROM tech_stacks WHERE team_id = ?"

	rows, err := tsr.DB.QueryContext(ctx, query, teamID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tech stacks by team ID: %v", err)
	}

	defer rows.Close()

	var techStacks []*model.TechStack

	for rows.Next() {
		var techStack model.TechStack
		if err := rows.Scan(&techStack.ID, &techStack.Technology, &techStack.TeamID); err != nil {
			return nil, fmt.Errorf("failed to scan tech stack row: %v", err)
		}
		techStacks = append(techStacks, &techStack)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading tech stack rows: %v", err)
	}

	return techStacks, nil
}
