package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ZoinMe/team-service/model"
	"github.com/ZoinMe/team-service/service"
	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	teamService *service.TeamService
}

func NewTeamHandler(teamService *service.TeamService) *TeamHandler {
	return &TeamHandler{teamService}
}

func (h *TeamHandler) GetTeams(c *gin.Context) {
	teams, err := h.teamService.GetAllTeams(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teams)
}

func (h *TeamHandler) GetTeamByID(c *gin.Context) {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}
	team, err := h.teamService.GetTeamByID(c.Request.Context(), int64(teamID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Team with ID %d not found", teamID)})
		return
	}
	c.JSON(http.StatusOK, team)
}

func (h *TeamHandler) CreateTeam(c *gin.Context) {
	var team model.Team
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTeam, err := h.teamService.CreateTeam(c.Request.Context(), &team)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTeam)
}

func (h *TeamHandler) UpdateTeam(c *gin.Context) {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}
	var updatedTeam model.Team
	if err := c.ShouldBindJSON(&updatedTeam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTeam.ID = int64(teamID)
	team, err := h.teamService.UpdateTeam(c.Request.Context(), &updatedTeam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, team)
}

func (h *TeamHandler) DeleteTeam(c *gin.Context) {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}
	err = h.teamService.DeleteTeam(c.Request.Context(), int64(teamID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Team deleted successfully"})
}
