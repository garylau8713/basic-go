package main

import (
	"basic-go/webook/internal/repository"
	"basic-go/webook/internal/repository/dao"
	"basic-go/webook/internal/service"
	"basic-go/webook/internal/web"
	"basic-go/webook/internal/web/middleware"
	"basic-go/webook/pkg/ginx/middleware/ratelimit"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

func main() {

	// 3-7：因为学习部署简单的web app，所以先把复杂的功能注释掉，之后再恢复
	//db := initDB()
	//server := initWebServer()
	//initUserHdl(db, server)
	server := gin.Default()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, it started.")
	})
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
		// 允许前端访问的你的后端响应中带的头部
		ExposeHeaders: []string{"x-jwt-token"},

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

	// 限流尽可能放在前面
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	// 这块已经用于接入限流插件
	server.Use(ratelimit.NewBuilder(redisClient, time.Second, 1).Build())
	useJWT(server)
	//useSession(server)

	return server
}

func useJWT(server *gin.Engine) {
	login := &middleware.LoginJWTMiddlewareBuilder{}
	server.Use(login.CheckLogin())

}

func useSession(server *gin.Engine) {
	//Init Session
	//  存储数据的，也就是userId存哪里
	//// 直接存cookie
	store := cookie.NewStore([]byte("secret"))
	// 不存在cookie里面，现在用别的存
	// 基于内存的实现，single instance利用memcache来存session
	//store, err := redis.NewStore(16, "tcp", "localhost:6379", "",
	//	[]byte("GONdbXwcBYhLJjYq7EX2cyKkzCR7XiC5"),
	//	[]byte("l564gjjcmHIksYTGmDSliSjDFj7an4mk"))
	//if err != nil {
	//	panic(err)
	//}
	//store := memstore.NewStore([]byte("GONdbXwcBYhLJjYq7EX2cyKkzCR7XiC5"),
	//	[]byte("l564gjjcmHIksYTGmDSliSjDFj7an4mk"))
	login := &middleware.LoginMiddlewareBuilder{}
	// First one to init sessions
	// Second one to use for checking login
	server.Use(sessions.Sessions("ssid", store),
		login.CheckLogin())
}
