package test

import (
	"net/http"
	"testing"

	"github.com/FriendManagement/shared"
	"github.com/FriendManagement/shared/config"
	"github.com/FriendManagement/test"
	"github.com/stretchr/testify/assert"
)

//TestRouterDiagnosticEndPoint : Test Diagnostic Endpoint
func TestRouterDiagnosticEndPoint(t *testing.T) {
	cfg, err := config.New("../../shared/config/")

	assert.Empty(t, err)

	configuration := *cfg
	routerInstance := shared.NewRouter(configuration)
	router := routerInstance.SetupRouter()

	response := test.DispatchRequest(router, "GET", "/api/v1/ping", nil)

	assert.Equal(t, http.StatusOK, response.Code)
}
