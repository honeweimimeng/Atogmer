package net

import (
	"litecluster/driver"
	"litecluster/driver/event"
	"testing"
	"time"
)

func TestNormalSocketEvent(t *testing.T) {
	ctx := event.Process(&Registry{}).LoadProperty(driver.DefaultConfig())
	group := event.UseEventBus(ctx)
	group.Execute()
	time.Sleep(200 * time.Second)
}
