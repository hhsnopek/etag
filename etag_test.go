package etag

import (
	"testing"

	"github.com/tj/assert"
)

var body = []byte("Hello World")

func TestGenerate(t *testing.T) {
	result := Generate(body, false)
	assert.NotEmpty(t, result)
	assert.NotContains(t, result, "W/")
}

func TestGenerateWeak(t *testing.T) {
	result := Generate(body, true)
	assert.NotEmpty(t, result)
	assert.Contains(t, result, "W/")
}
