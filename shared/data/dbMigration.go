package data

import (
	model "github.com/FriendManagement/module/friend/model"
	"github.com/FriendManagement/shared/config"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

//DbMigration : instance to hold dbmigration instance
type DbMigration struct {
	connection *gorm.DB
}

//NewDbMigration : intantiate new DBMigration instance
func NewDbMigration(cfg *config.Configuration) (*DbMigration, error) {
	dbFactory, err := NewDbFactory(cfg)

	if err != nil {
		glog.Errorf("%s", err)
		return nil, err
	}

	conn, err := dbFactory.DBConnection()

	if err != nil {
		glog.Errorf("%s", err)
		return nil, err
	}

	return &DbMigration{connection: conn}, nil
}

//Migrate : function to invoke gorm's automigrate
func (d *DbMigration) Migrate() (bool, error) {
	glog.Info("Start Database Migration")

	d.connection.AutoMigrate(
		&model.Connection{},
	)

	glog.Info("Database migration finished")
	return true, nil
}
