package gateway

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/le0nar/time-control-read/internal/util"
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
	token, err := getTokenFromHeader(header)
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

func  getTokenFromHeader(header string) (string, error) {
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", errors.New("invalid auth header")
	}

	token := headerParts[1]

	return token, nil
}
