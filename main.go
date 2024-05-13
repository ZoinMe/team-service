package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ZoinMe/team-service/handler"
	"github.com/ZoinMe/team-service/repository"
	"github.com/ZoinMe/team-service/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to the database
	db, err := sql.Open("mysql", "root:ganesh123@tcp/zoinme?parseTime=true")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize Gin router
	router := gin.Default()

	// Initialize repositories
	teamRepo := repository.NewTeamRepository(db)
	teamUserRepo := repository.NewTeamUserRepository(db)
	requestRepo := repository.NewRequestRepository(db)
	techStackRepo := repository.NewTechStackRepository(db)

	// Initialize services
	teamService := service.NewTeamService(teamRepo)
	teamUserService := service.NewTeamUserService(teamUserRepo)
	requestService := service.NewRequestService(requestRepo)
	techStackService := service.NewTechStackService(techStackRepo)

	// Initialize handlers
	teamHandler := handler.NewTeamHandler(teamService)
	teamUserHandler := handler.NewTeamUserHandler(teamUserService)
	requestHandler := handler.NewRequestHandler(requestService)
	techStackHandler := handler.NewTechStackHandler(techStackService)

	// Define APIs for each entity
	router.GET("/team", teamHandler.GetTeams)
	router.GET("/team/:id", teamHandler.GetTeamByID)
	router.POST("/team", teamHandler.CreateTeam)
	router.PUT("/team/:id", teamHandler.UpdateTeam)
	router.DELETE("/team/:id", teamHandler.DeleteTeam)

	router.GET("/teamuser", teamUserHandler.GetTeamUsers)
	router.POST("/teamuser", teamUserHandler.AddUserToTeam)
	router.DELETE("/teamuser/:id", teamUserHandler.RemoveUserFromTeam)
	router.GET("/team/:id/user", teamUserHandler.GetUsersByTeamID)

	router.GET("/request", requestHandler.GetRequests)
	router.GET("/request/:id", requestHandler.GetRequestByID)
	router.POST("/request", requestHandler.CreateRequest)
	router.PUT("/request/:id", requestHandler.UpdateRequest)
	router.DELETE("/request/:id", requestHandler.DeleteRequest)
	router.GET("/team/:id/request", requestHandler.GetRequestsByTeamID)

	router.GET("/techstack", techStackHandler.GetTechStacks)
	router.GET("/techstack/:id", techStackHandler.GetTechStackByID)
	router.POST("/techstack", techStackHandler.CreateTechStack)
	router.PUT("/techstack/:id", techStackHandler.UpdateTechStack)
	router.DELETE("/techstack/:id", techStackHandler.DeleteTechStack)
	router.GET("/team/:id/techstack", techStackHandler.GetTechStacksByTeamID)

	// Start the server
	port := ":8080"
	log.Printf("Server started on port %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
