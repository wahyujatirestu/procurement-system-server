package services

import (
	"errors"
	"strings"
	"github.com/go-playground/validator/v10"
	"github.com/wahyujatirestu/simple-procurement-system/dto"
	"github.com/wahyujatirestu/simple-procurement-system/models"
	"github.com/wahyujatirestu/simple-procurement-system/repositories"
)

var validate = validator.New()

type SupplierService interface {
	Create(req dto.CreateSupplierRequest) (*dto.SupplierResponse, error)
	FindAll() ([]dto.SupplierResponse, error)
}


type supplierService struct {
	repo repositories.SupplierRepository
}

func NewSupplierService(repo repositories.SupplierRepository) SupplierService {
	return &supplierService{repo}
}

func (s *supplierService) Create(req dto.CreateSupplierRequest) (*dto.SupplierResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, errors.New("supplier name is required")
	}

	email := strings.TrimSpace(req.Email)
	if email == "" {
		return nil, errors.New("email is required")
	}

	address := strings.TrimSpace(req.Address)
	if address == "" {
		return nil, errors.New("address is required")
	}

	supplier := models.Supplier{
		Name:    name,
		Email:   req.Email,
		Address: req.Address,
	}

	if err := s.repo.Create(&supplier); err != nil {
		return nil, err
	}

	return &dto.SupplierResponse{
		ID:      supplier.ID,
		Name:    supplier.Name,
		Email:   supplier.Email,
		Address: supplier.Address,
	}, nil
}


func (s *supplierService) FindAll() ([]dto.SupplierResponse, error) {
	suppliers, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.SupplierResponse, 0)
	for _, supplier := range suppliers {
		responses = append(responses, dto.SupplierResponse{
			ID:      supplier.ID,
			Name:    supplier.Name,
			Email:   supplier.Email,
			Address: supplier.Address,
		})
	}

	return responses, nil
}

