package net

import (
	"atogmer/driver"
	"atogmer/driver/event"
	"testing"
	"time"
)

func TestNormalSocketEvent(t *testing.T) {
	ctx := event.Process(&Registry{}).LoadProperty(driver.DefaultConfig())
	event.UseEventBus(ctx).Execute()
	time.Sleep(200 * time.Second)
}
