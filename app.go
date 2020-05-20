package pico

import (
	"log"
)

// App type
type App struct{}

// Start func
func (app *App) Start() {
	log.Println("Core: Start() method is not implemented")
}

// Stop func
func (app *App) Stop() {
	log.Println("Core: Stop() method is not implemented")
}
