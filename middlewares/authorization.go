package middlewares

// import (
// 	"final-project-3/database"
// 	"final-project-3/models"
// 	"net/http"
// 	"strconv"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// )

// func ProductAuthorization() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		db := database.GetDB()
// 		// misal ada user yg mau akses product berdasarkan productId
// 		// nya melalui param. Param = path variable
// 		productId, err := strconv.Atoi(c.Param("productId"))
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 				"error":   "Bad Request",
// 				"message": "Invalid parameter (gada productId)",
// 			})
// 			return
// 		}
// 		userData := c.MustGet("userData").(jwt.MapClaims)
// 		userId := uint(userData["id"].(float64))
// 		Product := models.Product{}

// 		// cek product yg dicari berdasarkan product id nya ada atau engga
// 		err = db.Select("user_id").First(&Product, uint(productId)).Error

// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 				"error":   "Not Found",
// 				"message": "Product not found",
// 			})
// 			return
// 		}

// 		// product nya ada, tp cek dulu userId nya sama dengan
// 		// userId si user yg lg login ngga?
// 		if Product.UserId != userId {
// 			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
// 				"error":   "Unauthorized",
// 				"message": "You are not allowed to access this product",
// 			})
// 			return
// 		}

// 		c.Next()
// 	}
// }
