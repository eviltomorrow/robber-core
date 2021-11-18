package tests

import (
	"log"
	"os"

	"github.com/eviltomorrow/robber-core/pkg/zlog"
)

func InitLogConfig(path string) {
	global, prop, err := zlog.InitLogger(&zlog.Config{
		Level:            "debug",
		Format:           "text",
		DisableTimestamp: false,
		File: zlog.FileLogConfig{
			Filename:   path,
			MaxSize:    20,
			MaxDays:    10,
			MaxBackups: 30,
		},
		DisableStacktrace: true,
	})
	if err != nil {
		log.Printf("Error: Setup log config failure, nest error: %v", err)
		os.Exit(1)
	}
	zlog.ReplaceGlobals(global, prop)
}
