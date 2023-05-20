package middlewares

import (
	"final-project-3/helpers"
	"net/http"

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
		// nyimpen claim untuk diambil di endpoint berikutnya
		c.Set("userData", verifyToken)
		c.Next() // lanjut ke endpoint berikutnya, yakni ke controller
	}
}
