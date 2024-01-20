package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucasSSalgado/firstCRUD.git/model"
	"github.com/lucasSSalgado/firstCRUD.git/service"
)

type StockController struct {
	service *service.StockService
}

func NewStockController(service *service.StockService) *StockController {
	return &StockController{
		service: service,
	}
}

func (c *StockController) InitRoutes() {
	app := gin.Default()
	api := app.Group("/api/stock")

	//api.GET("/:id", )
	api.POST("/", c.saveStock)
	//api.PUT("/:id", )
	//app.DELETE("/:id", )

	app.Run(":3000")
}

func (c *StockController) findByID(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err},
		)
		return
	}

	stock, err := c.service.FindByID(id)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err},
		)
		return
	}

	ctx.JSON(http.StatusOK, stock)
}

func (c *StockController) saveStock(ctx *gin.Context) {
	stock := new(model.Stock)

	if err := ctx.ShouldBindJSON(&stock); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err},
		)
		return
	}

	id, err := c.service.SaveStock(*stock)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	}

	ctx.JSON(http.StatusCreated, id)
}
