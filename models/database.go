package models

import (
	"github.com/ix-pay/ixpay/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg config.DBConfig) error {
	var err error
	dsn := buildDSN(cfg)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// 自动迁移表结构
	if err := DB.AutoMigrate(&User{}); err != nil {
		return err
	}

	return nil
}

func CloseDB() error {
	if DB != nil {
		sqlDB, _ := DB.DB()
		return sqlDB.Close()
	}
	return nil
}

func buildDSN(cfg config.DBConfig) string {
	return "host=" + cfg.Host +
		" user=" + cfg.User +
		" password=" + cfg.Password +
		" dbname=" + cfg.DBName +
		" port=" + cfg.Port +
		" sslmode=" + cfg.SSLMode
}
