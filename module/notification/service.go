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
func (s *Service) Subscribe(request *messages.NotificationRequest) (bool, error) {
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
	tx := s.db.Begin()
	connection.Subscribe = true

	if err := tx.Save(connection).Error; err != nil {
		glog.Errorf("Error when create subscibe request,error : %s", err.Error())
		return false, err
	}

	tx.Commit()

	return true, nil
}

//Block : Service for block update to an email
func (s *Service) Block(request *messages.NotificationRequest) (bool, error) {
	if request.Requestor == "" || request.Target == "" {
		return false, errors.New("Requestor or Target is undefined")
	}

	var connection model.Connection
	if err := s.db.Where("email1 =? AND email2=?", request.Requestor, request.Target).First(&connection).Error; err != nil {
		glog.Errorf("Error when create subscibe request,error : %s", err.Error())
		return false, err
	}

	if connection.ID == uuid.Nil {
		tx := s.db.Begin()
		error1 := tx.Create(&model.Connection{Email1: request.Requestor, Email2: request.Target, Blocked: true, Subscribe: false}).Error

		if error1 != nil {
			glog.Errorf("error create connection %s", error1.Error())
			tx.Rollback()
			return true, nil
		}
		tx.Commit()

		return true, nil
	}

	if connection.Blocked {
		return false, errors.New("Already Blocked")
	}

	tx := s.db.Begin()
	connection.Blocked = true
	connection.Subscribe = false

	if err := tx.Save(connection).Error; err != nil {
		tx.Rollback()
		glog.Errorf("Error when create subscibe request,error : %s", err.Error())
		return false, err
	}

	tx.Commit()

	return true, nil
}
