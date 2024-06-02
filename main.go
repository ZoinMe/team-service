package main

import (
	"database/sql"
	request3 "github.com/ZoinMe/team-service/handler/request"
	team3 "github.com/ZoinMe/team-service/handler/team"
	teamUser3 "github.com/ZoinMe/team-service/handler/teamUser"
	techstack3 "github.com/ZoinMe/team-service/handler/techstack"
	request2 "github.com/ZoinMe/team-service/service/request"
	team2 "github.com/ZoinMe/team-service/service/team"
	teamUser2 "github.com/ZoinMe/team-service/service/teamUser"
	techstack2 "github.com/ZoinMe/team-service/service/techstack"
	"github.com/ZoinMe/team-service/stores/request"
	"github.com/ZoinMe/team-service/stores/team"
	"github.com/ZoinMe/team-service/stores/teamUser"
	"github.com/ZoinMe/team-service/stores/techstack"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to the database
	db, err := sql.Open("mysql", "root:password@tcp/zoinme?parseTime=true")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	defer db.Close()

	// Initialize Gin router
	router := gin.Default()

	// Initialize repositories
	teamRepo := team.NewTeamRepository(db)
	teamUserRepo := teamUser.NewTeamUserRepository(db)
	requestRepo := request.NewRequestRepository(db)
	techStackRepo := techstack.NewTechStackRepository(db)

	// Initialize services
	teamService := team2.NewTeamService(teamRepo)
	teamUserService := teamUser2.NewTeamUserService(teamUserRepo)
	requestService := request2.NewRequestService(requestRepo)
	techStackService := techstack2.NewTechStackService(techStackRepo)

	// Initialize handlers
	teamHandler := team3.NewTeamHandler(teamService)
	teamUserHandler := teamUser3.NewTeamUserHandler(teamUserService)
	requestHandler := request3.NewRequestHandler(requestService)
	techStackHandler := techstack3.NewTechStackHandler(techStackService)

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
