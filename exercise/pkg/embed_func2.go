package main

import "fmt"

//多重继承

type Camera struct {
}

type Phone struct {
}

func (c *Camera) TakeApicture() string {
	return "Click"
}

func (p *Phone) Call() string {
	return "Ring Ring"
}

type CameraPhone struct {
	Camera
	Phone
}

func main() {
	cp := new(CameraPhone)
	fmt.Println("Our new CameraPhone exhibits multiple behaviors...")
	fmt.Println("It exhibits behavior of a Camera: ", cp.TakeApicture())
	fmt.Println("It works like a Phone too: ", cp.Call())
}
