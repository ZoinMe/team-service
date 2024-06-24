package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	// Import the generated protobuf package
	pb "github.com/ZoinMe/team-service/proto"

	// Import your service packages
	commentHandler "github.com/ZoinMe/team-service/handler/comment"
	requestHandler "github.com/ZoinMe/team-service/handler/request"
	teamHandler "github.com/ZoinMe/team-service/handler/team"
	teamUserHandler "github.com/ZoinMe/team-service/handler/teamUser"
	techstackHandler "github.com/ZoinMe/team-service/handler/techstack"

	commentService "github.com/ZoinMe/team-service/service/comment"
	requestService "github.com/ZoinMe/team-service/service/request"
	teamService "github.com/ZoinMe/team-service/service/team"
	teamUserService "github.com/ZoinMe/team-service/service/teamUser"
	techstackService "github.com/ZoinMe/team-service/service/techstack"

	// Import your store packages
	"github.com/ZoinMe/team-service/stores/comment"
	"github.com/ZoinMe/team-service/stores/request"
	"github.com/ZoinMe/team-service/stores/team"
	"github.com/ZoinMe/team-service/stores/teamUser"
	"github.com/ZoinMe/team-service/stores/techstack"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get database environment variables
	dbuser := os.Getenv("DB_USER_AIVEN")
	dbpassword := os.Getenv("DB_PASSWORD_AIVEN")
	dbhost := os.Getenv("DB_HOST_AIVEN")
	dbport := os.Getenv("DB_PORT_AIVEN")
	dbdbname := os.Getenv("DB_NAME_AIVEN")
	grpcPort := os.Getenv("GRPC_PORT")
	port := os.Getenv("PORT")

	// Data Source Name for MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbuser, dbpassword, dbhost, dbport, dbdbname)
	// Connect to the MySQL database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Ensure the database connection is available
	err = db.Ping()
	if err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	// Initialize repositories
	teamRepo := team.NewTeamRepository(db)
	teamUserRepo := teamUser.NewTeamUserRepository(db)
	requestRepo := request.NewRequestRepository(db)
	techStackRepo := techstack.NewTechStackRepository(db)
	commentRepo := comment.NewCommentRepository(db)

	// Initialize services
	teamService := teamService.NewTeamService(teamRepo)
	teamUserService := teamUserService.NewTeamUserService(teamUserRepo)
	requestService := requestService.NewRequestService(requestRepo)
	techStackService := techstackService.NewTechStackService(techStackRepo)
	commentService := commentService.NewCommentService(commentRepo)

	// Initialize handlers
	teamHandler := teamHandler.NewTeamHandler(teamService)
	teamUserHandler := teamUserHandler.NewTeamUserHandler(teamUserService)
	requestHandler := requestHandler.NewRequestHandler(requestService)
	techStackHandler := techstackHandler.NewTechStackHandler(techStackService)
	commentHandler := commentHandler.NewCommentHandler(commentService)

	// Set up a listener for the gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the gRPC services
	pb.RegisterRequestCommentServiceServer(grpcServer, &RequestCommentServer{
		requestService: requestService,
		commentService: commentService,
	})

	// Initialize Gin router for REST API
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Allow requests from your frontend domain
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

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
	router.GET("/user/:id/team", teamUserHandler.GetTeamsByUserID)

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

	router.GET("/teams/:team_id/comments", commentHandler.GetByTeamID)
	router.GET("/comments/:id", commentHandler.GetByID)
	router.POST("/comments", commentHandler.Create)
	router.PUT("/comments/:id", commentHandler.Update)
	router.DELETE("/comments/:id", commentHandler.Delete)

	// Start servers and handle graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Run the gRPC server in a goroutine
	go func() {
		log.Println("Starting gRPC server on port :", grpcPort, "...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Run the HTTP server in a goroutine
	go func() {
		httpPort := fmt.Sprintf(":%s", port)
		log.Printf("Starting HTTP server on port %s ...", httpPort)
		if err := http.ListenAndServe(httpPort, router); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// Wait for a termination signal
	<-stop
	log.Println("Shutting down servers...")

	// Graceful stop for gRPC server
	grpcServer.GracefulStop()
}

// RequestCommentServer is the gRPC server for handling request and comment related operations.
type RequestCommentServer struct {
	pb.UnimplementedRequestCommentServiceServer
	requestService *requestService.RequestService
	commentService *commentService.CommentService
}

// CreateRequest handles the gRPC request to create a new request.
func (s *RequestCommentServer) CreateRequest(ctx context.Context, in *pb.InsertRequest) (*pb.InsertResponse, error) {
	log.Printf("Received CreateRequest: %v", in.Request)

	// Implement business logic for creating a request
	// Placeholder example - replace with actual logic
	response := &pb.InsertResponse{
		Success: true,
		Message: "Request created successfully",
	}
	return response, nil
}

// CreateComment handles the gRPC request to create a new comment.
func (s *RequestCommentServer) CreateComment(ctx context.Context, in *pb.InsertComment) (*pb.InsertResponse, error) {
	log.Printf("Received CreateComment: %v", in.Comment)

	// Implement business logic for creating a comment
	// Placeholder example - replace with actual logic
	response := &pb.InsertResponse{
		Success: true,
		Message: "Comment created successfully",
	}
	return response, nil
}
