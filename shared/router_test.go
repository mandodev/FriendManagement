package shared

import (
	"net/http"
	"testing"

	"github.com/FriendManagement/test"
	"github.com/stretchr/testify/assert"
)

//TestRouterDiagnosticEndPoint : Test Diagnostic Endpoint
func TestRouterDiagnosticEndPoint(t *testing.T) {
	router := SetupRouter()

	response := test.DispatchRequest(router, "GET", "/v1/api/ping", nil)

	assert.Equal(t, http.StatusOK, response.Code)
}
