package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wahyujatirestu/simple-procurement-system/dto"
	"github.com/wahyujatirestu/simple-procurement-system/models"
	"github.com/wahyujatirestu/simple-procurement-system/services"
	"github.com/wahyujatirestu/simple-procurement-system/utils"
)

type PurchasingController struct {
	service services.PurchasingService
}

func NewPurchasingController(s services.PurchasingService) *PurchasingController {
	return &PurchasingController{service: s}
}
func (c *PurchasingController) Create(ctx *fiber.Ctx) error {
	var req dto.CreatePurchasingRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utils.Error(ctx, 400, "invalid request")
	}

	user := ctx.Locals("user").(models.User)

	result, err := c.service.Create(ctx.Context(), user.ID, req)
	if err != nil {
		return utils.Error(ctx, 400, err.Error())
	}

	return utils.Success(ctx, 201, "purchase created", result)
}

func (c *PurchasingController) FindAll(ctx *fiber.Ctx) error {
	result, err := c.service.FindAll()
	if err != nil {
		return utils.Error(ctx, 400, "failed to fetch purchase")
	}

	return utils.Success(ctx, 200, "success", result)
}
