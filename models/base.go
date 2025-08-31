package models

import (
	"time"

	"github.com/ix-pay/ixpay/utils"
	"gorm.io/gorm"
)

// 替换gorm.Model的自定义基模型
type BaseModel struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt int64
	UpdatedAt int64
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

var (
	sf *utils.Snowflake
)

func InitSnowflake(MachineId string) error {
	var err error
	sf, err = utils.SetupSnowflake(MachineId)
	return err
}

// BeforeCreate GORM钩子函数
func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if m.ID == 0 {
		m.ID = sf.Generate()
	}
	m.CreatedAt = time.Now().Unix()
	return nil
}

// BeforeUpdate GORM钩子函数
func (m *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = time.Now().Unix()
	return nil
}
