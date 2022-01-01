package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Header struct {
	Token string `header:"auth_token"`
}

type Claim struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

func IsAuthenticated(c *gin.Context) {

	// get token from header
	header := Header{}
	if err := c.ShouldBindHeader(&header); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
	}

	// check if token present
	if len(header.Token) < 1 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no token sent"})
		return
	}

	// extract token from jwt
	token, err := jwt.ParseWithClaims(header.Token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There is an error")
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "not authorised"})
		return
	}

	// if the token is valid
	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		c.Set("user_id", claims.UserId)
	}

	c.Next()

}
