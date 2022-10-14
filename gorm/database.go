package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var globaldb *Database

type Database struct {
	connection *gorm.DB
}

func (database *Database) Get() (*gorm.DB, error) {
	if database.connection == nil {
		return nil, ErrNoDatabaseSet
	}
	return database.connection, nil
}

func (database *Database) Set(dsn string) error {
	db, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: false,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		DSN:                       dsn,
		DefaultStringSize:         256,
	}))

	if err != nil {
		return err
	}

	database.connection = db
	globaldb = database

	return nil
}

func Get() (*gorm.DB, error) {
	if globaldb == nil || globaldb.connection == nil {
		return nil, ErrNoDatabaseSet
	}

	return globaldb.connection, nil
}
