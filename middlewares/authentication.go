package middlewares

import (
	"final-project-3/database"
	"final-project-3/helpers"
	"final-project-3/models"
	"final-project-3/repositories/user_repository/user_pg"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken // dapet MapClaims yg isinya id dan email

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}

		data := verifyToken.(jwt.MapClaims)

		id := uint(data["id"].(float64))
		if _, isExist := data["email"]; !isExist {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": "invalid token",
			})
			return
		}
		email := data["email"].(string)

		initialUser := &models.User{}
		initialUser.ID = id

		db := database.GetPostgresInstance()
		userRepo := user_pg.NewUserPG(db)

		// does a user exist with this id?
		_, errNotFound := userRepo.GetUserByID(id)
		if errNotFound != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": "invalid token",
			})
			return
		}

		// if exist, is the email is the same with the one from the token?
		user, err := userRepo.GetUserByID(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": "invalid token",
			})
			return
		}

		if user.Email != email {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": "invalid token",
			})
			return
		}

		// nyimpen claim untuk diambil di endpoint berikutnya
		c.Set("userData", verifyToken)
		c.Next() // lanjut ke endpoint berikutnya, yakni ke controller
	}
}