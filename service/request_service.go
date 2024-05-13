package service

import (
	"context"
	"fmt"
	"github.com/ZoinMe/team-service/model"
	"github.com/ZoinMe/team-service/repository"
)

type RequestService struct {
	requestRepository *repository.RequestRepository
}

func NewRequestService(requestRepository *repository.RequestRepository) *RequestService {
	return &RequestService{requestRepository}
}

func (rs *RequestService) GetAllRequests(ctx context.Context) ([]*model.Request, error) {
	requests, err := rs.requestRepository.GetAllRequests(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all requests: %v", err)
	}
	return requests, nil
}

func (rs *RequestService) GetRequestByID(ctx context.Context, id uint) (*model.Request, error) {
	request, err := rs.requestRepository.GetRequestByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get request by ID: %v", err)
	}
	return request, nil
}

func (rs *RequestService) CreateRequest(ctx context.Context, req *model.Request) (*model.Request, error) {
	createdRequest, err := rs.requestRepository.CreateRequest(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	return createdRequest, nil
}

func (rs *RequestService) UpdateRequest(ctx context.Context, updatedRequest *model.Request) (*model.Request, error) {
	updatedRequest, err := rs.requestRepository.UpdateRequest(ctx, updatedRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to update request: %v", err)
	}
	return updatedRequest, nil
}

func (rs *RequestService) DeleteRequest(ctx context.Context, id uint) error {
	err := rs.requestRepository.DeleteRequest(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete request: %v", err)
	}
	return nil
}

func (rs *RequestService) GetRequestsByTeamID(ctx context.Context, teamID int64) ([]*model.Request, error) {
	requests, err := rs.requestRepository.GetRequestsByTeamID(ctx, teamID)
	if err != nil {
		return nil, fmt.Errorf("failed to get requests by team ID: %v", err)
	}
	return requests, nil
}