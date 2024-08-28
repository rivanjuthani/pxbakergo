package main

import (
	"fmt"
)

func main() {
	px := NewPerimeterX("", true)
	result := px.SubmitSensor()
	fmt.Println(result)
	// px.PxUuid = "9d92d005-5e00-11ef-90d2-1568ec2dc4f9"
	// fmt.Println(px.PxHello("", "2"))
}
