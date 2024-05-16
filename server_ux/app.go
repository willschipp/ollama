package main

import (
	"context"

	"server_ux/pkg/registry"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetImages() []string {
	// return response
	return registry.GetImages()
}

func (a *App) GetTags(name string) []string {
	// return response
	return registry.GetTags(name)
}
