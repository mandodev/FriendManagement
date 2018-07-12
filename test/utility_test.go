package test

import (
	"net/http"
	"testing"

	"github.com/FriendManagement/shared"
	"github.com/FriendManagement/shared/config"
	"github.com/stretchr/testify/assert"
)

func TestUtilityDispatchFunction(t *testing.T) {
	cfg, err := config.New("../shared/config/")
	assert.Empty(t, err)
	configuration := *cfg

	routerInstance := shared.NewRouter(configuration)
	router := routerInstance.SetupRouter()

	response := DispatchRequest(router, "GET", "/api/v1/ping", nil)
	assert.Equal(t, http.StatusOK, response.Code)
}
