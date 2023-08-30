package infrastructure

import (
	"ginweb/adapter/controller"
	"ginweb/ent"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Store  *ent.Client
	Router *gin.Engine
}

func SetUpServer(store *ent.Client) (*Server, error) {
	server := &Server{Store: store}
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	v1 := r.Group("api/v1")
	{
		helCtrl := controller.NewHelloController()
		userCtrl := controller.NewUserController(store)
		groupCtrl := controller.NewGroupController(store)
		hello := v1.Group("/hello")
		{
			hello.GET("", helCtrl.HelloController)
		}

		users := v1.Group("/users")
		{
			users.GET("", userCtrl.GetUsers)
			users.GET("/:id", userCtrl.GetUserByID)
			users.POST("", userCtrl.CreateUser)
			users.PUT("/:id", userCtrl.UpdateUser)
			users.DELETE("/:id", userCtrl.DeleteUser)
		}

		groups := v1.Group("/groups")
		{
			groups.POST("", groupCtrl.CreateGroup)
			groups.PUT("/:id", groupCtrl.UpdateGroupName)
			groups.DELETE("/:id")
			groups.GET("/:id")
		}
	}

	server.Router = r

	return server, nil
}
