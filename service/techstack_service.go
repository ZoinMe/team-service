package service

import (
	"context"
	"fmt"

	"github.com/ZoinMe/team-service/model"
	"github.com/ZoinMe/team-service/repository"
)

type TechStackService struct {
	techStackRepository *repository.TechStackRepository
}

func NewTechStackService(techStackRepository *repository.TechStackRepository) *TechStackService {
	return &TechStackService{techStackRepository}
}

func (tss *TechStackService) GetAllTechStacks(ctx context.Context) ([]*model.TechStack, error) {
	techStacks, err := tss.techStackRepository.GetAllTechStacks(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all tech stacks: %v", err)
	}
	return techStacks, nil
}

func (tss *TechStackService) GetTechStackByID(ctx context.Context, id int64) (*model.TechStack, error) {
	techStack, err := tss.techStackRepository.GetTechStackByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get tech stack by ID: %v", err)
	}
	return techStack, nil
}

func (tss *TechStackService) CreateTechStack(ctx context.Context, techStack *model.TechStack) (*model.TechStack, error) {
	createdTechStack, err := tss.techStackRepository.CreateTechStack(ctx, techStack)
	if err != nil {
		return nil, fmt.Errorf("failed to create tech stack: %v", err)
	}
	return createdTechStack, nil
}

func (tss *TechStackService) UpdateTechStack(ctx context.Context, updatedTechStack *model.TechStack) (*model.TechStack, error) {
	updatedTechStack, err := tss.techStackRepository.UpdateTechStack(ctx, updatedTechStack)
	if err != nil {
		return nil, fmt.Errorf("failed to update tech stack: %v", err)
	}
	return updatedTechStack, nil
}

func (tss *TechStackService) DeleteTechStack(ctx context.Context, id int64) error {
	err := tss.techStackRepository.DeleteTechStack(ctx, id)
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
