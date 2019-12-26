package main

import (
	"secret-santa/app"
	"fmt"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		panic(err)
	}
	fmt.Print(a)
}