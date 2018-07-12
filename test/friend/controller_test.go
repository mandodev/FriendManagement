package friend

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/FriendManagement/shared"
	"github.com/FriendManagement/shared/config"
	"github.com/FriendManagement/test"
	"github.com/stretchr/testify/assert"
)

func TestCreateConnection(t *testing.T) {
	cfg, err := config.New("../../shared/config/")
	assert.Empty(t, err)
	configuration := *cfg

	payload := bytes.NewBuffer([]byte(`{"friends":["andy@example.com", "john@example.com"]}`))
	routerInstance := shared.NewRouter(configuration)
	router := routerInstance.SetupRouter()

	response := test.DispatchRequest(router, "POST", "/api/v1/friend/connect", payload)
	s := string(response.Body.Bytes())
	assert.NotEmpty(t, s)
	assert.Equal(t, http.StatusOK, response.Code)

}

func TestFriendList(t *testing.T) {
	cfg, err := config.New("../../shared/config/")
	assert.Empty(t, err)
	configuration := *cfg

	payload := bytes.NewBuffer([]byte(`{"email":"andy@example.com"}`))
	routerInstance := shared.NewRouter(configuration)
	router := routerInstance.SetupRouter()

	response := test.DispatchRequest(router, "POST", "/api/v1/friend/list", payload)
	s := string(response.Body.Bytes())
	assert.NotEmpty(t, s)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetCommonFriend(t *testing.T) {
	cfg, err := config.New("../../shared/config/")
	assert.Empty(t, err)
	configuration := *cfg

	payloadOne := bytes.NewBuffer([]byte(`{"friends":["friend1@example.com", "common@example.com"]}`))
	payloadTwo := bytes.NewBuffer([]byte(`{"friends":["friend2@example.com", "common@example.com"]}`))

	routerInstance := shared.NewRouter(configuration)
	router := routerInstance.SetupRouter()

	responseOne := test.DispatchRequest(router, "POST", "/api/v1/friend/connect", payloadOne)
	bodyOne := string(responseOne.Body.Bytes())
	assert.NotEmpty(t, bodyOne)

	responseTwo := test.DispatchRequest(router, "POST", "/api/v1/friend/connect", payloadTwo)
	bodyTwo := string(responseTwo.Body.Bytes())
	assert.NotEmpty(t, bodyTwo)

	payloadEmail := bytes.NewBuffer([]byte(`{"friends":["friend1@example.com", "friend2@example.com"]}`))

	responseCommon := test.DispatchRequest(router, "POST", "/api/v1/friend/common", payloadEmail)
	bodyCommon := string(responseCommon.Body.Bytes())
	assert.Equal(t, "{\"count\":1,\"friends\":[\"common@example.com\"],\"success\":true}", bodyCommon)
}
