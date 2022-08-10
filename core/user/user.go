package user_core

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		// 用户 UUID 36 位
		UUID string `gorm:"type:varchar(36);unique_index"`
		// 用户密码 Hash 后的密码
		Password string `gorm:"type:varchar(100)"`
		// 用户名
		UserName sql.NullString `gorm:"type:varchar(10);unique_index"`
		// 用户邮箱
		Email sql.NullString `gorm:"type:varchar(50);unique_index"`
		// 用户手机号
		Phone sql.NullString `gorm:"type:varchar(20);unique_index"`
		// 用户状态
		Status bool `gorm:"type:bool;default:true"`
		// 用户类型
		Type sql.NullInt64 `gorm:"type:int;default:0"`
		// 用户头像
		AvatarUrl string `gorm:"type:varchar(100)"`
	}

	UserStore interface {
		// 存储一条用户记录
		Save(ctx context.Context, user *User) error
		// 根据 UUID 查找用户
		FindByUUID(ctx context.Context, uuid string) (*User, error)
		// 查找一个用户，根据用户名
		FindByUserName(ctx context.Context, username string) (*User, error)
		// 查找一个用户，根据邮箱
		FindByEmail(ctx context.Context, email string) (*User, error)
		// 查找一个用户，根据手机号
		FindByPhone(ctx context.Context, phone string) (*User, error)
		// 根据 UUID 删除用户
		DeleteByUUID(ctx context.Context, uuid string) error
	}
)
