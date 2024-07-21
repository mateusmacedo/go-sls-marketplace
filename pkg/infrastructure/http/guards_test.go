package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMethodAllowed(t *testing.T) {
	allowedMethods := []string{"GET", "POST", "PUT"}
	guard := NewHttpMethodGuard(allowedMethods)

	t.Run("Allowed method", func(t *testing.T) {
		method := "GET"
		allowed := guard.IsMethodAllowed(method)
		assert.True(t, allowed)
	})

	t.Run("Disallowed method", func(t *testing.T) {
		method := "DELETE"
		allowed := guard.IsMethodAllowed(method)
		assert.False(t, allowed)
	})
}
