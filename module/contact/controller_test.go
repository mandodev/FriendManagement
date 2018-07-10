package contact

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/FriendManagement/shared"
	"github.com/FriendManagement/test"
	"github.com/stretchr/testify/assert"
)

func TestCreateConnection(t *testing.T) {
	payload := bytes.NewBuffer([]byte(`{friends:['andy@example.com', 'john@example.com']}`))
	router := shared.SetupRouter()
	response := test.DispatchRequest(router, "POST", "api/v1/friend", payload)
	assert.Equal(t, http.StatusTemporaryRedirect, response.Code)
}
