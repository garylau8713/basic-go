package web

import (
	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/service"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	// 和上面比起来，用 ` 看起来就比较清爽
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	bizLogin             = "login"
)

// UserHandler Used for define all user related routes
type UserHandler struct {

	// Pre compile
	emailRegexExp    *regexp.Regexp
	passwordRegexExp *regexp.Regexp
	svc              *service.UserService
}

// Follow the dependency injection pattern

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		emailRegexExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRegexExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		svc:              svc,
	}
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	// Those two should be the ones that strictly followed the RESTful coding style.
	//server.POST("/users", h.SignUp)
	//server.PUT("/users", h.SignUp)
	// This is RESTful style get method to get the user base info.
	//server.GET("/users/:id", h.Profile)
	ug := server.Group("/users")
	// POST /users/signup
	ug.POST("/signup", h.SignUp)
	// POST /users/login
	ug.POST("/login", h.Login)
	// POST /users/edit
	ug.POST("/edit", h.Edit)
	// This endpoint is used to get the user basic profile info.
	// GET /users/profile
	ug.GET("/profile", h.Profile)
}

func (h *UserHandler) SignUp(ctx *gin.Context) {
	// Define a SignUpRequest Struct (internal class)
	type SignUpReq struct {
		// Here is a tag, defined the tag name email in json
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	var req SignUpReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	isEmail, err := h.emailRegexExp.MatchString(req.Email)
	//isEmail, err := regexp.Match(emailRegexPattern, []byte(req.Email))

	if err != nil {
		ctx.String(http.StatusOK, "System Internal Error.")
		return
	}

	if !isEmail {
		ctx.String(http.StatusOK, "illegal email format.")
		return
	}

	isPassword, err := h.passwordRegexExp.MatchString(req.Password)

	if err != nil {
		ctx.String(http.StatusOK, "System Internal Error.")
		return
	}

	if !isPassword {
		ctx.String(http.StatusOK, "illegal password format.")
		return
	}

	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusOK, "Two passwords are not the same.")
		return
	}

	err = h.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.String(http.StatusOK, "Internal System Error.")
		return
	}
	ctx.String(http.StatusOK, "hello, you have signed up successfully.")

}

func (h *UserHandler) Login(ctx *gin.Context) {

}

func (h *UserHandler) Edit(ctx *gin.Context) {

}

func (h *UserHandler) Profile(ctx *gin.Context) {

}
