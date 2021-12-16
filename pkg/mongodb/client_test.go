package mongodb

import (
	"testing"

	"github.com/eviltomorrow/robber-core/tests"
)

func init() {
	tests.InitLogConfig("/tmp/data.log")
}

func TestCreateClient(t *testing.T) {
	DSN = "mongodb://127.0.0.1:27017"
	MaxOpen = 10
	Build()
}
