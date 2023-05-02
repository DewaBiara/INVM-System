package bootsrapper

import (
	"time"

	userControllerPkg "github.com/DewaBiara/INVM-System/internal/user/controller"
	userRepositoryPkg "github.com/DewaBiara/INVM-System/internal/user/repository/impl"
	userServicePkg "github.com/DewaBiara/INVM-System/internal/user/service/impl"
	"github.com/DewaBiara/INVM-System/pkg/routes"
	jwtPkg "github.com/DewaBiara/INVM-System/pkg/utils/jwt_service/impl"
	passwordPkg "github.com/DewaBiara/INVM-System/pkg/utils/password/impl"
	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

func InitController(e *echo.Echo, db *gorm.DB, conf map[string]string) {
	passwordFunc := passwordPkg.NewPasswordFuncImpl()
	jwtService := jwtPkg.NewJWTService(conf["JWT_SECRET"], 1*time.Hour)

	// User
	userRepository := userRepositoryPkg.NewUserRepositoryImpl(db)
	userService := userServicePkg.NewUserServiceImpl(userRepository, passwordFunc, jwtService)
	userController := userControllerPkg.NewUserController(userService, jwtService)

	route := routes.NewRoutes(userController)
	route.Init(e, conf)
}
