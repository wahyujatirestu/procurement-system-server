package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wahyujatirestu/simple-procurement-system/dto"
	"github.com/wahyujatirestu/simple-procurement-system/services"
	"github.com/wahyujatirestu/simple-procurement-system/utils"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err :=ctx.BodyParser(&req); err != nil {
		return utils.Error(ctx, 400, "invalid request")
	}

	res, err := c.authService.Register(req)
	if err != nil {
		return utils.Error(ctx, 400, err.Error())
	}

	return utils.Success(ctx, 201, "register successfully", res)
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var req dto.LoginRequest
	if err :=ctx.BodyParser(&req); err != nil {
		return utils.Error(ctx, 400, "invalid request")
	}	

	res, err := c.authService.Login(req)
	if err != nil {
		return utils.Error(ctx, 400, err.Error())
	}	
	
	return utils.Success(ctx, 200, "login successfully", res)
}