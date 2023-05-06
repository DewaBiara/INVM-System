package controller

import (
	"net/http"
	"strconv"

	"github.com/DewaBiara/INVM-System/internal/inventory/dto"
	"github.com/DewaBiara/INVM-System/internal/inventory/service"
	"github.com/DewaBiara/INVM-System/pkg/utils"
	"github.com/DewaBiara/INVM-System/pkg/utils/jwt_service"
	"github.com/labstack/echo/v4"
)

type PurchaseController struct {
	purchaseService service.PurchaseService
	jwtService      jwt_service.JWTService
}

func NewPurchaseController(purchaseService service.PurchaseService, jwtService jwt_service.JWTService) *PurchaseController {
	return &PurchaseController{
		purchaseService: purchaseService,
		jwtService:      jwtService,
	}
}

func (u *PurchaseController) CreatePurchase(c echo.Context) error {
	purchase := new(dto.CreatePurchaseRequest)
	if err := c.Bind(purchase); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrBadRequestBody.Error())
	}

	if err := c.Validate(purchase); err != nil {
		return err
	}

	err := u.purchaseService.CreatePurchase(c.Request().Context(), purchase)

	if err != nil {
		switch err {
		case utils.ErrUsernameAlreadyExist:
			fallthrough
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success creating purchase",
	})
}

func (u *PurchaseController) UpdatePurchase(c echo.Context) error {

	purchase := new(dto.UpdatePurchaseRequest)
	if err := c.Bind(purchase); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrBadRequestBody.Error())
	}

	if err := c.Validate(purchase); err != nil {
		return err
	}

	err := u.purchaseService.UpdatePurchase(c.Request().Context(), purchase.ID, purchase)
	if err != nil {
		switch err {
		case utils.ErrUserNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		case utils.ErrUsernameAlreadyExist:
			fallthrough
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update purchase",
	})
}

func (u *PurchaseController) GetSinglePurchase(c echo.Context) error {
	purchaseID := c.Param("purchase_id")
	purchase, err := u.purchaseService.GetSinglePurchase(c.Request().Context(), purchaseID)
	if err != nil {
		if err == utils.ErrDocumentNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	claims := u.jwtService.GetClaims(&c)
	role := claims["role"].(string)

	switch {
	case role == "pegawai":
		fallthrough
	case role == "admin":
		return c.JSON(http.StatusOK, echo.Map{
			"message": "success getting purchase",
			"data":    purchase,
		})
	default:
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
}

func (u *PurchaseController) GetPagePurchase(c echo.Context) error {

	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrInvalidNumber.Error())
	}

	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "20"
	}
	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrInvalidNumber.Error())
	}

	purchase, err := u.purchaseService.GetPagePurchase(c.Request().Context(), int(pageInt), int(limitInt))
	if err != nil {
		if err == utils.ErrDocumentNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success getting document",
		"data":    purchase,
		"meta": echo.Map{
			"page":  pageInt,
			"limit": limitInt,
		},
	})
}
