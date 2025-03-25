package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	eventHandlers "github.com/hokamsingh/go-backend-template/internal/handlers/event"
	speakerHandlers "github.com/hokamsingh/go-backend-template/internal/handlers/speaker"
	userHandlers "github.com/hokamsingh/go-backend-template/internal/handlers/user"
	"github.com/hokamsingh/go-backend-template/internal/middleware"
	"github.com/hokamsingh/go-backend-template/internal/repository"
	"github.com/hokamsingh/go-backend-template/internal/service"
	"gorm.io/gorm"
)

// Route defines the structure for dynamic routing
type Route struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

// Controller defines the structure for a controller with routes
type Controller struct {
	Routes []Route
}

// SetupRouter dynamically sets up routes
func SetupRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode("release")

	r := gin.Default()
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Serve static files
	r.Static("/static", "./static")

	// Serve index.html at "/"
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Welcome to Go Backend",
		})
	})

	// Initialize Repositories
	userRepo := repository.NewUserRepository(db)
	eventRepo := repository.NewEventRepository(db)
	speakerRepo := repository.NewSpeakerRepository(db)

	// Initialize Services
	userService := service.NewUserService(userRepo)

	// Initialize Controllers
	userController := userHandlers.NewUserController(userService)
	eventController := eventHandlers.NewEventController(eventRepo)
	speakerController := speakerHandlers.NewSpeakerController(speakerRepo)

	// API routes
	api := r.Group("/api")
	{
		// User routes
		users := api.Group("/users")
		{
			users.GET("", userController.GetAllUsers)
			users.GET("/:id", userController.GetUserByID)
			users.POST("", userController.CreateUser)
		}

		// Event routes
		events := api.Group("/events")
		{
			events.GET("", eventController.GetAllEvents)
			events.GET("/:id", eventController.GetEventByID)
		}

		// Speaker routes
		speakers := api.Group("/speakers")
		{
			speakers.GET("", speakerController.GetAllSpeakers)
		}
	}

	return r
}
