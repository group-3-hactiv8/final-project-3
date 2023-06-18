package middlewares

import (
	"final-project-3/database"
	"final-project-3/models"
	"final-project-3/repositories/user_repository/user_pg"
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

		// cek Task yg dicari berdasarkan Task id nya ada atau engga
		err = db.Select("user_id").First(&task, uint(taskId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Task not found",
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
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))

		userRepo := user_pg.NewUserPG(db)
		initialUser, err := userRepo.GetUserByID(userId)
		if err != nil {
			// Handle error ketika mendapatkan user
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal Server Error",
				"message": "Failed to retrieve user data",
			})
			return
		}

		// Cek role user, jika bukan "admin", berikan response Forbidden
		if initialUser.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this category feature",
			})
			return
		}

		c.Next()
	}
}
