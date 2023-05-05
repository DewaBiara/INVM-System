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

type SupplierController struct {
	supplierService service.SupplierService
	jwtService      jwt_service.JWTService
}

func NewSupplierController(supplierService service.SupplierService) *SupplierController {
	return &SupplierController{
		supplierService: supplierService,
	}
}

func (u *SupplierController) CreateSupplier(c echo.Context) error {
	supplier := new(dto.CreateSupplierRequest)
	if err := c.Bind(supplier); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrBadRequestBody.Error())
	}

	if err := c.Validate(supplier); err != nil {
		return err
	}

	err := u.supplierService.CreateSupplier(c.Request().Context(), supplier)

	if err != nil {
		switch err {
		case utils.ErrUsernameAlreadyExist:
			fallthrough
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success creating supplier",
	})
}

func (u *SupplierController) UpdateSupplier(c echo.Context) error {

	supplier := new(dto.UpdateSupplierRequest)
	if err := c.Bind(supplier); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrBadRequestBody.Error())
	}

	if err := c.Validate(supplier); err != nil {
		return err
	}

	err := u.supplierService.UpdateSupplier(c.Request().Context(), supplier.ID, supplier)
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
		"message": "success update supplier",
	})
}

func (u *SupplierController) GetSingleSupplier(c echo.Context) error {
	supplierID := c.Param("supplier_id")
	supplier, err := u.supplierService.GetSingleSupplier(c.Request().Context(), supplierID)
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
			"message": "success getting supplier",
			"data":    supplier,
		})
	default:
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
}

func (u *SupplierController) GetPageSupplier(c echo.Context) error {

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

	supplier, err := u.supplierService.GetPageSupplier(c.Request().Context(), int(pageInt), int(limitInt))
	if err != nil {
		if err == utils.ErrDocumentNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success getting document",
		"data":    supplier,
		"meta": echo.Map{
			"page":  pageInt,
			"limit": limitInt,
		},
	})
}
