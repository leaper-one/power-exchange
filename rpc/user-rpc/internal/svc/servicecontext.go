package svc

import (
	core "github.com/leaper-one/power-exchange/core/user"
	"github.com/leaper-one/power-exchange/rpc/user-rpc/internal/config"
	store "github.com/leaper-one/power-exchange/store/user"
	"github.com/leaper-one/power-exchange/util"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	// DbEngin *gorm.DB
	DbEngin core.UserStore
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DataSourceName), &gorm.Config{})
	if err != nil {
		logx.Error(err)
	}
	//自动同步更新表结构
	db.AutoMigrate(&core.User{})

	return &ServiceContext{
		Config:  c,
		DbEngin: store.NewUserStore(&util.DB{Write: db, Read: db}),
	}
}
