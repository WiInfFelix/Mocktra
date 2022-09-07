package main

import (
	"fmt"
	"mocktra/generator"
)

func main() {
	fmt.Println("Starting app...")
	//userinterface.MakeUI()
	generator.InitDefaults()
	fmt.Println(generator.DefaultLocations)
	fmt.Println(generator.DefaultItems)
	generator.SimulateDay()
}
