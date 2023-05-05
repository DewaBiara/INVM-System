package bootsrapper

import (
	"time"

	itemControllerPkg "github.com/DewaBiara/INVM-System/internal/inventory/controller"
	itemRepositoryPkg "github.com/DewaBiara/INVM-System/internal/inventory/repository/impl"
	itemServicePkg "github.com/DewaBiara/INVM-System/internal/inventory/service/impl"
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

	//Item
	itemRepository := itemRepositoryPkg.NewItemRepositoryImpl(db)
	itemService := itemServicePkg.NewItemServiceImpl(itemRepository)
	itemController := itemControllerPkg.NewItemController(itemService)

	route := routes.NewRoutes(userController, itemController)
	route.Init(e, conf)
}
