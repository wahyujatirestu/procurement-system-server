package services

import (
	"errors"
	"strings"

	"github.com/wahyujatirestu/simple-procurement-system/dto"
	"github.com/wahyujatirestu/simple-procurement-system/models"
	"github.com/wahyujatirestu/simple-procurement-system/repositories"
)

type ItemService interface {
	Create(req dto.CreateItemRequest) (*dto.ItemResponse, error)
	FindAll() ([]dto.ItemResponse, error)
	FindById(id uint) (*dto.ItemResponse, error)
	Update(id uint, req dto.UpdateItemRequest) (*dto.ItemResponse, error)
	Delete(id uint) error
}


type itemService struct {
	repo repositories.ItemRepository
}

func NewItemService(repo repositories.ItemRepository) ItemService {
	return &itemService{repo}
}

func (s *itemService) Create(req dto.CreateItemRequest) (*dto.ItemResponse, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, errors.New("item name is required")
	}

	if req.Stock < 0 {
		return nil, errors.New("stock must be greater than 0")
	}

	if req.Price < 0 {
		return nil, errors.New("price must be greater than 0")
	}

	item := models.Item{
		Name:  name,
		Stock: req.Stock,
		Price: req.Price,
	}

	if err := s.repo.Create(&item); err != nil {
		return nil, err
	}

	return &dto.ItemResponse{
		ID:    item.ID,
		Name:  item.Name,
		Stock: item.Stock,
		Price: item.Price,
	}, nil
}


func (s *itemService) FindAll() ([]dto.ItemResponse, error) {
	items, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.ItemResponse, 0)
	for _, item := range items {
		responses = append(responses, dto.ItemResponse{
			ID:    item.ID,
			Name:  item.Name,
			Stock: item.Stock,
			Price: item.Price,
		})
	}

	return responses, nil
}


func (s *itemService) FindById(id uint) (*dto.ItemResponse, error) {
	item, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	return &dto.ItemResponse{
		ID:    item.ID,
		Name:  item.Name,
		Stock: item.Stock,
		Price: item.Price,
	}, nil
}


func (s *itemService) Update(id uint, req dto.UpdateItemRequest) (*dto.ItemResponse, error) {
	item, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(req.Name) != "" {
		item.Name = req.Name
	}
	if req.Stock >= 0 {
		item.Stock = req.Stock
	}
	if req.Price >= 0 {
		item.Price = req.Price
	}

	if err := s.repo.Update(item); err != nil {
		return nil, err
	}

	return &dto.ItemResponse{
		ID:    item.ID,
		Name:  item.Name,
		Stock: item.Stock,
		Price: item.Price,
	}, nil
}


func (s *itemService) Delete(id uint) error {
	item, err := s.repo.FindById(id)
	if err != nil {
		return err
	}
	if item == nil {
		return errors.New("item not found")
	}

	return s.repo.Delete(id)
}