package service

import (
	"context"
	"fmt"
	"github.com/ZoinMe/team-service/model"
	"github.com/ZoinMe/team-service/repository"
)

type TeamService struct {
	teamRepository *repository.TeamRepository
}

func NewTeamService(teamRepository *repository.TeamRepository) *TeamService {
	return &TeamService{teamRepository}
}

func (ts *TeamService) GetAllTeams(ctx context.Context) ([]*model.Team, error) {
	teams, err := ts.teamRepository.GetAllTeams(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all teams: %v", err)
	}
	return teams, nil
}

func (ts *TeamService) GetTeamByID(ctx context.Context, id int64) (*model.Team, error) {
	team, err := ts.teamRepository.GetTeamByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get team by ID: %v", err)
	}
	return team, nil
}

func (ts *TeamService) CreateTeam(ctx context.Context, team *model.Team) (*model.Team, error) {
	createdTeam, err := ts.teamRepository.CreateTeam(ctx, team)
	if err != nil {
		return nil, fmt.Errorf("failed to create team: %v", err)
	}
	return createdTeam, nil
}

func (ts *TeamService) UpdateTeam(ctx context.Context, updatedTeam *model.Team) (*model.Team, error) {
	updatedTeam, err := ts.teamRepository.UpdateTeam(ctx, updatedTeam)
	if err != nil {
		return nil, fmt.Errorf("failed to update team: %v", err)
	}
	return updatedTeam, nil
}

func (ts *TeamService) DeleteTeam(ctx context.Context, id int64) error {
	err := ts.teamRepository.DeleteTeam(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete team: %v", err)
	}
	return nil
}
