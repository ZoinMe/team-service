package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ZoinMe/team-service/model"
	"github.com/ZoinMe/team-service/service"
	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	requestService *service.RequestService
}

func NewRequestHandler(requestService *service.RequestService) *RequestHandler {
	return &RequestHandler{requestService}
}

func (h *RequestHandler) GetRequests(c *gin.Context) {
	requests, err := h.requestService.GetAllRequests(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, requests)
}

func (h *RequestHandler) GetRequestByID(c *gin.Context) {
	requestID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	request, err := h.requestService.GetRequestByID(c.Request.Context(), uint(requestID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Request with ID %d not found", requestID)})
		return
	}
	c.JSON(http.StatusOK, request)
}

func (h *RequestHandler) CreateRequest(c *gin.Context) {
	var req model.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newRequest, err := h.requestService.CreateRequest(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newRequest)
}

func (h *RequestHandler) UpdateRequest(c *gin.Context) {
	requestID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	var updatedReq model.Request
	if err := c.ShouldBindJSON(&updatedReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedReq.ID = int64(requestID)
	req, err := h.requestService.UpdateRequest(c.Request.Context(), &updatedReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *RequestHandler) DeleteRequest(c *gin.Context) {
	requestID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	err = h.requestService.DeleteRequest(c.Request.Context(), uint(requestID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Request deleted successfully"})
}

func (rh *RequestHandler) GetRequestsByTeamID(c *gin.Context) {
	teamIDStr := c.Param("id")
	teamID, err := strconv.ParseInt(teamIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	requests, err := rh.requestService.GetRequestsByTeamID(c.Request.Context(), teamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, requests)
}