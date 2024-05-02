package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMigrationUp(t *testing.T) {
	m := Migrate("up")
	assert.Nil(t, m)
}

func TestMigrationDown(t *testing.T) {
	m := Migrate("down")
	assert.Nil(t, m)
}
