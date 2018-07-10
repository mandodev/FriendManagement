package test

import (
	"net/http"
	"testing"

	"github.com/FriendManagement/shared"
	"github.com/stretchr/testify/assert"
)

//TestRouterDiagnosticEndPoint : Test Diagnostic Endpoint
func TestRouterDiagnosticEndPoint(t *testing.T) {
	router := shared.SetupRouter()

	response := DispatchRequest(router, "GET", "/v1/api/ping", nil)

	assert.Equal(t, http.StatusOK, response.Code)
}
