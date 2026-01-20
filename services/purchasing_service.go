package services

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"github.com/wahyujatirestu/simple-procurement-system/dto"
	"github.com/wahyujatirestu/simple-procurement-system/models"
	"github.com/wahyujatirestu/simple-procurement-system/repositories"
	"github.com/wahyujatirestu/simple-procurement-system/webhook"
)

type PurchasingService interface {
	Create(ctx context.Context, userID uint, req dto.CreatePurchasingRequest) (*dto.PurchasingResponse, error)
	FindAll() ([]dto.PurchasingResponse, error)
}

type purchasingService struct {
	txManager        repositories.TransactionManagerRepository
	purchasingRepo   repositories.PurchasingRepository
	purchasingDetRepo repositories.PurchasingDetailRepository
	itemRepo         repositories.ItemRepository
	supplierRepo     repositories.SupplierRepository
	userRepo         repositories.UserRepository
	webhookClient    webhook.Client
}


func NewPurchasingService(
	txManager repositories.TransactionManagerRepository,
	purchasingRepo repositories.PurchasingRepository,
	purchasingDetRepo repositories.PurchasingDetailRepository,
	itemRepo repositories.ItemRepository,
	supplierRepo repositories.SupplierRepository,
	userRepo repositories.UserRepository,
	webhookClient webhook.Client,
) PurchasingService {
	return &purchasingService{
		txManager:        txManager,
		purchasingRepo:   purchasingRepo,
		purchasingDetRepo: purchasingDetRepo,
		itemRepo:         itemRepo,
		supplierRepo:     supplierRepo,
		userRepo:         userRepo,
		webhookClient:    webhookClient,
	}
}

func (s *purchasingService) Create(
	ctx context.Context,
	userID uint,
	req dto.CreatePurchasingRequest,
) (*dto.PurchasingResponse, error) {

	var response *dto.PurchasingResponse
	var webhookPayload dto.PurchaseWebhookPayload

	err := s.txManager.WithTransaction(ctx, func(tx *gorm.DB) error {

		user, err := s.userRepo.FindById(userID)
		if err != nil {
			return errors.New("user not found")
		}

		supplier, err := s.supplierRepo.FindById(req.SupplierID)
		if err != nil {
			return errors.New("supplier not found")
		}

		purchasing := models.Purchasing{
			Date:       time.Now(),
			SupplierID: supplier.ID,
			UserID:     user.ID,
		}

		if err := tx.Create(&purchasing).Error; err != nil {
			return err
		}

		var (
			grandTotal float64
			itemsResp  []dto.PurchasingDetailResponse
			webhookItems []dto.PurchaseWebhookItem
		)

		for _, reqItem := range req.Items {
			var item models.Item
			if err := tx.First(&item, reqItem.ItemID).Error; err != nil {
				return errors.New("item not found")
			}

			if item.Stock < reqItem.Qty {
				return errors.New("stock not sufficient")
			}

			subTotal := float64(reqItem.Qty) * item.Price
			grandTotal += subTotal

			detail := models.PurchasingDetail{
				PurchasingID: purchasing.ID,
				ItemID:       item.ID,
				Qty:          reqItem.Qty,
				SubTotal:     subTotal,
			}

			if err := tx.Create(&detail).Error; err != nil {
				return err
			}

			item.Stock += reqItem.Qty
			if err := tx.Save(&item).Error; err != nil {
				return err
			}

			itemsResp = append(itemsResp, dto.PurchasingDetailResponse{
				ItemID:   item.ID,
				ItemName: item.Name,
				Qty:      reqItem.Qty,
				Price:    item.Price,
				SubTotal: subTotal,
			})

			webhookItems = append(webhookItems, dto.PurchaseWebhookItem{
				ItemID:   item.ID,
				ItemName: item.Name,
				Qty:      reqItem.Qty,
				SubTotal: subTotal,
			})
		}

		if err := tx.Model(&purchasing).
			Update("grand_total", grandTotal).Error; err != nil {
			return err
		}

		response = &dto.PurchasingResponse{
			ID:         purchasing.ID,
			Date:       purchasing.Date,
			GrandTotal: grandTotal,
			Supplier: dto.SupplierMiniResponse{
				ID:   supplier.ID,
				Name: supplier.Name,
			},
			User: dto.UserMiniResponse{
				ID:       user.ID,
				Username: user.Username,
			},
			Details: itemsResp,
		}

		webhookPayload = dto.PurchaseWebhookPayload{
			PurchaseID: purchasing.ID,
			Supplier: dto.SupplierMiniResponse{
				ID:   supplier.ID,
				Name: supplier.Name,
			},
			User: dto.UserMiniResponse{
				ID:       user.ID,
				Username: user.Username,
			},
			GrandTotal: grandTotal,
			Items:      webhookItems,
			CreatedAt:  purchasing.Date,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	go s.webhookClient.SendPurchaseCreated(webhookPayload)

	return response, nil
}

func (s *purchasingService) FindAll() ([]dto.PurchasingResponse, error) {
	purchasings, err := s.purchasingRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var response []dto.PurchasingResponse

	for _, p := range purchasings {
		var details []dto.PurchasingDetailResponse

		for _, d := range p.Details {
			details = append(details, dto.PurchasingDetailResponse{
				ItemID: d.ItemID,
				ItemName: d.Item.Name,
				Qty: d.Qty,
				Price: d.Item.Price,
				SubTotal: d.SubTotal,
			})
		}

		response = append(response, dto.PurchasingResponse{
			ID: p.ID,
			Date: p.Date,
			GrandTotal: p.GrandTotal,
			Supplier: dto.SupplierMiniResponse{
				ID: p.Supplier.ID,
				Name: p.Supplier.Name,
			},
			User: dto.UserMiniResponse{
				ID: p.User.ID,
				Username: p.User.Username,
			},
			Details: details,
		})
	}
	
	return response, nil
}