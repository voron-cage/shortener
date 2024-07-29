package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"net/http"
)

func (s *Service) ContextUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHead := AuthorizationHeaderRequest{}
		c.BindHeader(&authHead)
		token, err := jwt.Parse(authHead.Authorization, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				c.Status(http.StatusUnauthorized)
				return "", nil
			}
			return []byte(s.secretKey), nil
		})
		if err == nil {
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {
				email := claims["email"].(string)
				user, err := s.tntStorage.GetUser(email)
				if err != nil {
					s.logger.Warn("failed to fetch user from tarantool", zap.Error(err))
				}
				c.Set("user", user)
			}
		}
		c.Next()
	}
}
