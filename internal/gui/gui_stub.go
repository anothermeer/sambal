//go:build !gui

package gui

import "fmt"

func Available() bool {
	return false
}

func Run() {
	fmt.Println("GUI not compiled.")
}
