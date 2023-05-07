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

type ItemController struct {
	itemService service.ItemService
	jwtService  jwt_service.JWTService
}

func NewItemController(itemService service.ItemService, jwtService jwt_service.JWTService) *ItemController {
	return &ItemController{
		itemService: itemService,
		jwtService:  jwtService,
	}
}

func (u *ItemController) CreateItem(c echo.Context) error {
	claims := u.jwtService.GetClaims(&c)
	role := claims["role"].(string)
	if role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
	item := new(dto.CreateItemRequest)
	if err := c.Bind(item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrBadRequestBody.Error())
	}

	if err := c.Validate(item); err != nil {
		return err
	}

	err := u.itemService.CreateItem(c.Request().Context(), item)

	if err != nil {
		switch err {
		case utils.ErrItemAlreadyExist:
			fallthrough
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success creating item",
		"data":    item,
	})
}

func (u *ItemController) UpdateItem(c echo.Context) error {
	claims := u.jwtService.GetClaims(&c)
	role := claims["role"].(string)
	if role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
	item := new(dto.UpdateItemRequest)
	if err := c.Bind(item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrBadRequestBody.Error())
	}

	if err := c.Validate(item); err != nil {
		return err
	}

	err := u.itemService.UpdateItem(c.Request().Context(), item.ID, item)
	if err != nil {
		switch err {
		case utils.ErrItemNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		case utils.ErrItemAlreadyExist:
			fallthrough
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update item",
		"data":    item,
	})
}

func (u *ItemController) GetSingleItem(c echo.Context) error {
	itemID := c.Param("item_id")
	item, err := u.itemService.GetSingleItem(c.Request().Context(), itemID)
	if err != nil {
		if err == utils.ErrItemNotFound {
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
			"message": "success getting item",
			"data":    item,
		})
	default:
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
}

func (u *ItemController) GetPageItem(c echo.Context) error {

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

	item, err := u.itemService.GetPageItem(c.Request().Context(), int(pageInt), int(limitInt))
	if err != nil {
		if err == utils.ErrItemNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success getting document",
		"data":    item,
		"meta": echo.Map{
			"page":  pageInt,
			"limit": limitInt,
		},
	})
}

func (d *ItemController) DeleteItem(c echo.Context) error {
	claims := d.jwtService.GetClaims(&c)
	role := claims["role"].(string)
	if role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
	itemID := c.Param("item_id")
	err := d.itemService.DeleteItem(c.Request().Context(), itemID)
	if err != nil {
		switch err {
		case utils.ErrItemNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success deleting item",
	})
}
