package main

type Container struct {
	NumberOfBall int  `json:"number_of_ball"`
	Verified     bool `json:"verified"`
}

func NewContainer() Container {
	return Container{
		NumberOfBall: 0,
		Verified:     false,
	}
}

func (c *Container) increaseNumberOfBall() {
	c.NumberOfBall += 1
}

func (c *Container) setVerified() {
	c.Verified = true
}
