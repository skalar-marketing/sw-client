package sw

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClient_Context(t *testing.T) {
	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
	assert.NoError(t, err)
	result, err := client.Context()
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
