package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"product-inventory/internal/database"
	"product-inventory/internal/models"
)

type StockReq struct {
	ProductID    uuid.UUID       `json:"product_id"`
	SubVariantID uuid.UUID       `json:"subvariant_id"`
	Quantity     decimal.Decimal `json:"quantity"`
}

func AddStock(c *gin.Context) {
	var req StockReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ❗ Block negative quantity
	if req.Quantity.LessThan(decimal.NewFromInt(0)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "quantity cannot be negative"})
		return
	}

	tx := database.DB.Begin()

	var sv models.SubVariant
	if err := tx.
		Set("gorm:query_option", "FOR UPDATE").
		First(&sv, "id = ?", req.SubVariantID).Error; err != nil {

		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "subvariant not found"})
		return
	}

	sv.Stock = sv.Stock.Add(req.Quantity)
	if err := tx.Save(&sv).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tr := models.StockTransaction{
		ID:              uuid.New(),
		ProductID:       req.ProductID,
		SubVariantID:    req.SubVariantID,
		Quantity:        req.Quantity,
		TransactionType: "IN",
		TransactionDate: time.Now(),
	}
	tx.Create(&tr)

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "stock added"})
}

func RemoveStock(c *gin.Context) {
	var req StockReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ❗ Block negative quantity
	if req.Quantity.LessThan(decimal.NewFromInt(0)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "quantity cannot be negative"})
		return
	}

	tx := database.DB.Begin()

	var sv models.SubVariant
	if err := tx.
		Set("gorm:query_option", "FOR UPDATE").
		First(&sv, "id = ?", req.SubVariantID).Error; err != nil {

		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "subvariant not found"})
		return
	}

	// ❗ Correct comparison: prevent negative stock
	if req.Quantity.GreaterThan(sv.Stock) {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient stock"})
		return
	}

	sv.Stock = sv.Stock.Sub(req.Quantity)
	if err := tx.Save(&sv).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tr := models.StockTransaction{
		ID:              uuid.New(),
		ProductID:       req.ProductID,
		SubVariantID:    req.SubVariantID,
		Quantity:        req.Quantity,
		TransactionType: "OUT",
		TransactionDate: time.Now(),
	}
	tx.Create(&tr)

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "stock removed"})
}

func StockReport(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")

	if from == "" || to == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "from and to dates are required"})
		return
	}

	fromTime := from + " 00:00:00"
	toTime := to + " 23:59:59"

	var txs []models.StockTransaction

	database.DB.
		Where("transaction_date BETWEEN ? AND ?", fromTime, toTime).
		Order("transaction_date ASC").
		Find(&txs)

	c.JSON(http.StatusOK, txs)
}
