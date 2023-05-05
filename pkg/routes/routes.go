package routes

import (
	itemControllerPkg "github.com/DewaBiara/INVM-System/internal/inventory/controller"
	userControllerPkg "github.com/DewaBiara/INVM-System/internal/user/controller"
	"github.com/DewaBiara/INVM-System/pkg/utils/validation"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Routes struct {
	userController *userControllerPkg.UserController
	itemController *itemControllerPkg.ItemController
}

func NewRoutes(userController *userControllerPkg.UserController, itemController *itemControllerPkg.ItemController) *Routes {
	return &Routes{
		userController: userController,
		itemController: itemController,
	}
}

func (r *Routes) Init(e *echo.Echo, conf map[string]string) {
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Recover())

	e.Validator = &validation.CustomValidator{Validator: validator.New()}

	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(conf["JWT_SECRET"]),
	})

	v1 := e.Group("/v1")

	// Users
	users := v1.Group("/users")
	users.POST("/signup/", r.userController.SignUpUser)
	users.POST("/login/", r.userController.LoginUser)

	usersWithAuth := users.Group("", jwtMiddleware)
	usersWithAuth.GET("/", r.userController.GetBriefUsers)
	usersWithAuth.PUT("/", r.userController.UpdateUser)

	// Items
	items := v1.Group("/items")
	items.POST("/", r.itemController.CreateItem)
	items.PUT("/", r.itemController.UpdateItem)
	items.GET("/:item_id/", r.itemController.GetSingleItem)
	items.GET("/", r.itemController.GetPageItem)

	// Suppliers
	// Purchases
	// Sales
}
