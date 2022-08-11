package user

import (
	"context"
	"errors"

	core "github.com/leaper-one/power-exchange/core/user"
	"github.com/leaper-one/power-exchange/util"

	// "github.com/leaper-one/pkg/util"

	"gorm.io/gorm"
)

func NewUserStore(db *util.DB) core.UserStore {
	return &userStore{db: db}
}

type userStore struct {
	db *util.DB
}

func toUpdateParams(user *core.User) map[string]interface{} {
	return map[string]interface{}{}
}

func update(db *util.DB, user *core.User) (int64, error) {
	updates := toUpdateParams(user)
	tx := db.Update().Model(user).Where("uuid = ?", user.UUID).Or("ID = ?", user.ID).Updates(updates)
	return tx.RowsAffected, tx.Error
}

// Save a user
func (s *userStore) Save(_ context.Context, user *core.User) error {
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
func (s *userStore) FindByUUID(_ context.Context, uuid string) (*core.User, error) {
	user := core.User{}
	if err := s.db.View().Where("uuid = ?", uuid).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, nil
}

// Find a user by user name
func (s *userStore) FindByUserName(_ context.Context, userName string) (*core.User, error) {
	user := core.User{}
	if err := s.db.View().Where("user_name = ?", userName).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, nil
}

// Find a user by email
func (s *userStore) FindByEmail(_ context.Context, email string) (*core.User, error) {
	user := core.User{}
	if err := s.db.View().Where("email = ?", email).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, nil
}

// Find a user by phone
func (s *userStore) FindByPhone(_ context.Context, phone string) (*core.User, error) {
	user := core.User{}
	if err := s.db.View().Where("phone = ?", phone).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, nil
}

// Delete a user by uuid
func (s *userStore) DeleteByUUID(_ context.Context, uuid string) error {
	if err := s.db.Update().Where("uuid = ?", uuid).Delete(&core.User{}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
