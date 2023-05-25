package middlewares

import (
	"final-project-3/database"
	"final-project-3/models"
	"net/http"
	"strconv"

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
