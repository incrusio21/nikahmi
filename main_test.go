package main

import (
	"testing"

	"github.com/incrusio21/nikahmi/config"
	"github.com/stretchr/testify/assert"
)

func TestViperConfig(t *testing.T) {
	_, err := config.Read()
	assert.Nil(t, err)
}
