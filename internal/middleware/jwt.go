package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pndwrzk/taskhub-service/config"
	"github.com/pndwrzk/taskhub-service/internal/common/response"
	"github.com/pndwrzk/taskhub-service/internal/common/utils"
	errConst "github.com/pndwrzk/taskhub-service/internal/constants/error"
)

// JWTAuth middleware wajib login
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response.HttpResponse(response.ParamHTTPResp{
				Code:  http.StatusUnauthorized,
				Gin:   ctx,
				Error: errConst.ErrTokenNotFound,
			})
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.HttpResponse(response.ParamHTTPResp{
				Code:  http.StatusUnauthorized,
				Gin:   ctx,
				Error: errConst.ErrTokenFormat,
			})
			ctx.Abort()
			return
		}

		tokenStr := parts[1]

		claims := &utils.Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.App.JWTAccessSecret), nil
		})

		if err != nil || !token.Valid {
			if err == jwt.ErrTokenExpired {
				response.HttpResponse(response.ParamHTTPResp{
					Code:  http.StatusUnauthorized,
					Gin:   ctx,
					Error: errConst.ErrTokenExpired,
				})
				ctx.Abort()
				return
			}

			response.HttpResponse(response.ParamHTTPResp{
				Code:  http.StatusUnauthorized,
				Gin:   ctx,
				Error: errConst.ErrInvalidToken,
			})
			ctx.Abort()
			return
		}

		// Set user_id ke context
		ctx.Set("user_id", claims.UserID)
		ctx.Next()
	}
}
