package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

const secretKey = ""

// Function to verify JWT tokens
func authenticateMiddleware(c *gin.Context) {
	// Retrieve the token from the cookie
	tokenString, err := c.Cookie("token")
	if err != nil {
		fmt.Println("Token missing in cookie")
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	// Verify the token
	token, err := verifyToken(tokenString)
	if err != nil {
		fmt.Printf("Token verification failed: %v\\n", err)
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	// Print information about the verified token
	fmt.Printf("Token verified successfully. Claims: %+v\\n", token.Claims)

	// Continue with the next middleware or route handler
	c.Next()
}

// Function to verify JWT tokens
func verifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return the verified token
	return token, nil
}
