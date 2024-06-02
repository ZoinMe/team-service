package teamUser

import (
	"github.com/ZoinMe/team-service/service/teamUser"
	"net/http"
	"strconv"

	"github.com/ZoinMe/team-service/model"
	"github.com/gin-gonic/gin"
)

type TeamUserHandler struct {
	teamUserService *teamUser.TeamUserService
}

func NewTeamUserHandler(teamUserService *teamUser.TeamUserService) *TeamUserHandler {
	return &TeamUserHandler{teamUserService}
}

func (h *TeamUserHandler) GetTeamUsers(c *gin.Context) {
	teamUsers, err := h.teamUserService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teamUsers)
}

func (h *TeamUserHandler) AddUserToTeam(c *gin.Context) {
	var teamUser model.TeamUser
	if err := c.ShouldBindJSON(&teamUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTeamUser, err := h.teamUserService.Create(c.Request.Context(), &teamUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTeamUser)
}

func (h *TeamUserHandler) RemoveUserFromTeam(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	err = h.teamUserService.Delete(c.Request.Context(), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User removed from team successfully"})
}

func (tuh *TeamUserHandler) GetUsersByTeamID(c *gin.Context) {
	teamIDStr := c.Param("id")
	teamID, err := strconv.ParseInt(teamIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	teamUsers, err := tuh.teamUserService.GetUsersByTeamID(c.Request.Context(), teamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teamUsers)
}
