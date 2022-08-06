package core

import (
	"gorm.io/gorm"
)

type MixinUser struct {
	gorm.Model
	// User UUID
	UUID string `gorm:"type:varchar(36);unique_index"`
	// Mixin UUID
	ClientID string `json:"client_id,omitempty"`
	// Mixin session id
	SessionID string `json:"session_id,omitempty"`
	// Mixin Pin Token
	PinToken string `json:"pin_token,omitempty"`
	// Mixin Private Key
	PrivateKey string `json:"private_key,omitempty"`
}
