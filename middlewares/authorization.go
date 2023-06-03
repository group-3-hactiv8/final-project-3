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
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		initialUser := &models.User{}
		initialUser.ID = userId

		userRepo := user_pg.NewUserPG(db)
		userRepo.GetUserByID(initialUser)
		// abis di Get, objek initialUser akan terupdate,
		// smua attribute nya akan terisi.

		// user nya fix ada karna udh di cek di authentication,
		// tp cek dulu role nya "admin" bukan?
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
