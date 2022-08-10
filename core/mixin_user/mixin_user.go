package core

import (
	"gorm.io/gorm"
)

type MixinUser struct {
	gorm.Model
	// User UUID
	UUID string `gorm:"type:varchar(36);unique_index"`
	// Mixin UUID
	ClientID string `gorm:"type:varchar(36);unique_index"`
	// Mixin session id
	SessionID string `gorm:"type:varchar(100);unique_index"`
	// Mixin Pin Token
	PinToken string `gorm:"type:varchar(100);unique_index"`
	// Mixin Private Key
	PrivateKey string `gorm:"type:varchar(100);unique_index"`
}
