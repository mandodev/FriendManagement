package data

import (
	"errors"

	"github.com/FriendManagement/shared/config"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"

	// importing all possible database dialect
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//DbFactory :  model for database factory instance
type DbFactory struct {
	config config.Configuration
}

//NewDbFactory : function to generate new database factory
func NewDbFactory(cfg *config.Configuration) (*DbFactory, error) {
	if cfg == nil {
		return nil, errors.New("Error Intantiate new db instance, config is null")
	}
	return &DbFactory{config: *cfg}, nil
}

//DBConnection : open connection to Database server
func (factory *DbFactory) DBConnection() (*gorm.DB, error) {
	db, err := gorm.Open(factory.config.Database.DbType, factory.config.Database.ConnectionURI)
	if err != nil {
		glog.Fatalf("Failed to establish connection to Database, error: %s", err)
		return nil, err
	}

	return db, err
}
