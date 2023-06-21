package microservice

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/nfwGytautas/gdev/jwt"
)

func GetTokenInfo(c *gin.Context) (jwt.TokenInfo, error) {
	tokenInfo, exists := c.Get("TokenInfo")
	if !exists {
		// Error
		return jwt.TokenInfo{}, errors.New("token info not set, make sure authentication middleware is used")
	}

	return tokenInfo.(jwt.TokenInfo), nil
}
