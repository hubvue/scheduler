package scheduler

// Handler Schedule handler interface
type Handler interface {
	Start(data interface{}) interface{}
	End(data interface{}) interface{}
}

// Scheduler Schedule methods interface
type Scheduler interface {
	Add(handler Handler)
	Run(data interface{}) interface{}
	Remove(handler Handler)
	Clear()
}

type schedule struct {
	handlers []Handler
}

// Add a handler to the schedule
func (schedule *schedule) Add(handler Handler) {
	schedule.handlers = append(schedule.handlers, handler)
}

// Run execute all handlers in the schedule
func (schedule *schedule) Run(data interface{}) interface{} {
	return schedule.runEndHandler(schedule.runStartHandler(data))
}

// Remove remove handler from scheduler
func (schedule *schedule) Remove(handler Handler) {
	idx := -1
	for i := 0; i < len(schedule.handlers); i++ {
		if schedule.handlers[i] == handler {
			idx = i
			break
		}
	}
	if idx != -1 {
		schedule.handlers = append(schedule.handlers[:idx], schedule.handlers[idx+1:]...)
	}
}

// Clear all handlers in the air conditioner
func (schedule *schedule) Clear() {
	schedule.handlers = make([]Handler, 0)
}

func (schedule *schedule) runStartHandler(data interface{}) interface{} {
	for i := 0; i < len(schedule.handlers); i++ {
		data = schedule.handlers[i].Start(data)
	}
	return data
}
func (schedule *schedule) runEndHandler(data interface{}) interface{} {
	for i := len(schedule.handlers) - 1; i >= 0; i-- {
		data = schedule.handlers[i].End(data)
	}
	return data
}
