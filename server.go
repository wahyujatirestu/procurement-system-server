package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wahyujatirestu/simple-procurement-system/config"
	"github.com/wahyujatirestu/simple-procurement-system/controllers"
	"github.com/wahyujatirestu/simple-procurement-system/database"
	"github.com/wahyujatirestu/simple-procurement-system/middleware"
	"github.com/wahyujatirestu/simple-procurement-system/repositories"
	"github.com/wahyujatirestu/simple-procurement-system/routes"
	"github.com/wahyujatirestu/simple-procurement-system/services"
	utilsservice "github.com/wahyujatirestu/simple-procurement-system/utils/services"
	"github.com/wahyujatirestu/simple-procurement-system/webhook"
	"gorm.io/gorm"
)	

type Server struct {
	cfg *config.Config
	app *fiber.App
	db  *gorm.DB
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := database.NewDB(cfg.DB)


	userRepo := repositories.NewUserRepository(db)
	jwtService := utilsservice.NewJwtService(cfg.JWT)
	authMw := middleware.NewAuthMiddleware(jwtService)
	authService := services.NewAuthService(userRepo, jwtService)
	authController := controllers.NewAuthController(authService)
	itemRepo := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepo)
	itemController := controllers.NewItemController(itemService)
	supplierRepo := repositories.NewSupplierRepository(db)
	supplierService := services.NewSupplierService(supplierRepo)
	supplierController := controllers.NewSupplierController(supplierService)
	webhookClient := webhook.NewClient()
	purchasingRepo := repositories.NewPurchasingRepository(db)
	purchasingDetRepo := repositories.NewPurchasingDetailRepository(db)
	purchasingService := services.NewPurchasingService(
		repositories.NewTransactionManagerRepository(db),
		purchasingRepo,
		purchasingDetRepo,
		itemRepo,
		supplierRepo,
		userRepo,
		webhookClient,
	)
	purchasingController := controllers.NewPurchasingController(purchasingService)

	app := fiber.New()

	api := app.Group("/api/v1")
	routes.AuthRoutes(api, authController)
	routes.ItemRoute(api, itemController, authMw)
	routes.SupplierRoute(api, supplierController, authMw)
	routes.PurchasingRoute(api, purchasingController, authMw)

	return &Server{
		cfg: cfg,
		app: app,
		db:  db,
	}
}

func (s *Server) Run() {
	log.Fatal(s.app.Listen(":" + s.cfg.API.Port))
}