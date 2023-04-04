package net

import (
	"litecluster/driver/event"
	"testing"
	"time"
)

func TestNormalSocketEvent(t *testing.T) {
	socket := &SocketTrigger{}
	eventCtx := event.UseTriggerEventLoop(socket).Init()
	go eventCtx.Boot().Execute()
	go func() {
		time.Sleep(2 * time.Second)
		socket.adviceChannel <- socket
	}()
	time.Sleep(20 * time.Second)
}
