package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSha256Hash(t *testing.T) {
	assert.True(t, IsSha256Hash("e08abf9698f7d27e634de0d36cc974a0d908ec41c0a7e5e5738d2431f9a700e3"))
	// Not hex
	assert.False(t, IsSha256Hash("e08abf9698f7d27e634de0d36cc974a0d908ec41c0a7e5e5738d2431f9a700ez"))
	// Not 64 chars
	assert.False(t, IsSha256Hash(""))
}
