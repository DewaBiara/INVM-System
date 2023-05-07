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

type SaleController struct {
	saleService service.SaleService
	jwtService  jwt_service.JWTService
}

func NewSaleController(saleService service.SaleService, jwtService jwt_service.JWTService) *SaleController {
	return &SaleController{
		saleService: saleService,
		jwtService:  jwtService,
	}
}

func (u *SaleController) CreateSale(c echo.Context) error {
	sale := new(dto.CreateSaleRequest)
	if err := c.Bind(sale); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrBadRequestBody.Error())
	}

	if err := c.Validate(sale); err != nil {
		return err
	}

	err := u.saleService.CreateSale(c.Request().Context(), sale)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success creating sale",
	})
}

func (u *SaleController) UpdateSale(c echo.Context) error {
	claims := u.jwtService.GetClaims(&c)
	role := claims["role"].(string)
	if role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
	sale := new(dto.UpdateSaleRequest)
	if err := c.Bind(sale); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrBadRequestBody.Error())
	}

	if err := c.Validate(sale); err != nil {
		return err
	}

	err := u.saleService.UpdateSale(c.Request().Context(), sale.ID, sale)
	if err != nil {
		switch err {
		case utils.ErrSaleNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update sale",
	})
}

func (u *SaleController) GetSingleSale(c echo.Context) error {
	saleID := c.Param("sale_id")
	sale, err := u.saleService.GetSingleSale(c.Request().Context(), saleID)
	if err != nil {
		if err == utils.ErrSaleNotFound {
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
			"message": "success getting sale",
			"data":    sale,
		})
	default:
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
}

func (u *SaleController) GetPageSale(c echo.Context) error {

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

	sale, err := u.saleService.GetPageSale(c.Request().Context(), int(pageInt), int(limitInt))
	if err != nil {
		if err == utils.ErrSaleNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success getting document",
		"data":    sale,
		"meta": echo.Map{
			"page":  pageInt,
			"limit": limitInt,
		},
	})
}

func (d *SaleController) DeleteSale(c echo.Context) error {
	claims := d.jwtService.GetClaims(&c)
	role := claims["role"].(string)
	if role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
	saleID := c.Param("sale_id")
	err := d.saleService.DeleteSale(c.Request().Context(), saleID)
	if err != nil {
		switch err {
		case utils.ErrSaleNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success deleting sale",
	})
}
