package api

import (
	"fmt"
	db "github.com/carlruan/simple_bank/db/sqlc"
	"github.com/carlruan/simple_bank/token"
	"github.com/carlruan/simple_bank/util"
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Server for service
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer Create a server and set the router
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	//tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.GET("/health", server.healthCheck)
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)
	authRoutes.DELETE("/accounts/:id", server.deleteAccount)
	authRoutes.PUT("/accounts/:id", server.updateAccount)

	authRoutes.POST("/transfers", server.createTransfer)
	server.router = router
}

// Run server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) healthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, "healthy")
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
