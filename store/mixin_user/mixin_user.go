package mixin_user

import (
	"context"
	"errors"

	core "github.com/leaper-one/power-exchange/core/mixin_user"
	"github.com/leaper-one/power-exchange/util"

	// "github.com/leaper-one/pkg/util"

	"gorm.io/gorm"
)

func NewUserStore(db *util.DB) core.MixinUserStore {
	return &mixinUserStore{db: db}
	// return &userStore{db: db}
}

type mixinUserStore struct {
	db *util.DB
}

func toUpdateParams(user *core.MixinUser) map[string]interface{} {
	return map[string]interface{}{}
}

func update(db *util.DB, user *core.MixinUser) (int64, error) {
	updates := toUpdateParams(user)
	tx := db.Update().Model(user).Where("uuid = ?", user.UUID).Or("client_id = ?", user.ClientID).Or("ID = ?", user.ID).Updates(updates)
	return tx.RowsAffected, tx.Error
}

// Save a mixin user
func (s *mixinUserStore) Save(_ context.Context, user *core.MixinUser) error {
	return s.db.Tx(func(tx *util.DB) error {
		var rows int64
		var err error

		rows, err = update(tx, user)
		if err != nil {
			return err
		}

		if rows == 0 {
			return tx.Update().Create(user).Error
		}

		return nil
	})
}

// Find a user by UUID
func (s *mixinUserStore) FindByUUID(_ context.Context, uuid string) (*core.MixinUser, error) {
	user := core.MixinUser{}
	if err := s.db.View().Where("uuid = ?", uuid).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, nil
}

// Find a user by client_id
func (s *mixinUserStore) FindByClientID(_ context.Context, clientID string) (*core.MixinUser, error) {
	user := core.MixinUser{}
	if err := s.db.View().Where("client_id = ?", clientID).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, nil
}

// Delete a user by uuid
func (s *mixinUserStore) DeleteByUUID(_ context.Context, uuid string) error {
	if err := s.db.Update().Where("uuid = ?", uuid).Delete(&core.MixinUser{}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}

// Delete a user by client_id
func (s *mixinUserStore) DeleteByClientID(_ context.Context, clientID string) error {
	if err := s.db.Update().Where("client_id = ?", clientID).Delete(&core.MixinUser{}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
