package friend

import "github.com/FriendManagement/common"

//Connection : Friend connection struct
type Connection struct {
	common.BaseModel
	Email1    string `gorm:"varchar(100);not null"`
	Email2    string `gorm:"varchar(100);not null"`
	Blocked   bool   `gorm:"not null"`
	Subscribe bool   `gorm:"not null"`
}
