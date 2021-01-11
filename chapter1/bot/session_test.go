package bot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewSession tests
// go test -run TestNewSession -v
func TestNewSession(t *testing.T) {
	s := NewAgentSession(dfStaging, gcpProjectID)
	assert.NotNil(t, s)
	t.Logf("%s", s.path)
}
