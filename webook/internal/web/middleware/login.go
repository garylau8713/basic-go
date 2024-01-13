package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginMiddlewareBuilder struct {
}

func (m *LoginMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if path == "/users/signup" ||
			path == "/users/login" {
			// Don't need to check for those two paths.
			return
		}
		sess := sessions.Default(ctx)
		if sess.Get("userId") == nil {
			// Abort
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
