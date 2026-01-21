package controllers

import (
	"strconv"

	"github.com/wahyujatirestu/simple-procurement-system/dto"
	"github.com/wahyujatirestu/simple-procurement-system/services"
	"github.com/wahyujatirestu/simple-procurement-system/utils"

	"github.com/gofiber/fiber/v2"
)

type SupplierController struct {
	service services.SupplierService
}

func NewSupplierController(s services.SupplierService) *SupplierController {
	return &SupplierController{s}
}

func (c *SupplierController) Create(ctx *fiber.Ctx) error {
	var req dto.CreateSupplierRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utils.Error(ctx, 400, "invalid request")
	}

	res, err := c.service.Create(req)
	if err != nil {
		return utils.Error(ctx, 400, err.Error())
	}

	return utils.Success(ctx, 201, "supplier created successfully", res)
}


func (c *SupplierController) FindAll(ctx *fiber.Ctx) error {
	suppliers, err := c.service.FindAll()
	if err != nil {
		return utils.Error(ctx, 500, "failed to fetch suppliers")
	}

	return utils.Success(ctx, 200, "success", suppliers)
}

func (c *SupplierController) FindById(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	supplier, err := c.service.FindById(uint(id))
	if err != nil {
		return utils.Error(ctx, 500, "failed to fetch supplier")
	}

	return utils.Success(ctx, 200, "success", supplier)
}

func (c *SupplierController) Update(ctx *fiber.Ctx) error {
	var req dto.UpdateSupplierRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utils.Error(ctx, 400, "invalid request")
	}	

	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	supplier, err := c.service.Update(uint(id), req)
	if err != nil {
		return utils.Error(ctx, 500, "failed to update supplier")
	}	

	return utils.Success(ctx, 200, "supplier updated successfully",supplier)
}

func (c *SupplierController) Delete(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	err = c.service.Delete(uint(id))
	if err != nil {
		return utils.Error(ctx, 500, "failed to delete supplier")
	}

	return utils.Success(ctx, 200, "supplier deleted successfully", nil)
}