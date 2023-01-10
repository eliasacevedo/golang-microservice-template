package services

import (
	"github.com/eliasacevedo/golang-microservice-template/config"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDefaultDB() (*gorm.DB, error) {
	return GetDB(config.GetDatabase())
}

func GetDB(connectionString string) (*gorm.DB, error) {
	if db == nil {
		c := gorm.Config{}
		d, err := gorm.Open(sqlserver.Open(connectionString), &c)
		if err != nil {
			return nil, err
		}
		db = d
	}

	return db, nil
}
