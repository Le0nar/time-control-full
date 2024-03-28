package gateway

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/le0nar/time-control/internal/util"
)

type GatewayHandler struct {
}

func NewGatewayHandler() *GatewayHandler {
	return &GatewayHandler{}
}

const (
	authorizationHeader = "Authorization"
	validateUrl 		= "http://localhost:8001/auth-service/employee/validate"
)

func (gh *GatewayHandler) IdentityEmployee(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	token, err := util.GetTokenFromHeader(header)
	if err != nil {
		util.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
   
	var jsonStr = []byte(``)
    req, err := http.NewRequest(http.MethodPost, validateUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

    req.Header.Set(authorizationHeader, token)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
		util.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

    body, _ := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		util.NewErrorResponse(c, http.StatusUnauthorized, string(body))
		return
	}
}
