package routes

import (
	inventoryControllerPkg "github.com/DewaBiara/INVM-System/internal/inventory/controller"
	userControllerPkg "github.com/DewaBiara/INVM-System/internal/user/controller"
	"github.com/DewaBiara/INVM-System/pkg/utils/validation"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Routes struct {
	userController     *userControllerPkg.UserController
	itemController     *inventoryControllerPkg.ItemController
	supplierController *inventoryControllerPkg.SupplierController
	purchaseController *inventoryControllerPkg.PurchaseController
	saleController     *inventoryControllerPkg.SaleController
}

func NewRoutes(userController *userControllerPkg.UserController, itemController *inventoryControllerPkg.ItemController, supplierController *inventoryControllerPkg.SupplierController,
	purchaseController *inventoryControllerPkg.PurchaseController, saleController *inventoryControllerPkg.SaleController) *Routes {
	return &Routes{
		userController:     userController,
		itemController:     itemController,
		supplierController: supplierController,
		purchaseController: purchaseController,
		saleController:     saleController,
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
	items.POST("/", r.itemController.CreateItem, jwtMiddleware)
	items.PUT("/", r.itemController.UpdateItem, jwtMiddleware)
	items.GET("/:item_id/", r.itemController.GetSingleItem, jwtMiddleware)
	items.GET("/", r.itemController.GetPageItem)
	items.DELETE("/:item_id/", r.itemController.DeleteItem, jwtMiddleware)

	// Suppliers
	suppliers := v1.Group("/suppliers")
	suppliers.POST("/", r.supplierController.CreateSupplier, jwtMiddleware)
	suppliers.PUT("/", r.supplierController.UpdateSupplier, jwtMiddleware)
	suppliers.GET("/:supplier_id/", r.supplierController.GetSingleSupplier, jwtMiddleware)
	suppliers.GET("/", r.supplierController.GetPageSupplier)
	suppliers.DELETE("/:supplier_id/", r.supplierController.DeleteSupplier)

	// Purchases
	purchases := v1.Group("/purchases")
	purchases.POST("/", r.purchaseController.CreatePurchase)
	purchases.PUT("/", r.purchaseController.UpdatePurchase, jwtMiddleware)
	purchases.GET("/:purchase_id/", r.purchaseController.GetSinglePurchase, jwtMiddleware)
	purchases.GET("/", r.purchaseController.GetPagePurchase)
	purchases.DELETE("/:purchase_id/", r.purchaseController.DeletePurchase, jwtMiddleware)

	// Sales
	sales := v1.Group("/sales")
	sales.POST("/", r.saleController.CreateSale)
	sales.PUT("/", r.saleController.UpdateSale, jwtMiddleware)
	sales.GET("/:sale_id/", r.saleController.GetSingleSale, jwtMiddleware)
	sales.GET("/", r.saleController.GetPageSale)
	sales.DELETE("/:sale_id/", r.saleController.DeleteSale, jwtMiddleware)
}
