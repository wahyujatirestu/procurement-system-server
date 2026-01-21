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

// Create Supplier
// @Summary Create new supplier
// @Tags Suppliers
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CreateSupplierRequest true "Supplier payload"
// @Success 201 {object} dto.SupplierResponse
// @Failure 400 {object} dto.SupplierResponse
// @Failure 500 {object} dto.SupplierResponse
// @Router /suppliers [post]
func (c *SupplierController) Create(ctx *fiber.Ctx) error {
	var req dto.CreateSupplierRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utils.Error(ctx, 400, "invalid request")
	}

	res, err := c.service.Create(req)
	if err != nil {
		return utils.Error(ctx, 500, err.Error())
	}

	return utils.Success(ctx, 201, "supplier created successfully", res)
}

// Get Suppliers
// @Summary Get list of suppliers
// @Tags Suppliers
// @Security BearerAuth
// @Produce json
// @Success 200 {array} dto.SupplierResponse
// @Failure 500 {object} dto.SupplierResponse
// @Router /suppliers [get]
func (c *SupplierController) FindAll(ctx *fiber.Ctx) error {
	suppliers, err := c.service.FindAll()
	if err != nil {
		return utils.Error(ctx, 500, "failed to fetch suppliers")
	}

	return utils.Success(ctx, 200, "success", suppliers)
}

// Get Supplier
// @Summary Get supplier by id
// @Tags Suppliers
// @Security BearerAuth
// @Produce json
// @Param id path int true "Supplier ID"
// @Success 200 {object} dto.SupplierResponse
// @Failure 500 {object} dto.SupplierResponse
// @Router /suppliers/{id} [get]
func (c *SupplierController) FindById(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	supplier, err := c.service.FindById(uint(id))
	if err != nil {
		return utils.Error(ctx, 500, "failed to fetch supplier")
	}

	return utils.Success(ctx, 200, "success", supplier)
}

// Update Supplier
// @Summary Update supplier by id
// @Tags Suppliers
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Supplier ID"
// @Param request body dto.UpdateSupplierRequest true "Supplier payload"
// @Success 200 {object} dto.SupplierResponse
// @Failure 400 {object} dto.SupplierResponse
// @Failure 500 {object} dto.SupplierResponse
// @Router /suppliers/{id} [put]
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

// Delete Supplier
// @Summary Delete supplier by id
// @Tags Suppliers
// @Security BearerAuth
// @Produce json
// @Param id path int true "Supplier ID"
// @Success 200 {object} dto.SupplierResponse
// @Failure 404 {object} dto.SupplierResponse
// @Router /suppliers/{id} [delete]
func (c *SupplierController) Delete(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	err = c.service.Delete(uint(id))
	if err != nil {
		return utils.Error(ctx, 500, "failed to delete supplier")
	}

	return utils.Success(ctx, 200, "supplier deleted successfully", nil)
}