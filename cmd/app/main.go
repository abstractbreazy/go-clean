package main

import (
	"log"

	"go-clean/internal/app"
)

func initApp() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatal("Fail to create app: ", err)
	}

	app.SetGlobalApp(a)
}

func main() {
	initApp()
	a, err := app.GetGlobalApp()
	if err != nil {
		log.Fatal(err)
	}

	a.StartHTTP()
}
