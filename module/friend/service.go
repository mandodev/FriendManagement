package friend

import (
	"errors"

	model "github.com/FriendManagement/module/friend/model"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

//Service : Service for contact module
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

//CreateConnection : function to connect two email as a friend
func (s *Service) CreateConnection(emails []string) (bool, error) {
	if emails == nil || len(emails) != 2 {
		return false, errors.New("Email collection is empty or lenght is not is equal 2")
	}

	var connections []model.Connection

	tx := s.db.Begin()
	if tx.Error != nil {
		glog.Errorf("Failed to begin transaction when creating Friend Connection : %s", tx.Error)
		return false, errors.New("Failed to begin transaction when creating Friend Connection " + tx.Error.Error())
	}

	tx.Where("(email1 = ? AND email2 = ?) OR (email1 = ? AND email2 = ?)", emails[0], emails[1], emails[1], emails[0]).Find(&connections)

	if len(connections) > 0 {
		for _, v := range connections {
			if v.Blocked {
				return false, errors.New("Friend Requet failed, user is blocked your email")
			}
		}

		if len(connections) == 2 {
			return false, errors.New("Both of you already friend :)")
		}

		if err := tx.Create(&model.Connection{Email1: connections[0].Email2, Email2: connections[0].Email1, Blocked: false, Subscribe: false}).Error; err != nil {
			tx.Rollback()
			return false, errors.New("Error when create connection " + err.Error())
		}

		tx.Commit()

		return true, nil
	}

	error1 := tx.Create(&model.Connection{Email1: emails[0], Email2: emails[1], Blocked: false, Subscribe: false}).Error
	error2 := tx.Create(&model.Connection{Email1: emails[1], Email2: emails[0], Blocked: false, Subscribe: false}).Error

	if error1 != nil || error2 != nil {
		glog.Errorf("error create connection %s", error1.Error())
		glog.Errorf("error create connection %s", error2.Error())

		tx.Rollback()
		return false, errors.New("Error when create connection")
	}
	tx.Commit()
	return true, nil

}

//ConnectionList : function to list all of user's friend
func (s *Service) ConnectionList(email string) ([]string, error) {
	if email == "" {
		return nil, errors.New("Email field empty")
	}

	var mails []string
	var connections []model.Connection

	errs := s.db.Where("Email1 = ?", email).Find(&connections, &model.Connection{Email1: email}).Pluck("Email2", &mails).GetErrors()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if len(mails) == 0 {
		return nil, errors.New("This user doesn't have any friend")
	}

	return mails, nil
}

//CommonConnection : function to get common connection between two emails
func (s *Service) CommonConnection(emails []string) ([]string, error) {
	if emails == nil || len(emails) != 2 {
		return nil, errors.New("Email collection is empty or lenght is not is equal 2")
	}

	var mails []string
	var connections []model.Connection

	errs := s.db.Select("ANY_VALUE(Email1), Email2").Where("email1=? OR email1=?", emails[0], emails[1]).Group("email2").Having("count(*) > 1").Find(&connections).Pluck("Email2", &mails).GetErrors()

	if len(errs) > 0 {
		glog.Error(errs)
		return nil, errors.New("error occured when get common friend, please see logs")
	}

	if len(mails) == 0 {
		return nil, errors.New("This user doesn't have any friend")
	}

	return mails, nil
}
