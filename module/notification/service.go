package notification

import (
	"errors"

	"github.com/satori/uuid"

	model "github.com/FriendManagement/module/friend/model"
	messages "github.com/FriendManagement/module/notification/messages"

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

//Subscribe : Service for subscribe update to an email
func (s *Service) Subscribe(request *messages.SubscribeRequest) (bool, error) {
	if request.Requestor == "" || request.Target == "" {
		return false, errors.New("Requestor or Target is undefined")
	}

	var connection model.Connection
	if err := s.db.Where("email1 =? AND email2=?", request.Requestor, request.Target).First(&connection).Error; err != nil {
		glog.Errorf("Error when create subscibe request,error : %s", err.Error())
		return false, err
	}

	if connection.ID == uuid.Nil {
		return false, errors.New("No friend connection found")
	}

	if connection.Subscribe {
		return false, errors.New("Already Subscribe")
	}

	connection.Subscribe = true
	s.db.Save(&connection)
	return true, nil
}
