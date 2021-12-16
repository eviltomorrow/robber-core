package mysql

import (
	"testing"

	"github.com/eviltomorrow/robber-core/tests"
)

func init() {
	tests.InitLogConfig("/tmp/robber-datasource/data.log")
}

func TestCreateClient(t *testing.T) {
	DSN = "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true&loc=Local"
	MaxOpen = 10
	MinOpen = 5
	Build()
}
