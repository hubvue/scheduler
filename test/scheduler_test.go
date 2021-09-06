package test

import (
	"testing"

	"github.com/hubvue/scheduler"
)

type handler1 struct{}

func (handler handler1) Start(data interface{}) interface{} {
	num := data.(int)
	return num + num
}

func (handler handler1) End(data interface{}) interface{} {
	num := data.(int)
	return num * num
}

type handler2 struct{}

func (handler handler2) Start(data interface{}) interface{} {
	num := data.(int)
	return num * num
}

func (handler handler2) End(data interface{}) interface{} {
	num := data.(int)
	return num + num
}

func TestScheduler(t *testing.T) {
	t.Run("Single handler", func(t *testing.T) {
		scheduler.Add(handler1{})
		res := scheduler.Run(2)
		scheduler.Clear()
		expected := 16

		if res != expected {
			t.Errorf("expected '%d' but got '%d'", expected, res)
		}
	})
	t.Run("Multiple handlers", func(t *testing.T) {
		scheduler.Add(handler1{})
		scheduler.Add(handler2{})
		res := scheduler.Run(2)
		scheduler.Clear()
		expected := 1024

		if res != expected {
			t.Errorf("expected '%d' but got '%d'", expected, res)
		}
	})
	t.Run("Remove handler", func(t *testing.T) {
		scheduler.Add(handler1{})
		scheduler.Add(handler2{})
		scheduler.Remove(handler1{})
		res := scheduler.Run(2)
		scheduler.Clear()
		expected := 8

		if res != expected {
			t.Errorf("expected '%d' but got '%d'", expected, res)
		}
	})
	t.Run("New schedule instance", func(t *testing.T) {
		schedule := scheduler.New()
		schedule.Add(handler1{})
		schedule.Add(handler2{})
		res := schedule.Run(2)
		schedule.Clear()
		expected := 1024

		if res != expected {
			t.Errorf("expected '%d' but got '%d'", expected, res)
		}
	})
}
