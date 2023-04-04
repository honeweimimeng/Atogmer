package event

import "litecluster/driver"

type TriggerRegistry struct {
	eventTrigger Trigger
	lastTrigger  Trigger
}

func (r *TriggerRegistry) Process(ctx driver.ExecutorContext) {
	if bus, ok := ctx.Group().(*ExecutorBus); ok {
		if r.lastTrigger != nil {
			r.lastTrigger.Child(r.eventTrigger)
			return
		}
		r.BuildChild(bus.eventTrigger, r.eventTrigger)
	}
}

func (r *TriggerRegistry) BuildChild(trigger Trigger, child Trigger) {
	if trigger.Next() != nil {
		r.BuildChild(trigger.Next(), child)
		return
	}
	trigger.Child(child)
	r.lastTrigger = child
}
