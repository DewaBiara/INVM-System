package controller

import (
	"github.com/DewaBiara/INVM-System/internal/inventory/service"
)

type ItemController struct {
	inventoryService service.ItemService
}

func NewItemController(itemService service.ItemService) *ItemController {
	return &ItemController{
		itemService: itemService,
	}
}
