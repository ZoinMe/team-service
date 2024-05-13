package service

import (
	"context"
	"fmt"
	"github.com/ZoinMe/team-service/model"
	"github.com/ZoinMe/team-service/repository"
)

type TeamUserService struct {
	teamUserRepository *repository.TeamUserRepository
}

func NewTeamUserService(teamUserRepository *repository.TeamUserRepository) *TeamUserService {
	return &TeamUserService{teamUserRepository}
}

func (tus *TeamUserService) GetAllTeamUsers(ctx context.Context) ([]*model.TeamUser, error) {
	teamUsers, err := tus.teamUserRepository.GetAllTeamUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all team users: %v", err)
	}
	return teamUsers, nil
}

func (tus *TeamUserService) GetTeamUserByID(ctx context.Context, id uint) (*model.TeamUser, error) {
	teamUser, err := tus.teamUserRepository.GetTeamUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get team user by ID: %v", err)
	}
	return teamUser, nil
}

func (tus *TeamUserService) CreateTeamUser(ctx context.Context, teamUser *model.TeamUser) (*model.TeamUser, error) {
	createdTeamUser, err := tus.teamUserRepository.CreateTeamUser(ctx, teamUser)
	if err != nil {
		return nil, fmt.Errorf("failed to create team user: %v", err)
	}
	return createdTeamUser, nil
}

func (tus *TeamUserService) DeleteTeamUser(ctx context.Context, id uint) error {
	err := tus.teamUserRepository.DeleteTeamUser(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete team user: %v", err)
	}
	return nil
}

func (tus *TeamUserService) GetUsersByTeamID(ctx context.Context, teamID int64) ([]*model.TeamUser, error) {
	teamUsers, err := tus.teamUserRepository.GetUsersByTeamID(ctx, teamID)
	if err != nil {
		return nil, fmt.Errorf("failed to get users by team ID: %v", err)
	}
	return teamUsers, nil
}