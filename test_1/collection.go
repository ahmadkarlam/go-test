package main

const MaxBall = 2

type Collection struct {
	containers map[int]Container
}

func NewCollection() Collection {
	return Collection{containers: make(map[int]Container)}
}

func (c *Collection) isThereAFullContainer() bool {
	for _, container := range c.containers {
		if container.Verified {
			return true
		}
	}

	return false
}

func (c *Collection) fill(key int) {
	container := NewContainer()
	if v, ok := c.containers[key]; ok {
		container = v
	}
	container.increaseNumberOfBall()
	if container.NumberOfBall == MaxBall {
		container.setVerified()
	}
	c.containers[key] = container
}
