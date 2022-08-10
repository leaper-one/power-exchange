package mixin_user_core

import (
	"context"

	"gorm.io/gorm"
)

type (
	MixinUser struct {
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

	MixinUserStore interface {
		// Save a mixin user
		Save(ctx context.Context, mixinUser *MixinUser) error
		// Find a mixin user by UUID
		FindByUUID(ctx context.Context, uuid string) (*MixinUser, error)
		// Find a mixin user by client id
		FindByClientID(ctx context.Context, clientID string) (*MixinUser, error)
		// Delete a mixin user by UUID
		DeleteByUUID(ctx context.Context, uuid string) error
		// Delete a mixin user by client id
		DeleteByClientID(ctx context.Context, clientID string) error
	}
)
