package microservice

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nfwGytautas/gdev/array"
)

// Utility function for parsing gin request body, return true if parsed, false otherwise
func GinParseRequestBody(c *gin.Context, out any) bool {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, EndpointError{
			Description: "Failed to read request body",
			Error:       err,
		})
		return false
	}

	err = json.Unmarshal(body, &out)
	if err != nil {
		c.JSON(http.StatusBadRequest, EndpointError{
			Description: "Failed to unmarshal request body",
			Error:       err,
		})
		return false
	}

	return true
}

// Get token info from gin context
func GetTokenInfo(c *gin.Context) (TokenInfo, error) {
	tokenInfo, exists := c.Get("TokenInfo")
	if !exists {
		// Error
		return TokenInfo{}, errors.New("token info not set, make sure authentication middleware is used")
	}

	return tokenInfo.(TokenInfo), nil
}

/*
Middleware for authenticating

Usage:
r.Use(microservice.AuthenticationMiddleware())
*/
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := parseTokenOrFail(c)
		if err != nil {
			return
		}

		c.Next()
	}
}

/*
Middleware for authorization

Usage:
r.Use(microservice.AuthorizationMiddleware([]string{"role"}))
*/
func AuthorizationMiddleware(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		info, err := parseTokenOrFail(c)
		if err != nil {
			return
		}

		if !array.IsElementInArray(roles, info.Role) {
			c.JSON(http.StatusUnauthorized, EndpointError{
				Description: "Access denied, check with your IT department",
				Error:       nil,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Parse token or abort request
func parseTokenOrFail(c *gin.Context) (TokenInfo, error) {
	tokenString := c.Query("token")

	if tokenString == "" {
		// Token empty check if it is inside Authorization header
		tokenString = c.Request.Header.Get("Authorization")

		// Since this is bearer token we need to parse the token out
		if len(strings.Split(tokenString, " ")) == 2 {
			tokenString = strings.Split(tokenString, " ")[1]
		} else {
			c.JSON(http.StatusBadRequest, EndpointError{
				Description: "Token not specified",
				Error:       nil,
			})
			c.Abort()
			return TokenInfo{}, errors.New("token not specified")
		}
	}

	info, err := ParseToken(tokenString)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, EndpointError{
			Description: "Failed to parse token",
			Error:       err,
		})
		c.Abort()
		return TokenInfo{}, err
	}

	if !info.Valid {
		c.JSON(http.StatusUnauthorized, EndpointError{
			Description: "Access denied, token invalid",
			Error:       nil,
		})
		c.Abort()
		return TokenInfo{}, errors.New("token invalid")
	}

	c.Set("TokenInfo", info)
	return info, nil
}
