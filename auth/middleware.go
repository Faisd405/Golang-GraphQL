package auth

import (
	"context"
	"graphql-template/jwt"
	"graphql-template/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.Next()
			return
		}

		tokenStr := header
		username, err := jwt.ParseToken(tokenStr)
		if err != nil {
			c.Next()
			return
		}

		user := models.User{Username: username}
		id, err := models.GetUserIdByUsername(username)
		if err != nil {
			c.Next()
			return
		}
		user.ID = strconv.Itoa(id)
		// put it in context
		ctx := context.WithValue(c.Request.Context(), userCtxKey, &user)

		// and call the next with our new context
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func ForContext(ctx context.Context) *models.User {
	raw, _ := ctx.Value(userCtxKey).(*models.User)
	return raw
}
