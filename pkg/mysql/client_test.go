package mysql

import (
	"testing"

	"github.com/eviltomorrow/robber-core/tests"
)

func init() {
	tests.InitLogConfig("/var/log/robber-datasource-sina/data.log")
}

func TestCreateClient(t *testing.T) {
	DSN = "root:root@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=true&loc=Local"
	MaxOpen = 10
	MinOpen = 5
	Build()
}
