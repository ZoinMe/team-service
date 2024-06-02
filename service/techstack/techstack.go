package techstack

import (
	"context"
	"fmt"
	"github.com/ZoinMe/team-service/model"
	"github.com/ZoinMe/team-service/stores"
)

type TechStackService struct {
	techStackRepository stores.Techstack
}

func NewTechStackService(techStackRepository stores.Techstack) *TechStackService {
	return &TechStackService{techStackRepository}
}

func (tss *TechStackService) GetAll(ctx context.Context) ([]*model.TechStack, error) {
	techStacks, err := tss.techStackRepository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all tech stacks: %v", err)
	}
	return techStacks, nil
}

func (tss *TechStackService) GetByID(ctx context.Context, id int64) (*model.TechStack, error) {
	techStack, err := tss.techStackRepository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get tech stack by ID: %v", err)
	}
	return techStack, nil
}

func (tss *TechStackService) Create(ctx context.Context, techStack *model.TechStack) (*model.TechStack, error) {
	createdTechStack, err := tss.techStackRepository.Create(ctx, techStack)
	if err != nil {
		return nil, fmt.Errorf("failed to create tech stack: %v", err)
	}
	return createdTechStack, nil
}

func (tss *TechStackService) Update(ctx context.Context, updatedTechStack *model.TechStack) (*model.TechStack, error) {
	updatedTechStack, err := tss.techStackRepository.Update(ctx, updatedTechStack)
	if err != nil {
		return nil, fmt.Errorf("failed to update tech stack: %v", err)
	}
	return updatedTechStack, nil
}

func (tss *TechStackService) Delete(ctx context.Context, id int64) error {
	err := tss.techStackRepository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete tech stack: %v", err)
	}
	return nil
}

func (tss *TechStackService) GetTechStacksByTeamID(ctx context.Context, teamID int64) ([]*model.TechStack, error) {
	techStacks, err := tss.techStackRepository.GetTechStacksByTeamID(ctx, teamID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tech stacks by team ID: %v", err)
	}
	return techStacks, nil
}
