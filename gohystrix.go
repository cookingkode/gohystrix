
// Package gohystrix provides functionality to make calls which can take variable amount of time 
// Example could be call to external web services which can be latent
package gohystrix

import (
	"time"
)

type Command struct {
	todo     func(interface{}) interface{}
	fallback func(interface{}) interface{}
	timeout  time.Duration
}

// NewCommand creates a new Command object
// todo : main function to be called
// fallback : function to be called on timeout
// tout : timeout in milliseconds
//
// NOTE: main and fallback functions need to have
//     func (interface{}) interface{} 
// signature
func NewCommand(todo func(interface{}) interface{},
	fallback func(interface{}) interface{},
	tout time.Duration,
) *Command {

	return &Command{
		todo:     todo,
		timeout:  tout,
		fallback: fallback,
	}
}

// Run starts execution of the underlying function
// param is the parameter for the main/fallback function
func (c *Command) Run(param interface{}) interface{} {

	c1 := make(chan interface{})
	defer close(c1)

	go wrap(c1, c.todo, param)

	select {
	case res := <-c1:
		return res
	case <-time.After(time.Millisecond * c.timeout):
		return c.fallback(param)
	}

}

func wrap(c chan interface{}, todo func(interface{}) interface{}, in interface{}) {
	c <- todo(in)
}
