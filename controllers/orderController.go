package controllers

import (
	"Demo/database"
	"Demo/models"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderCreatedForm struct {
	CustomerName string `json:"customer_name"`
	Quantity     int    `json:"quantity"`
	ProductName  string `json:"product_name"`
}

// @Tags   OrderAPI
// @Router /orders [post]
// @Param  Request body OrderCreatedForm true "Order Created data"
// @Param Authorization header string true "JWT token" default(Bearer)
func OrderCreate(c *gin.Context) {
	order := models.Order{}
	product := models.Product{}

	var createdForm OrderCreatedForm

	if err := c.ShouldBindJSON(&createdForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not Found createdForm data",
		})
		return
	}

	//check product
	if err := database.DB.Where("product_name=?", createdForm.ProductName).First(&product).Error; err != nil {
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

	//update product Quantity
	updatedQuantity := product.Quantity - createdForm.Quantity
	if updatedQuantity < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not enough products in stock",
		})
		return
	}

	if err := database.DB.Model(&product).Where("id=?", product.Id).Update("quantity", updatedQuantity).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update product quantity",
		})
		return
	}

	order = models.Order{
		CustomerName: createdForm.CustomerName,
		Price:        product.Price * createdForm.Quantity,
		Quantity:     createdForm.Quantity,
		CreateDate:   time.Now(),
		ProductId:    product.Id,
		Product:      product,
	}

	result := database.DB.Save(&order)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Database operation failed",
		})
		return

	}

	c.JSON(http.StatusOK, order)
}

// @Tags   OrderAPI
// @Router /orders [get]
// @Param Authorization header string true "JWT token" default(Bearer)
func AllOrders(c *gin.Context) {
	// 启用 GORM 的详细日志记录模式
	database.DB = database.DB.Debug()

	var order []models.Order

	if err := database.DB.Omit("Product").Find(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to find orders",
		})
		return
	}

	orderResponses := []map[string]interface{}{}

	for _, order := range order {
		orderResponse := map[string]interface{}{
			"Id":           order.Id,
			"CustomerName": order.CustomerName,
			"Quantity":     order.Quantity,
			"Price":        order.Price,
			"CreateDate":   order.CreateDate,
		}
		orderResponses = append(orderResponses, orderResponse)
	}
	c.JSON(http.StatusOK, orderResponses)
}

type OrderPutForm struct {
	CustomerName string `json:"customer_name"`
	Quantity     int    `json:"quantity"`
}

// @Tags   OrderAPI
// @Router /orders/{id} [put]
// @Param  id path int true "Order ID"
// @Param  Request body OrderPutForm true "Order Update data"
// @Param Authorization header string true "JWT token" default(Bearer)
func OrderPut(c *gin.Context) {
	orderId := c.Param("id")

	order := models.Order{}
	product := models.Product{}

	var PutForm OrderPutForm

	if err := c.ShouldBindJSON(&PutForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not Found PutForm data",
		})
		return
	}

	//Error func
	handleDBError := func(err error, message string) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": message})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Database operation failed"})
		}
	}

	//查詢 order & product
	if err := database.DB.Where("id=?", orderId).First(&order).Error; err != nil {
		handleDBError(err, "Order not found")
		return
	}
	if err := database.DB.Where("id=?", order.ProductId).First(&product).Error; err != nil {
		handleDBError(err, "Product not found")
		return
	}

	// //計算產品數量，如果 更新的數量<原訂單數量，則把數量加回product
	// quantity := PutForm.Quantity - order.Quantity

	// if quantity > 0 {
	// 	// Update product's quantity by adding the returned quantity
	// 	updatedQuantity := product.Quantity + quantity

	// 	if err := database.DB.Model(&product).Where("id = ?", product.Id).Update("quantity", updatedQuantity).Error; err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"message": "Failed to update product quantity",
	// 		})
	// 		return
	// 	}
	// } else if quantity < 0 {

	// 	updatedQuantity := product.Quantity + (-quantity)

	// 	if err := database.DB.Model(&product).Where("id = ?", product.Id).Update("quantity", updatedQuantity).Error; err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"message": "Failed to update product quantity",
	// 		})
	// 		return
	// 	}
	// }

	// ======================================================
	// 				|	|
	// 				|	|    簡化 計算產品數量 代碼
	// 				|	|
	// 				|	|
	// 			   \	 /
	// 			    \	/
	// 				 \ /
	// ======================================================

	quantity := PutForm.Quantity - order.Quantity
	fmt.Printf("Quantity difference: %d\n", quantity)

	updatedQuantity := product.Quantity + (-quantity)
	fmt.Printf("Updated product quantity: %d\n", updatedQuantity)

	if err := database.DB.Model(&product).Where("id = ?", product.Id).Update("quantity", updatedQuantity).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to update product quantity"})
		return
	}

	updates := map[string]interface{}{
		"CustomerName": PutForm.CustomerName,
		"Quantity":     PutForm.Quantity,
		"Price":        PutForm.Quantity * product.Price,
	}

	if err := database.DB.Where("id=?", order.ProductId).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to reload product information",
		})
		return
	}

	order.Product = product

	result := database.DB.Model(&order).Updates(updates)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Database operation failed",
		})
		return
	}

	c.JSON(http.StatusOK, order)
}

// @Tags   OrderAPI
// @Router /orders/{id} [delete]
// @Param  id path int true "Order ID"
// @Param Authorization header string true "JWT token" default(Bearer)
func OrderDelete(c *gin.Context) {
	orderId := c.Param("id")

	order := models.Order{}

	if err := database.DB.Where("id=?", orderId).First(&order).Error; err != nil {
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

	product := models.Product{}
	if err := database.DB.Where("id=?", order.ProductId).First(&product).Error; err != nil {
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

	updatedQuantity := product.Quantity + order.Quantity
	fmt.Printf("Updated product quantity: %d\n", updatedQuantity)

	if err := database.DB.Model(&product).Where("id = ?", product.Id).Update("quantity", updatedQuantity).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to update product quantity"})
		return
	}

	result := database.DB.Delete(&order)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Database operation failed",
		})
		return
	}

	c.JSON(http.StatusOK, order)
}
