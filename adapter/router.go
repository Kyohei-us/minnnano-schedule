package http

import (
	"fmt"
	"minnnano-schedule/domain/service"
	"minnnano-schedule/infra"
	"minnnano-schedule/infra/postgresql"
	"minnnano-schedule/usecase"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	apiVersion      = "/v1"
	healthCheckRoot = "/health_check"
	// user
	usersAPIRoot  = apiVersion + "/users"
	userIDParam   = "user_id"
	userNameParam = "user_name"
)

func InitRouter() *gin.Engine {
	// e := echo.New()
	// e.Use(
	// 	middleware.Logger(),
	// 	middleware.Recover(),
	// )

	// Set up gin router
	r := gin.Default()

	// Read env vars from .env file
	sesh_key, KeyOk := os.LookupEnv("SESSION_KEY")
	sesh_secret, SecretOk := os.LookupEnv("SESSION_SECRET")
	if !KeyOk || !SecretOk {
		panic("ERROR: Cannot read env vars")
	}
	// Sessions を使用する宣言
	r.Use(sessions.Sessions(sesh_key, sessions.NewCookieStore([]byte(sesh_secret))))

	// CSS などの static files
	r.Static("/static", "./views/static")
	// Load HTML files in views
	r.LoadHTMLGlob("views/*.html")

	// user
	postgresqlConn := infra.NewPostgreSQLConnector()
	userRepository := postgresql.NewUserRepository(postgresqlConn.Conn)
	userService := service.NewUserService(userRepository)
	userUsecase := usecase.NewUserUsecase(userService)

	userGroup := r.Group(usersAPIRoot)
	{
		handler := NewUserHandler(userUsecase)
		// v1/users
		relativePath := ""
		userGroup.GET(relativePath, handler.FindUsers())
		// v1/users/{user_id}
		relativePath = fmt.Sprintf("/:%s", userIDParam)
		userGroup.GET(relativePath, handler.FindUserById())
		relativePath = ""
		userGroup.POST(relativePath, handler.AddUser())
		userGroup.DELETE(relativePath, handler.DeleteUser())
	}

	return r
}
