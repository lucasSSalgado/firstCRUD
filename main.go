package main

import (
	"github.com/lucasSSalgado/firstCRUD.git/controller"
	"github.com/lucasSSalgado/firstCRUD.git/infra"
	"github.com/lucasSSalgado/firstCRUD.git/service"
)

func main() {
	db := infra.CreateConnection()
	stockService := service.NewStockService(db)
	stockController := controller.NewStockController(stockService)

	stockController.InitRoutes()
}