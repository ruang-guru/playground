package main

import "fmt"

// agar mudah dijalankan, implementasi interface akan ditaruh di file yang sama dengan context

type eyePowerStrategy interface { //ini adalah strategy
	useEyePower()
}

type cat struct {
	eyePower eyePowerStrategy
}

func createCat(eyePower eyePowerStrategy) *cat {
	return &cat{
		eyePower: eyePower,
	}
}
func (c *cat) walk() {
	fmt.Println("cat walk")
}
func (c *cat) talk() {
	fmt.Println("cat meow")
}
func (c *cat) sleep() {
	fmt.Println("cat sleep")
}
func (c *cat) sit() {
	fmt.Println("cat sit")
}
func (c *cat) eat() {
	fmt.Println("cat eat")
}
func (c *cat) useEyePower() {
	c.eyePower.useEyePower()
} // use eyePower implementation of useEyePower

func main() {
	noPower := &noEyePower{}
	laserPower := &laserEyePower{}

	normalCat := createCat(noPower)
	normalCat2 := createCat(noPower)
	laserCat := createCat(laserPower)
	normalCat.useEyePower()
	normalCat2.useEyePower()
	laserCat.useEyePower()
}

type laserEyePower struct{}

func (e *laserEyePower) useEyePower() {
	fmt.Println("shoot laser")
}

type noEyePower struct{}

func (e *noEyePower) useEyePower() {
	fmt.Println("do nothing")
}
