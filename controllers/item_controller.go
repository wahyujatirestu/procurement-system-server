package controllers

import (
	"strconv"

	"github.com/wahyujatirestu/simple-procurement-system/dto"
	"github.com/wahyujatirestu/simple-procurement-system/services"
	"github.com/wahyujatirestu/simple-procurement-system/utils"

	"github.com/gofiber/fiber/v2"
)

type ItemController struct {
	service services.ItemService
}

func NewItemController(s services.ItemService) *ItemController {
	return &ItemController{s}
}

func (c *ItemController) Create(ctx *fiber.Ctx) error {
	var req dto.CreateItemRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utils.Error(ctx, 400, "invalid request")
	}

	res, err := c.service.Create(req)
	if err != nil {
		return utils.Error(ctx, 400, err.Error())
	}

	return utils.Success(ctx, 201, "item created successfully", res)
}


func (c *ItemController) FindAll(ctx *fiber.Ctx) error {
	items, err := c.service.FindAll()
	if err != nil {
		return utils.Error(ctx, 500, "failed to fetch items")
	}

	return utils.Success(ctx, 200, "success", items)
}

func (c *ItemController) FindById(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	item, err := c.service.FindById(uint(id))
	if err != nil {
		return utils.Error(ctx, 500, "failed to fetch item")
	}

	return utils.Success(ctx, 200, "success", item)
	
}

func (c *ItemController) Update(ctx *fiber.Ctx) error {
	var req dto.UpdateItemRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utils.Error(ctx, 400, "invalid request")
	}

	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	item, err := c.service.Update(uint(id), req)
	if err != nil {
		return utils.Error(ctx, 500, "failed to update item")
	}

	return utils.Success(ctx, 200, "item updated successfully",item)
	
}

func (c *ItemController) Delete(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	err = c.service.Delete(uint(id))
	if err != nil {
		return utils.Error(ctx, 500, "failed to delete item")
	}

	return utils.Success(ctx, 200, "item deleted successfully", nil)
}