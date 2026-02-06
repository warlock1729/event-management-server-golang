package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/warlock1729/first-go-project/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authoriztion token not found"})
		return
	}
	parsedToken, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	userID, err := utils.ExtractJWTClaims(parsedToken)
	context.Set("userID",userID)
	context.Next()

}
