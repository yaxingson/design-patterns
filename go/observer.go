package main

type Listener func(...interface{}) interface{}

type EventEmitter struct {
	listeners []Listener
}

func (e *EventEmitter) on()     {}
func (e *EventEmitter) emit()   {}
func (e *EventEmitter) remove() {}

func init() {

}
