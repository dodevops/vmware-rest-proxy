package endpoints

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
	"vmware-rest-proxy/internal"
)

// The Endpoint interface is used to have a common API for all available endpoints
type Endpoint interface {
	// Register needs to be available to register the endpoint with the engine
	Register(engine *gin.Engine)
}

// RequestData holds username and password from a request
type RequestData struct {
	// The Username issued in the Authorization header
	Username string
	// The Password issued in the Authorization header
	Password string
}

// HandleRequest manages incoming requests and extracts authorization data and optionally fails them if no authorization
// is present.
func HandleRequest(context *gin.Context) (RequestData, bool) {
	if u, p, err := getAuthData(context.GetHeader("Authorization")); err != nil {
		var missingAuthenticationHeaderError internal.MissingAuthorizationHeaderError
		if errors.As(err, &missingAuthenticationHeaderError) {
			context.AbortWithStatusJSON(401, gin.H{
				"error": "Missing authentication header",
			})
		} else {
			context.AbortWithStatusJSON(400, gin.H{
				"error": fmt.Sprintf("Error getting Authorization header: %s", err),
			})
		}
		return RequestData{}, false
	} else {
		return RequestData{
			Username: u,
			Password: p,
		}, true
	}
}

// getAuthData extracts the username and password from an authorization header
func getAuthData(authHeader string) (string, string, error) {
	if re, err := regexp.Compile("^Basic (.+)$"); err != nil {
		return "", "", err
	} else if re.MatchString(authHeader) {
		bS := re.FindStringSubmatch(authHeader)[1]
		if dS, err := base64.StdEncoding.DecodeString(bS); err != nil {
			return "", "", err
		} else {
			uP := strings.Split(string(dS), ":")
			return uP[0], uP[1], nil
		}
	} else {
		return "", "", internal.MissingAuthorizationHeaderError{}
	}
}
