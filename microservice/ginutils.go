package microservice

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
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
