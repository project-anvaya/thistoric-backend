package jwt

import (
	"net/http"
	"os"
	"strings"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireUser() gin.HandlerFunc {
	jwksURL := os.Getenv("SUPABASE_JWKS_URL")
	jwks, err := keyfunc.NewDefault([]string{jwksURL})
	if err != nil {
		panic("jwks init failed: " + err.Error())
	}

	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if !strings.HasPrefix(h, "Bearer ") { c.AbortWithStatus(http.StatusUnauthorized); return }
		raw := strings.TrimPrefix(h, "Bearer ")
		tok, err := jwt.Parse(raw, jwks.Keyfunc)
		if err != nil || !tok.Valid { c.AbortWithStatus(http.StatusUnauthorized); return }
		if claims, ok := tok.Claims.(jwt.MapClaims); ok {
			if sub, ok := claims["sub"].(string); ok { c.Set("sub", sub) }
			if email, ok := claims["email"].(string); ok { c.Set("email", email) }
		}
		c.Next()
	}
}
