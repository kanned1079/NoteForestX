package utils

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
)

func ShowFigure(str string, version string, auth string) {
	figure.NewFigure(str, "rectangles", true).Print()
	lens, _ := fmt.Printf("Server v%s\tby %s\n", version, auth)
	for i := 0; i < lens+12; i++ {
		fmt.Print("=")
	}
	fmt.Println()
}
