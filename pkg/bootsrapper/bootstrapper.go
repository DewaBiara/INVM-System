package bootsrapper

import (
	"time"

	inventoryControllerPkg "github.com/DewaBiara/INVM-System/internal/inventory/controller"
	inventoryRepositoryPkg "github.com/DewaBiara/INVM-System/internal/inventory/repository/impl"
	inventoryServicePkg "github.com/DewaBiara/INVM-System/internal/inventory/service/impl"
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
	itemRepository := inventoryRepositoryPkg.NewItemRepositoryImpl(db)
	itemService := inventoryServicePkg.NewItemServiceImpl(itemRepository)
	itemController := inventoryControllerPkg.NewItemController(itemService, jwtService)

	//Supplier
	supplierRepository := inventoryRepositoryPkg.NewSupplierRepositoryImpl(db)
	supplierService := inventoryServicePkg.NewSupplierServiceImpl(supplierRepository)
	supplierController := inventoryControllerPkg.NewSupplierController(supplierService, jwtService)

	//Purchase
	purchaseRepository := inventoryRepositoryPkg.NewPurchaseRepositoryImpl(db)
	purchaseService := inventoryServicePkg.NewPurchaseServiceImpl(purchaseRepository)
	purchaseController := inventoryControllerPkg.NewPurchaseController(purchaseService, jwtService)

	//Sales
	saleRepository := inventoryRepositoryPkg.NewSaleRepositoryImpl(db)
	saleService := inventoryServicePkg.NewSaleServiceImpl(saleRepository)
	saleController := inventoryControllerPkg.NewSaleController(saleService, jwtService)

	route := routes.NewRoutes(userController, itemController, supplierController, purchaseController, saleController)
	route.Init(e, conf)
}
