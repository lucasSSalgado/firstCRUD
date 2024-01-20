package service

import (
	"github.com/lucasSSalgado/firstCRUD.git/model"
	"gorm.io/gorm"
)

type StockService struct {
	db *gorm.DB
}

func NewStockService(db *gorm.DB) *StockService {
	return &StockService{
		db: db,
	}
}

func (s *StockService) FindByID(id uint64) (model.Stock, error) {
	stock := new(model.Stock)
	resp := s.db.First(&stock, id)
	if resp.Error != nil {
		return model.Stock{}, resp.Error
	}

	return *stock, nil
}

func (s *StockService) SaveStock(stock model.Stock) (uint64, error) {
	result := s.db.Create(&stock)
	if result.Error != nil {
		return 0, result.Error
	}

	return stock.ID, nil
}