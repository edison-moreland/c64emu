package clock

import (
	"fmt"
	"time"
)

// Clock needs to be a Ticker that can be started and stopped at will, with an adjustable interval

type Clock struct {
	out          chan bool
	lastTickTime time.Time
	interval     time.Duration
	active       bool
}

func New(i time.Duration) *Clock {
	return &Clock{
		out:          make(chan bool),
		lastTickTime: time.Now(),
		interval:     i,
		active:       false,
	}
}

func (c *Clock) C() <-chan bool {
	return c.out
}

func (c *Clock) Start() {
	c.active = true
	c.afterInterval()
}

func (c *Clock) afterInterval() {
	lastTickDuration := c.Tick()

	if lastTickDuration > c.interval {
		fmt.Printf("clock --- clock_drift=%s\n", c.interval-lastTickDuration)
	}

	if c.active {
		time.AfterFunc(c.interval, c.afterInterval)
	}
}

func (c *Clock) Stop() {
	c.active = false
}

func (c *Clock) Interval(i time.Duration) {
	c.interval = i
}

func (c *Clock) Tick() time.Duration {
	c.out <- true
	tickTime := time.Now()

	tickDuration := c.lastTickTime.Sub(tickTime)

	c.lastTickTime = tickTime

	return tickDuration
}
