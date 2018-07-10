package test

import (
	"net/http"
	"testing"

	"github.com/FriendManagement/shared"
	"github.com/stretchr/testify/assert"
)

func TestUtilityDispatchFunction(t *testing.T) {
	router := shared.SetupRouter()
	response := DispatchRequest(router, "GET", "/v1/api/ping", nil)
	assert.Equal(t, http.StatusOK, response.Code)
}
