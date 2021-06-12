package module

import (
	"fmt"
	"os"
	"os/signal"
)

// NewSignalChannel 创建 signal channel
func NewSignalChannel() chan os.Signal {
	fmt.Println("[mini]build signal channel")
	channel := make(chan os.Signal, 1)
	signal.Notify(channel)
	return channel
}
