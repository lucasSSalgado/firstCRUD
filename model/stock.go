package model

type Stock struct {
	ID     uint64 `gorm:"primary_key,autoIncrement"`
	Ticker string
	Price  float64
}
