package main

import (
	"serverpunk/internal/monitor"
	"serverpunk/internal/platform/server"
	"time"
)

func main() {
	monitor.Start(1*time.Second, 100.0)
	server.Start()
}
