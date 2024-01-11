package main

import (
	"basic-go/webook/internal/repository"
	"basic-go/webook/internal/repository/dao"
	"basic-go/webook/internal/service"
	"basic-go/webook/internal/web"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {

	db := initDB()

	server := initWebServer()

	initUserHdl(db, server)

	server.Run(":8080")
}

func initUserHdl(db *gorm.DB, server *gin.Engine) {
	ud := dao.NewUserDao(db)
	ur := repository.NewUserRepository(ud)
	us := service.NewUserService(ur)

	//hdl := &web.UserHandler{}
	hdl := web.NewUserHandler(us)
	hdl.RegisterRoutes(server)
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))

	if err != nil {
		panic(err)
	}

	err = dao.InitTables(db)

	if err != nil {
		panic(err)
	}
	return db
}

func initWebServer() *gin.Engine {
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
	return server
}
