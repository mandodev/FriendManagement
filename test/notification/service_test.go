package notification

import (
	"testing"

	"github.com/FriendManagement/module/notification"
	messages "github.com/FriendManagement/module/notification/messages"
	"github.com/FriendManagement/shared/config"
	"github.com/FriendManagement/shared/data"
	"github.com/stretchr/testify/assert"
)

func TestServiceSubscribe(t *testing.T) {
	cfg, _ := config.New("../../shared/config/")
	dbInstance, _ := data.NewDbFactory(cfg)
	conn, _ := dbInstance.DBConnection()

	service, err := notification.NewService(conn)
	request := &messages.NotificationRequest{Requestor: "friend1@example.com", Target: "common@example.com"}
	result, err := service.Subscribe(request)

	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestServiceBlockEmail(t *testing.T) {
	cfg, _ := config.New("../../shared/config/")
	dbInstance, _ := data.NewDbFactory(cfg)
	conn, _ := dbInstance.DBConnection()

	service, err := notification.NewService(conn)
	request := &messages.NotificationRequest{Requestor: "friend1@example.com", Target: "common@example.com"}
	result, err := service.Block(request)

	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestServiceUpdateList(t *testing.T) {
	cfg, _ := config.New("../../shared/config/")
	dbInstance, _ := data.NewDbFactory(cfg)
	conn, _ := dbInstance.DBConnection()

	service, err := notification.NewService(conn)
	request := &messages.UpdateRequest{Sender: "friend1@example.com", Text: "Hello! common@example.com"}
	result, err := service.Update(request)

	assert.NotNil(t, result)
	assert.Nil(t, err)
}
