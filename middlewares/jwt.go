package middlewares

import (
	"net/http"

	"github.com/davethio/task-5-pbi-btpns-DaveChristianThio/helpers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "Unauthorized"}
				helpers.JSONResponse(c, http.StatusUnauthorized, response)
				c.Abort() // Abort the middleware chain
				return
			}
		}

		stringToken := token

		claims := &helpers.JWTClaim{}

		tokenObj, err := jwt.ParseWithClaims(stringToken, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(helpers.JWT_KEY), nil
		})

		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				switch {
				case ve.Errors&jwt.ValidationErrorSignatureInvalid != 0:
					fallthrough
				case ve.Errors&jwt.ValidationErrorExpired != 0:
					response := map[string]string{"message": "Unauthorized"}
					helpers.JSONResponse(c, http.StatusUnauthorized, response)
					c.Abort() // Abort the middleware chain
					return
				default:
					response := map[string]string{"message": "Unauthorized"}
					helpers.JSONResponse(c, http.StatusUnauthorized, response)
					c.Abort() // Abort the middleware chain
					return
				}
			}
		}

		if !tokenObj.Valid {
			response := map[string]string{"message": "Unauthorized"}
			helpers.JSONResponse(c, http.StatusUnauthorized, response)
			c.Abort() // Abort the middleware chain
			return
		}

		// Access the UserID from the claims
		userID := claims.UserID // Replace "UserID" with the actual claim name

		// Now, you can use the userID in your middleware or pass it to the next handler
		// For example, you can set it in the Gin context for later use
		c.Set("userID", userID)

		// Your JWT validation logic here, e.g., check if the token is valid
		// If the token is valid, you can proceed to the next middleware/handler
		c.Next()
	}
}

