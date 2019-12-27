package main

import (
	"secret-santa/app"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		panic(err)
	}
	a.Start()
}
