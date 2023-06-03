package middlewares

import (
	"final-project-3/database"
	"final-project-3/models"
	"net/http"
	"strconv"
	"fmt"
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func TaskAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetPostgresInstance()
		// misal ada user yg mau akses task berdasarkan taskId
		// nya melalui param. Param = path variable
		taskId, err := strconv.Atoi(c.Param("taskId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid parameter (gada taskId)",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		task := models.Task{}

		// cek product yg dicari berdasarkan product id nya ada atau engga
		err = db.Select("user_id").First(&task, uint(taskId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Product not found",
			})
			return
		}

		// task nya ada, tp cek dulu userId nya sama dengan
		// userId si user yg lg login ngga?
		if task.UserId != userId {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this task",
			})
			return
		}

		c.Next()
	}
}


func CategoryAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetPostgresInstance()

		categoryId, err := strconv.Atoi(c.Param("categoryId"))
		if err != nil {
			badRequestError := fmt.Errorf("Invalid Category ID: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": badRequestError.Error()})
			return
		}

		// userData, ok := c.Get("userData").(jwt.MapClaims)
		userData, ok := c.MustGet("userData").(jwt.MapClaims)
		if !ok {
			unauthorizedError := errors.New("Unauthorized access")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": unauthorizedError.Error()})
			return
		}
		userID := uint(userData["id"].(float64))
		requestedCategory := &models.Category{}

		err = db.First(requestedCategory, "id = ?", categoryId).Error
		if err != nil {
			notFoundError := fmt.Errorf("Category not found: %v", err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": notFoundError.Error()})
			return
		}

		if requestedCategory.ID != userID {
			unauthorizedError := errors.New("You are not allowed to access this Category")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": unauthorizedError.Error()})
			return
		}

		c.Next()
	}
}

