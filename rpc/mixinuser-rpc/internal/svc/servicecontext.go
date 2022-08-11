package svc

import (
	core "github.com/leaper-one/power-exchange/core/mixin_user"
	"github.com/leaper-one/power-exchange/rpc/mixinuser-rpc/internal/config"
	store "github.com/leaper-one/power-exchange/store/mixin_user"
	"github.com/leaper-one/power-exchange/util"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DbEngin core.MixinUserStore
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DataSourceName), &gorm.Config{})
	if err != nil {
		logx.Error(err)
	}
	//自动同步更新表结构
	db.AutoMigrate(&core.MixinUser{})

	return &ServiceContext{
		Config:  c,
		DbEngin: store.NewMixinUserStore(&util.DB{Write: db, Read: db}),
	}
}
