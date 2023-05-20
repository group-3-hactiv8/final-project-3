package helpers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const secretKey = "secret"

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(secretKey))

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("sign in to proceed")
	headerToken := c.Request.Header.Get("Authorization")
	adaKataBearerDiToken := strings.HasPrefix(headerToken, "Bearer")

	if !adaKataBearerDiToken {
		return nil, errResponse
	}

	// headerToken = Bearer abc.klm.xyz
	stringToken := strings.Split(headerToken, " ")[1]

	// Parse = ngubah token dari string ke jwt.Token.
	// Pake secretKey, makannya pas buat token pas otentikasi
	// harus pake secretKey yg sama.
	token, _ := jwt.Parse(
		stringToken, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errResponse
			}
			return []byte(secretKey), nil
		})

	// Claims tugasnya mencasting jd jwt.MapClaims.
	// Ngecek bisa di cast atau engga, dan apakah tokennya valid.
	// Maksutnya valid disini apaya??? kan blom ngecek ke DB passnya
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	// isi dari MapClaims ini adalah id dan email yg digenerate
	// di awal pas login.
	return token.Claims.(jwt.MapClaims), nil
}
