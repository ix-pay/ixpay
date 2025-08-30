package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ix-pay/ixpay/models"
	"github.com/ix-pay/ixpay/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	AuthenticateUser(account, password string) (*models.User, error)
	Register(c *gin.Context, newUser *models.User) error
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) AuthenticateUser(account, password string) (*models.User, error) {
	var user models.User
	log.Println("account=" + account)
	log.Println("password=" + password)
	if err := models.DB.Where("account = ?", account).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("记录不存在")
		} else {
			log.Println("查询出错:", err)
		}
		return nil, err
	}
	// 修正点2：添加密码为空的安全检查
	if user.Password == "" {
		return nil, fmt.Errorf("用户密码字段为空")
	}

	log.Println("CompareHashAndPassword=" + user.Password)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		return &user, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (s *userService) Register(c *gin.Context, newUser *models.User) error {
	// 密码加密
	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, fmt.Sprintf("密码加密失败：%s\n", err.Error()))
		return err
	}
	newUser.Password = string(hashed)

	// 模拟数据库存储
	if err := models.DB.Create(newUser).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, fmt.Sprintf("用户已存在：%s\n", err.Error()))
		return err
	}
	return nil
}
