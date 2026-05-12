package transport

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SussyaPusya/swiftTalk/auth-service/internal/models"
	"github.com/SussyaPusya/swiftTalk/auth-service/internal/pkg"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateUser(ctx context.Context, user *models.CreateUserDTO) error
	Login(ctx context.Context, login *models.LoginDTO) (*models.GetUserDTO, error)
	GetUserByID(ctx context.Context, userID string) (*models.GetUserDTO, error)
	ChangePassword(ctx context.Context, userID string, newPassword string) error
	ChangeName(ctx context.Context, userID string, newName string) error
}
type router struct {
	s          Service
	j          *pkg.JWTManager
	middleware *Middleware
}

func NewRouter(s Service, middleware *Middleware, j *pkg.JWTManager) *router {
	return &router{s: s, middleware: middleware, j: j}
}

func (r *router) Run() {
	ginRouter := gin.Default()

	api := ginRouter.Group("/api")
	{
		api.POST("/login", r.login)
		api.POST("/register", r.createUser)
		users := ginRouter.Group("/users")
		{

			users.GET("/:userID", r.getUserByID)
			users.PUT("/:userID/password", r.changePassword)
			users.PUT("/:userID/name", r.changeName)
		}
		users.Use(r.middleware.AuthMiddleware())
	}

	ginRouter.Run(":8081")
}

func (r *router) login(c *gin.Context) {
	var req models.LoginDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(req)

	user, err := r.s.Login(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := r.j.GenerateAccessToken(user.Username, user.ID, models.AccessTokenTTL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := r.j.GenerateRefreshToken(user.Username, user.ID, models.RefreshTokenTTL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := models.LoginResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, resp)

}

func (r *router) createUser(c *gin.Context) {
	var req models.CreateUserDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := r.s.CreateUser(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created"})

}

func (r *router) getUserByID(c *gin.Context) {
	userID := c.Param("userID")

	user, err := r.s.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)

}

func (r *router) changePassword(c *gin.Context) {
	userID := c.Param("userID")
	var req models.ChangePasswordDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := r.s.ChangePassword(c.Request.Context(), userID, req.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password changed"})

}

func (r *router) changeName(c *gin.Context) {
	userID := c.Param("userID")
	var req models.ChangeNameDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := r.s.ChangeName(c.Request.Context(), userID, req.NewName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "name changed"})

}
