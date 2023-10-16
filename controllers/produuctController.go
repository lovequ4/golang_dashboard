package controllers

import (
	"Demo/database"
	"Demo/models"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductCreatedForm struct {
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
}

// @Tags   ProductAPI
// @Router /products [post]
// @Param  Request body ProductCreatedForm true "Product Created data"
func ProductCreate(c *gin.Context) {
	product := models.Product{}

	var createdForm ProductCreatedForm

	if err := c.ShouldBindJSON(&createdForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not Found createdForm data",
		})
		return
	}

	checkProductName := database.DB.Where("product_name=?", createdForm.ProductName).First(&product)
	if checkProductName.Error != nil {
		if errors.Is(checkProductName.Error, gorm.ErrRecordNotFound) {

			product = models.Product{
				ProductName: createdForm.ProductName,
				Price:       createdForm.Price,
				Quantity:    createdForm.Quantity,
				CreateDate:  time.Now(),
			}

			result := database.DB.Save(&product)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Database operation failed",
				})
				return
			} else {
				c.JSON(http.StatusOK, product)
			}

		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "An error occurred while checking product name",
			})
			return
		}

	} else {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Product name already exists",
		})
		return
	}

}

// @Tags   ProductAPI
// @Router /products [get]
func AllProducts(c *gin.Context) {
	var product []models.Product

	// .Error 用於檢查 Gorm 的查詢操作是否出錯。這是因為 Gorm 的 Find 方法返回一個錯誤對象，
	// 可以通過檢查錯誤對象的值來確定查詢是否成功。當 .Error 不為 nil 時，它表示查詢操作出現了
	// 錯誤。
	if err := database.DB.Find(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to find product",
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

type ProductPutForm struct {
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
}

// @Tags   ProductAPI
// @Router /products/{id} [put]
// @Param  id path int true "Product ID"
// @Param  Request body ProductPutForm true "Product Created data"
func ProductPut(c *gin.Context) {
	productId := c.Param("id")

	product := models.Product{}

	var PutForm ProductPutForm

	if err := c.ShouldBindJSON(&PutForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not Found PutForm data",
		})
		return
	}

	if err := database.DB.Where("id=?", productId).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Product not found",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Database operation failed",
			})
		}
		return
	}

	updates := map[string]interface{}{
		"ProductName": PutForm.ProductName,
		"Quantity":    PutForm.Quantity,
		"Price":       PutForm.Price,
	}

	result := database.DB.Model(&product).Updates(updates)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Database operation failed",
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

// @Tags   ProductAPI
// @Router /products/{id} [delete]
// @Param  id path int true "Product ID"
func ProductDelete(c *gin.Context) {
	productId := c.Param("id")

	product := models.Product{}

	if err := database.DB.Where("id=?", productId).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Product not found",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Database operation failed",
			})
		}
		return
	}

	result := database.DB.Delete(&product)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Database operation failed",
		})
		return
	}

	c.JSON(http.StatusOK, product)
}
