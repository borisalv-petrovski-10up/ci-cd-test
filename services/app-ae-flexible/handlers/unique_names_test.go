package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/borisalv-petrovski-10up/ci-cd-test/services/app-ae-flexible/handlers"
)

func TestRegisterUser(t *testing.T) {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	// Act
	handlers.UniqueNames(res, req)

	// Assert
	assert.Equal(t, "Hello, [Samuel John Falcooo]!", res.Body.String())
}
