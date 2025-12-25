// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"fim_server/core"
	"fim_server/fim_auth/auth_api/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	//mysqlDb.AutoMigrate(&auth_models.UserModel{})
	return &ServiceContext{
		Config: c,
		DB:     mysqlDb,
	}
}
