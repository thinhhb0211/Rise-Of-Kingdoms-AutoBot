package main

import (
	"Rise-Of-Kingdoms-AutoBot/pkg/config"
	"context"
)

// App struct
type App struct {
	ctx     context.Context
	devices []Device
	config  *config.Configuration
}

// NewApp creates a new App application struct
func NewApp(config *config.Configuration) *App {
	return &App{
		config: config,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type Device struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	Port string `json:"port"`
}

func (a *App) LoadConfig() *config.Configuration {
	return a.config
}

func (a *App) SetConfig(config *config.Configuration) {
	a.config = config
}

func (a *App) AddDevice(name, ip, port string) []Device {
	device := Device{Name: name, IP: ip, Port: port}
	a.devices = append(a.devices, device)
	return a.devices
}

func (a *App) GetDevices() []Device {
	return a.devices
}

func (a *App) DeleteDevice(name string) []Device {
	var newList []Device
	for _, d := range a.devices {
		if d.Name != name {
			newList = append(newList, d)
		}
	}
	a.devices = newList
	return a.devices
}
