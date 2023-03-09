package driver

import "litecluster/event"

type EventLooper interface {
	Registry(events []event.Proto)
}
