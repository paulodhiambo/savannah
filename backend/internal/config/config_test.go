package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoad(t *testing.T) {
	err := Load()
	assert.NoError(t, err, "Error loading config")
}
