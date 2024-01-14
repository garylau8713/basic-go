package middleware

import (
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginMiddlewareBuilder struct {
}

func (m *LoginMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if path == "/users/signup" ||
			path == "/users/login" {
			// Don't need to check for those two paths.
			return
		}
		sess := sessions.Default(ctx)
		userId := sess.Get("userId")
		if sess.Get("userId") == nil {
			// Abort
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		now := time.Now()
		const updateTimeKey = "update_time"
		// Try to get last refresh time
		val := sess.Get(updateTimeKey)
		lastUpdateTime, ok := val.(time.Time)
		if val == nil || !ok || (now.Sub(lastUpdateTime) > time.Second*10) {
			sess.Set(updateTimeKey, now)
			sess.Set("userId", userId)
			err := sess.Save()
			if err != nil {
				// Log
				fmt.Println(err)
			}
		}
	}
}
