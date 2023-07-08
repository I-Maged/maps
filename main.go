package main

import "fmt"

func main() {
	// colors:=make(map[string]string)

	colors := map[string]string{
		"yellow": "#ffff00",
		"red":    "#ff0000",
		"blue":   "#0000ff",
		"black":  "000",
	}

	delete(colors, "black")
	colors["white"] = "fff"

	print(colors)
}

func print(colors map[string]string) {
	for i, color := range colors {
		fmt.Println(i, color)
	}
}
