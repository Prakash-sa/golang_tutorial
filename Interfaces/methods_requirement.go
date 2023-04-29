package main

type car interface {
	startEngine() bool
	stopEngine()
}

type sportCar struct {
	started bool
}

func (c *sportCar) startEngine() bool {
	if c.started {
		return false
	}
	c.started = true
	return true
}

func (c *sportCar) stopEngine() {
	c.started = false
}

func main() {

}
