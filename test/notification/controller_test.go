package notification

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/FriendManagement/shared"
	"github.com/FriendManagement/shared/config"
	"github.com/FriendManagement/test"
	"github.com/stretchr/testify/assert"
)

func TestSubscribeUpdate(t *testing.T) {
	cfg, err := config.New("../../shared/config/")
	assert.Empty(t, err)
	configuration := *cfg

	payload := bytes.NewBuffer([]byte(`{"requestor":"requestor@example.com", "target":"target@example.com"}`))
	routerInstance := shared.NewRouter(configuration)
	router := routerInstance.SetupRouter()

	response := test.DispatchRequest(router, "POST", "/api/v1/notification/subscribe", payload)
	s := string(response.Body.Bytes())
	assert.NotEmpty(t, s)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "{\"success\":true}", s)
}
