package teamUser

import (
	"context"
	"fmt"
	"github.com/ZoinMe/team-service/model"
	"github.com/ZoinMe/team-service/stores"
)

type TeamUserService struct {
	teamUserRepository stores.TeamUser
}

func NewTeamUserService(teamUserRepository stores.TeamUser) *TeamUserService {
	return &TeamUserService{teamUserRepository}
}

func (tus *TeamUserService) GetAll(ctx context.Context) ([]*model.TeamUser, error) {
	teamUsers, err := tus.teamUserRepository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all team users: %v", err)
	}
	return teamUsers, nil
}

func (tus *TeamUserService) GetByID(ctx context.Context, id uint) (*model.TeamUser, error) {
	teamUser, err := tus.teamUserRepository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get team user by ID: %v", err)
	}
	return teamUser, nil
}

func (tus *TeamUserService) Create(ctx context.Context, teamUser *model.TeamUser) (*model.TeamUser, error) {
	createdTeamUser, err := tus.teamUserRepository.Create(ctx, teamUser)
	if err != nil {
		return nil, fmt.Errorf("failed to create team user: %v", err)
	}
	return createdTeamUser, nil
}

func (tus *TeamUserService) Delete(ctx context.Context, id uint) error {
	err := tus.teamUserRepository.Delete(ctx, id)
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
