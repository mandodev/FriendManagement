package friend

import (
	"testing"

	"github.com/FriendManagement/module/friend"
	"github.com/FriendManagement/shared/config"
	"github.com/FriendManagement/shared/data"
	"github.com/stretchr/testify/assert"
)

func TestServiceCreateConection(t *testing.T) {
	cfg, _ := config.New("../../shared/config/")
	dbInstance, _ := data.NewDbFactory(cfg)
	conn, _ := dbInstance.DBConnection()

	service, err := friend.NewService(conn)

	assert.Nil(t, err)
	assert.NotNil(t, service)

	emails := []string{"test1@example.com", "test2@example.com"}

	result, err := service.CreateConnection(emails)
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestServiceGetFriend(t *testing.T) {
	cfg, _ := config.New("../../shared/config/")
	dbInstance, _ := data.NewDbFactory(cfg)
	conn, _ := dbInstance.DBConnection()

	service, err := friend.NewService(conn)

	assert.Nil(t, err)
	assert.NotNil(t, service)

	email := "test1@example.com"

	result, err := service.ConnectionList(email)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 2)
}
