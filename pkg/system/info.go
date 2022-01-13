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
	RootDir     string
	Pid         = os.Getpid()
	LaunchTime  = time.Now()
	HostName    string
	OS          = runtime.GOOS
	Arch        = runtime.GOARCH
	RunningTime = func() string {
		return ztime.FormatDuration(time.Since(LaunchTime))
	}
	IP string
)

func init() {
	path, err := os.Executable()
	if err != nil {
		panic(fmt.Errorf("get executeable path failure, nest error: %v", err))
	}
	RootDir, err = filepath.Abs(path)
	if err != nil {
		panic(fmt.Errorf("abs RootDir failure, nest error: %v", err))
	}

	name, err := os.Hostname()
	if err == nil {
		HostName = name
	}

	localIP, err := znet.GetLocalIP()
	if err == nil {
		IP = localIP
	}
}
