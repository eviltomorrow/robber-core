package system

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/eviltomorrow/robber-core/pkg/znet"
	"github.com/eviltomorrow/robber-core/pkg/ztime"
)

var (
	Pid         = os.Getpid()
	Pwd         string
	LaunchTime  = time.Now()
	HostName    string
	OS          = runtime.GOOS
	Arch        = runtime.GOARCH
	RunningTime = func() string {
		return ztime.FormatDuration(time.Since(LaunchTime))
	}
	IP string
)

func InitEnv() error {
	path, err := filepath.Abs(".")
	if err != nil {
		return fmt.Errorf("get current dir failure, nest error: %v", err)
	}
	Pwd = path

	name, err := os.Hostname()
	if err == nil {
		HostName = name
	}

	localIP, err := znet.GetLocalIP()
	if err == nil {
		IP = localIP
	}
	return nil
}
