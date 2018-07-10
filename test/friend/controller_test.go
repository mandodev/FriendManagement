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
	assert.Equal(t, http.StatusOK, response.Code)
}
