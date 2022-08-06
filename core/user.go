package core

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// 用户 UUID 36 位
	UUID string `gorm:"type:varchar(36);unique_index"`
	// 用户密码 Hash 后的密码
	Password string `gorm:"type:varchar(100)"`
	// 用户名
	Username sql.NullString `gorm:"type:varchar(10);unique_index"`
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
