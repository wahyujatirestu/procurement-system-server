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

// Register godoc
// @Summary Register new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Register payload"
// @Success 201 {object} dto.AuthResponse
// @Failure 400 {object} map[string]string
// @Router /auth/register [post]
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

// Login godoc
// @Summary Login user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login payload"
// @Success 201 {object} dto.AuthResponse
// @Failure 400 {object} map[string]string
// @Router /auth/login [post]
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

// Logout godoc
// @Summary Logout user
// @Tags Auth
// @Accept json
// @Produce json
// @Success 201 {object} dto.AuthResponse
// @Failure 400 {object} map[string]string
// @Router /auth/logout [post]
func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	return utils.Success(ctx, 200, "logout successfully", nil)
}