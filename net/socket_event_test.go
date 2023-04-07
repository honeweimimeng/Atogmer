package net

import (
	"Atogmer/driver"
	"Atogmer/driver/event"
	"testing"
	"time"
)

func TestNormalSocketEvent(t *testing.T) {
	ctx := event.Process(&Registry{}).LoadProperty(driver.DefaultConfig())
	event.UseEventBus(ctx).Execute()
	time.Sleep(200 * time.Second)
}
