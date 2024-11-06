package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"my-gin-app/config"
	adminrout "my-gin-app/internal/router/adminRout"
	userroute "my-gin-app/internal/router/userRoute"
)

type Router interface {
	Start()
}

type impel struct {
	gin *gin.Engine
	db  *config.Config
}

func (i *impel) Start() {

	i.gin.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},

		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * 3600,
		AllowCredentials: true,
	}))

	userroute.UserRouter(i.gin, i.db)
	adminrout.AdminRoutes(i.gin, i.db)
	i.gin.Run()
}

func NewRouter(db *config.Config) Router {
	return &impel{
		gin: gin.New(),
		db:  db,
	}
}
