package request

import (
	"context"
	"fmt"
	"github.com/ZoinMe/team-service/model"
	"github.com/ZoinMe/team-service/stores"
)

type RequestService struct {
	requestRepository stores.Request
}

func NewRequestService(requestRepository stores.Request) *RequestService {
	return &RequestService{requestRepository}
}

func (rs *RequestService) GetAll(ctx context.Context) ([]*model.Request, error) {
	requests, err := rs.requestRepository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all requests: %v", err)
	}
	return requests, nil
}

func (rs *RequestService) GetByID(ctx context.Context, id uint) (*model.Request, error) {
	request, err := rs.requestRepository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get request by ID: %v", err)
	}
	return request, nil
}

func (rs *RequestService) Create(ctx context.Context, req *model.Request) (*model.Request, error) {
	createdRequest, err := rs.requestRepository.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	return createdRequest, nil
}

func (rs *RequestService) Update(ctx context.Context, updatedRequest *model.Request) (*model.Request, error) {
	updatedRequest, err := rs.requestRepository.Update(ctx, updatedRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to update request: %v", err)
	}
	return updatedRequest, nil
}

func (rs *RequestService) Delete(ctx context.Context, id uint) error {
	err := rs.requestRepository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete request: %v", err)
	}
	return nil
}

func (rs *RequestService) GetByTeamID(ctx context.Context, teamID int64) ([]*model.Request, error) {
	requests, err := rs.requestRepository.GetByTeamID(ctx, teamID)
	if err != nil {
		return nil, fmt.Errorf("failed to get requests by team ID: %v", err)
	}
	return requests, nil
}
