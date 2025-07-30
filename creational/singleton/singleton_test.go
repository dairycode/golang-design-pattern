package singleton_test

import (
	"testing"

	"github.com/dairycode/go-design-pattern/creational/singleton"
	"github.com/stretchr/testify/assert"
)

func TestGetEager(t *testing.T) {
	assert.True(t, singleton.GetEager() == singleton.GetEager())
}

func TestGetLazy(t *testing.T) {
	assert.True(t, singleton.GetLazy() == singleton.GetLazy())
}
