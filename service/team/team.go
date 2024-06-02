package team

import (
	"context"
	"fmt"
	"github.com/ZoinMe/team-service/model"
	"github.com/ZoinMe/team-service/stores"
)

type TeamService struct {
	teamRepository stores.Team
}

func NewTeamService(teamRepository stores.Team) *TeamService {
	return &TeamService{teamRepository}
}

func (ts *TeamService) GetAll(ctx context.Context) ([]*model.Team, error) {
	teams, err := ts.teamRepository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all teams: %v", err)
	}
	return teams, nil
}

func (ts *TeamService) GetByID(ctx context.Context, id int64) (*model.Team, error) {
	team, err := ts.teamRepository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get team by ID: %v", err)
	}
	return team, nil
}

func (ts *TeamService) Create(ctx context.Context, team *model.Team) (*model.Team, error) {
	createdTeam, err := ts.teamRepository.Create(ctx, team)
	if err != nil {
		return nil, fmt.Errorf("failed to create team: %v", err)
	}
	return createdTeam, nil
}

func (ts *TeamService) Update(ctx context.Context, updatedTeam *model.Team) (*model.Team, error) {
	updatedTeam, err := ts.teamRepository.Update(ctx, updatedTeam)
	if err != nil {
		return nil, fmt.Errorf("failed to update team: %v", err)
	}
	return updatedTeam, nil
}

func (ts *TeamService) Delete(ctx context.Context, id int64) error {
	err := ts.teamRepository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete team: %v", err)
	}
	return nil
}
