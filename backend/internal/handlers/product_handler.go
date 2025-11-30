package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"product-inventory/internal/database"
	"product-inventory/internal/models"
)

type CreateProductRequest struct {
	Product struct {
		ProductID    string `json:"productID"`
		ProductCode  string `json:"productCode"`
		ProductName  string `json:"productName"`
		ProductImage string `json:"productImage"`
		CreatedUser  string `json:"createdUser"`
		IsFavourite  bool   `json:"isFavourite"`
		Active       bool   `json:"active"`
		HSNCode      string `json:"hsnCode"`
	} `json:"product"`

	Variants []VariantInput `json:"variants"`
}

type VariantInput struct {
	Name    string   `json:"name"`
	Options []string `json:"options"`
}

func CreateProduct(c *gin.Context) {
	// ---------- 1. Bind into DTO (safe for strings) ----------
	var req CreateProductRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ---------- 2. Convert productID (string → int64) ----------
	productID, err := strconv.ParseInt(req.Product.ProductID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid productID"})
		return
	}

	// ---------- 3. Convert createdUser (string → uuid.UUID) ----------
	var createdUser uuid.UUID
	if req.Product.CreatedUser != "" { // allow empty
		parsed, err := uuid.Parse(req.Product.CreatedUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid createdUser UUID"})
			return
		}
		createdUser = parsed
	}

	// ---------- 4. Build the actual Product model ----------
	product := models.Product{
		ID:           uuid.New(),
		ProductID:    productID,
		ProductCode:  req.Product.ProductCode,
		ProductName:  req.Product.ProductName,
		ProductImage: req.Product.ProductImage,
		CreatedUser:  createdUser,
		IsFavourite:  req.Product.IsFavourite,
		Active:       req.Product.Active,
		HSNCode:      req.Product.HSNCode,
	}

	// ---------- 5. Save Product ----------
	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ---------- 6. Build variants + options ----------
	var allOptions [][]string

	for _, v := range req.Variants {
		variant := models.Variant{
			ID:        uuid.New(),
			ProductID: product.ID,
			Name:      v.Name,
		}
		database.DB.Create(&variant)

		var optionValues []string

		for _, opt := range v.Options {
			optRow := models.VariantOption{
				ID:        uuid.New(),
				VariantID: variant.ID,
				Value:     opt,
			}
			database.DB.Create(&optRow)

			optionValues = append(optionValues, opt)
		}

		allOptions = append(allOptions, optionValues)
	}

	// ---------- 7. Auto-generate subvariants ----------
	combinations := cartesianProduct(allOptions)

	for _, opts := range combinations {
		sv := models.SubVariant{
			ID:        uuid.New(),
			ProductID: product.ID,
			OptionIDs: opts,
			SKU:       generateSKU(product.ProductCode, opts),
			Stock:     decimal.NewFromInt(0),
		}
		database.DB.Create(&sv)
	}

	// ---------- 8. Success ----------
	c.JSON(http.StatusOK, gin.H{"message": "product created with auto subvariants"})
}

func ListProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	var products []models.Product

	err := database.DB.
		Preload("Variants").
		Preload("Variants.Options").
		Preload("SubVariants").
		Offset(offset).
		Limit(limit).
		Find(&products). // ✅ Find must be LAST
		Error

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data":  products,
		"page":  page,
		"limit": limit,
	})
}

func cartesianProduct(arr [][]string) [][]string {
	if len(arr) == 0 {
		return [][]string{}
	}

	result := [][]string{{}}

	for _, group := range arr {
		var temp [][]string
		for _, prefix := range result {
			for _, item := range group {
				newCombo := append([]string{}, prefix...)
				newCombo = append(newCombo, item)
				temp = append(temp, newCombo)
			}
		}
		result = temp
	}

	return result
}

func generateSKU(base string, opts []string) string {
	sku := base
	for _, o := range opts {
		sku += "-" + strings.ToUpper(o)
	}
	return sku
}
