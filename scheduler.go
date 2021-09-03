package scheduler

var globalSchedule = &schedule{}

// New create schedule instance
func New() Scheduler {
	return &schedule{}
}

// Add a handler to the schedule
func Add(handler Handler) {
	globalSchedule.Add(handler)
}

// Run execute all handlers in the schedule
func Run(data interface{}) interface{} {
	return globalSchedule.Run(data)
}

// Remove remove handler from scheduler
func Remove(handler Handler) {
	globalSchedule.Remove(handler)
}

// Clear all handlers in the air conditioner
func Clear() {
	globalSchedule.Clear()
}