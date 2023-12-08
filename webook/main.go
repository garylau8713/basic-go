package main

import (
	"basic-go/webook/internal/web"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func main() {

	server := gin.Default()

	// TODO: 跨域问题。 没太懂这块, 老师说回头会讲
	server.Use(cors.New(cors.Config{
		//AllowOrigins: []string{"http://localhost:3000"},
		//AllowAllOrigins: true,
		AllowCredentials: true,
		// Added the header "authorization" because the browser log out the error message said the server side
		// can not take "authorization" header. For the webook-fe, it does pass the "authorization" header.
		AllowHeaders: []string{"Content-Type", "authorization"},
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "your_company.com")
		},
		MaxAge: 12 * time.Hour,
	}), func(ctx *gin.Context) {
		fmt.Println("This is middleware")
	})

	//hdl := &web.UserHandler{}
	hdl := web.NewUserHandler()
	hdl.RegisterRoutes(server)

	server.Run(":8080")
}
