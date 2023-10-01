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

		stringToken := token // Corrected line

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

		// Your JWT validation logic here, e.g., check if the token is valid

		// If the token is valid, you can proceed to the next middleware/handler
		c.Next()
	}
}
