package techstack

import (
	"fmt"
	"github.com/ZoinMe/team-service/service/techstack"
	"net/http"
	"strconv"

	"github.com/ZoinMe/team-service/model"
	"github.com/gin-gonic/gin"
)

type TechStackHandler struct {
	techStackService *techstack.TechStackService
}

func NewTechStackHandler(techStackService *techstack.TechStackService) *TechStackHandler {
	return &TechStackHandler{techStackService}
}

func (h *TechStackHandler) GetTechStacks(c *gin.Context) {
	techStacks, err := h.techStackService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, techStacks)
}

func (h *TechStackHandler) GetTechStackByID(c *gin.Context) {
	techStackID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tech stack ID"})
		return
	}
	techStack, err := h.techStackService.GetByID(c.Request.Context(), int64(techStackID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Tech stack with ID %d not found", techStackID)})
		return
	}
	c.JSON(http.StatusOK, techStack)
}

func (h *TechStackHandler) CreateTechStack(c *gin.Context) {
	var techStack model.TechStack
	if err := c.ShouldBindJSON(&techStack); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTechStack, err := h.techStackService.Create(c.Request.Context(), &techStack)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTechStack)
}

func (h *TechStackHandler) UpdateTechStack(c *gin.Context) {
	techStackID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tech stack ID"})
		return
	}
	var updatedTechStack model.TechStack
	if err := c.ShouldBindJSON(&updatedTechStack); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTechStack.ID = int64(techStackID)
	techStack, err := h.techStackService.Update(c.Request.Context(), &updatedTechStack)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, techStack)
}

func (h *TechStackHandler) DeleteTechStack(c *gin.Context) {
	techStackID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tech stack ID"})
		return
	}
	err = h.techStackService.Delete(c.Request.Context(), int64(techStackID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tech stack deleted successfully"})
}

func (tsh *TechStackHandler) GetTechStacksByTeamID(c *gin.Context) {
	teamIDStr := c.Param("id")
	teamID, err := strconv.ParseInt(teamIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	techStacks, err := tsh.techStackService.GetTechStacksByTeamID(c.Request.Context(), teamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, techStacks)
}
