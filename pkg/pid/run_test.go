package pid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFile(t *testing.T) {
	_assert := assert.New(t)
	err := CreatePidFile("test.pid")
	_assert.Nil(err)

	DestroyFile()
}
