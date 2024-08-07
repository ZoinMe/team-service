package stores

import (
	"context"

	"github.com/ZoinMe/team-service/model"
)

type Request interface {
	GetAll(ctx context.Context) ([]*model.Request, error)
	GetByID(ctx context.Context, id uint) (*model.Request, error)
	Create(ctx context.Context, req *model.Request) (*model.Request, error)
	Update(ctx context.Context, updatedReq *model.Request) (*model.Request, error)
	Delete(ctx context.Context, id uint) error
	GetByTeamID(ctx context.Context, teamID int64) ([]*model.Request, error)
}

type Team interface {
	GetAll(ctx context.Context) ([]*model.Team, error)
	GetByID(ctx context.Context, id int64) (*model.Team, error)
	Create(ctx context.Context, team *model.Team) (*model.Team, error)
	Update(ctx context.Context, updatedTeam *model.Team) (*model.Team, error)
	Delete(ctx context.Context, id int64) error
}

type TeamUser interface {
	GetAll(ctx context.Context) ([]*model.TeamUser, error)
	GetByID(ctx context.Context, id uint) (*model.TeamUser, error)
	Create(ctx context.Context, teamUser *model.TeamUser) (*model.TeamUser, error)
	Delete(ctx context.Context, id uint) error
	GetUsersByTeamID(ctx context.Context, teamID string) ([]*model.TeamUser, error)
	GetTeamsByUserID(ctx context.Context, userID string) ([]*model.Team, error)
}

type Techstack interface {
	GetAll(ctx context.Context) ([]*model.TechStack, error)
	GetByID(ctx context.Context, id int64) (*model.TechStack, error)
	Create(ctx context.Context, techStack *model.TechStack) (*model.TechStack, error)
	Update(ctx context.Context, updatedTechStack *model.TechStack) (*model.TechStack, error)
	Delete(ctx context.Context, id int64) error
	GetTechStacksByTeamID(ctx context.Context, teamID int64) ([]*model.TechStack, error)
}

type Comment interface {
	GetAllCommentsByTeamID(ctx context.Context, teamID int64) ([]*model.Comment, error)
	GetByID(ctx context.Context, id int64) (*model.Comment, error)
	Create(ctx context.Context, comment *model.Comment) (*model.Comment, error)
	Update(ctx context.Context, updatedComment *model.Comment) (*model.Comment, error)
	Delete(ctx context.Context, id int64) error
}
