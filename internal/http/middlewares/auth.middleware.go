package middlewares

import (
	"snipz/internal/errors"
	"snipz/internal/services"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeaderKey = "authorization"

	AuthorizationType = "bearer"

	AuthorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenService services.TokenService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(AuthorizationHeaderKey)

		isEmpty := len(authHeader) == 0
		if isEmpty {
			err := errors.ErrEmptyAuthorizationHeader
			ctx.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		fields := strings.Fields(authHeader)
		isValid := len(fields) == 2
		if !isValid {
			err := errors.ErrInvalidAuthorizationHeader
			ctx.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		authType := strings.ToLower(fields[0])
		if authType != AuthorizationType {
			err := errors.ErrInvalidAuthorizationType
			ctx.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		accessToken := fields[1]
		payload, err := tokenService.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}
		ctx.Set(AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}
