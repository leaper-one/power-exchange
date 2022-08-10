package svc

import (
	core "github.com/leaper-one/power-exchange/core/user"
	"github.com/leaper-one/power-exchange/rpc/user-rpc/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DataSourceName), &gorm.Config{
		// NamingStrategy: schema.NamingStrategy{
		// TablePrefix:   "t_", // 表名前缀，`User` 的表名应该是 `t_users`
		// SingularTable: true,    // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		// },
	})
	if err != nil {
		// panic(err)
		logx.Error(err)
	}
	//自动同步更新表结构
	db.AutoMigrate(&core.User{})

	return &ServiceContext{
		Config:  c,
		DbEngin: db,
	}
}
