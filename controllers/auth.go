package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ix-pay/ixpay/container"
	"github.com/ix-pay/ixpay/models"
	"github.com/ix-pay/ixpay/service"
	"github.com/ix-pay/ixpay/utils"
)

type authController struct {
	ctr container.IContainer
}

func NewAuthController(ctr container.IContainer) *authController {
	return &authController{
		ctr: ctr,
	}
}

// Login
// @Summary 登录
// @Description 登录
// @Tags 基础功能
// @Accept json
// @Produce json
// @Param user body models.LoginCredentials true "用户信息"
// @Success 200 {object} utils.RespData{data=models.TokenUser} "成功响应"
// @Failure 400 {object} utils.RespData{data=string} "失败消息"
// @Router /auth/login [post]
func (con *authController) Login(c *gin.Context) {
	var creds models.LoginCredentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		utils.Error(c, http.StatusBadRequest, fmt.Sprintf("参数异常：%s\n", err.Error()))
		return
	}

	// 验证用户逻辑（示例）
	us := con.ctr.MustGet("authService").(service.AuthService)
	user, err := us.AuthenticateUser(creds.Account, creds.Password)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, fmt.Sprintf("认证失败：%s\n", err.Error()))
		return
	}
	j := con.ctr.GetJwt()
	token, err := j.GenerateJWT(user.ID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, fmt.Sprintf("生成token失败：%s\n", err.Error()))
		return
	}

	utils.Success(c, http.StatusOK, "", &models.TokenUser{
		Token: token,
		User: models.ProfileUser{
			Id:      strconv.FormatInt(user.ID, 10),
			Name:    user.Name,
			Account: user.Account,
		},
	})
}

// Register
// @Summary 注册
// @Description 注册
// @Tags 基础功能
// @Accept json
// @Produce json
// @Param user body models.RegisterUser true "用户信息"
// @Success 200 {object} utils.RespData{data=models.ProfileUser} "成功响应"
// @Failure 400 {object} utils.RespData{data=string} "失败消息"
// @Router /auth/register [post]
func (con *authController) Register(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		utils.Error(c, http.StatusBadRequest, fmt.Sprintf("参数异常：%s\n", err.Error()))
		return
	}

	us := con.ctr.MustGet("authService").(service.AuthService)
	if err := us.Register(c, &newUser); err != nil {
		return
	}

	utils.Success(c, http.StatusOK, "注册成功", &models.ProfileUser{
		Id:      strconv.FormatInt(newUser.ID, 10),
		Name:    newUser.Name,
		Account: newUser.Account,
	})
}

// GetProfile
// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags 基础功能
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} utils.RespData{data=models.ProfileUser} "成功响应"
// @Failure 400 {object} utils.RespData{data=string} "失败消息"
// @Router /auth/profile [get]
func (con *authController) GetProfile(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, "无效用户")
		return
	}

	us := con.ctr.MustGet("authService").(service.AuthService)
	pu, err := us.GetProfile(userID.(int64))

	if err != nil {
		utils.Error(c, http.StatusNotFound, fmt.Sprintf("用户不存在：%s\n", err.Error()))
		return
	}

	utils.Success(c, http.StatusOK, "", &pu)
}
