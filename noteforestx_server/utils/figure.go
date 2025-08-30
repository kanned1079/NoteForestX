package utils

import (
	"fmt"
)

func ShowFigure(str string, version string, auth string) {
	//figure.NewFigure(str, "rectangles", true).Print()
	lens, _ := fmt.Printf("%s v%s\tby %s\n", str, version, auth)
	for i := 0; i < lens+12; i++ {
		fmt.Print("=")
	}
	fmt.Println()
}
