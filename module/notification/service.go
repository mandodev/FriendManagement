package notification

import (
	"errors"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

//Service : Notification Controller
type Service struct {
	db *gorm.DB
}

//NewService : instantiate new contact service
func NewService(Db *gorm.DB) (*Service, error) {
	if Db == nil {
		glog.Error("failed to intantiate Service , Db instance is null")
		return nil, errors.New("failed to intantiate Service , Db instance is null")
	}
	return &Service{db: Db}, nil
}
