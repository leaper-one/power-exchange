package util

import (
	// "log"

	"github.com/sirupsen/logrus"
	// "gorm.io/driver/mysql"
	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// mixin_user "github.com/leaper-one/power-exchange/core/mixin_user"
	// user "github.com/leaper-one/power-exchange/core/user"
)

type DB struct {
	Write *gorm.DB
	Read  *gorm.DB
}

/*
func OpenDB(path string) *DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}
	err = db.AutoMigrate(&user.User{}, &mixin_user.MixinUser{})
	if err != nil {
		return nil
	}
	return &DB{Write: db, Read: db}
}

func OpenMysql(dsn string) *DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}
	err = db.AutoMigrate(&user.User{}, &mixin_user.MixinUser{})
	if err != nil {
		return nil
	}
	return &DB{Write: db, Read: db}
}
*/

func (db *DB) Update() *gorm.DB {
	return db.Write
}

func (db *DB) View() *gorm.DB {
	return db.Read
}

func (db *DB) Debug() *DB {
	return &DB{
		Write: db.Write.Debug(),
		Read:  db.Read.Debug(),
	}
}

func (db *DB) Begin() *DB {
	tx := db.Write.Begin()
	return &DB{
		Write: tx,
		Read:  db.Read,
	}
}

func (db *DB) Rollback() error {
	return db.Write.Rollback().Error
}

func (db *DB) Commit() error {
	return db.Write.Commit().Error
}

func (db *DB) RollbackUnlessCommitted() {
	if err := db.Write.Rollback().Error; err != nil {
		logrus.WithError(err).Error("DB: RollbackUnlessCommitted")
	}
}

func (db *DB) Tx(fn func(tx *DB) error) error {
	tx := db.Begin()
	defer tx.RollbackUnlessCommitted()

	if err := fn(tx); err != nil {
		return err
	}

	return tx.Commit()
}
