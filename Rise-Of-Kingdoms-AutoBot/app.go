package main

import (
	"Rise-Of-Kingdoms-AutoBot/pkg/config"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/electricbubble/gadb"
)

// App struct
type App struct {
	ctx    context.Context
	device gadb.Device
	config *config.Configuration
}

// NewApp creates a new App application struct
func NewApp(config *config.Configuration, device gadb.Device) *App {
	return &App{
		config: config,
		device: device,
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

func (a *App) Screenshot() string {
	// print screen shot
	screenImg, err := a.device.RunShellCommand("screencap -p")
	if err != nil {
		log.Fatalln("fail to take screen shot")
	}
	// Save to a file
	file, err := os.Create("screenshot.png")
	if err != nil {
		log.Fatalln("fail to create file")
	}
	defer file.Close()

	_, err = file.Write([]byte(screenImg))
	if err != nil {
		log.Fatalln("fail to write file")
	}

	data, err := os.ReadFile("screenshot.png")
	if err != nil {
		return ""
	}
	encoded := base64.StdEncoding.EncodeToString(data)
	return "data:image/png;base64," + encoded
}

type BuildingPositions map[string][2]int

func SavePositionsToFile(positions BuildingPositions, filepath string) error {
	data, err := json.MarshalIndent(positions, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath, data, 0644)
}

func (a *App) SaveBuildingPositionsJSON(positions map[string][2]int) error {
	return SavePositionsToFile(positions, "positions.json")
}

func (a *App) TapBuilding(x, y int) error {
	// Tap on the building
	cmd := fmt.Sprintf("input tap %d %d", x, y)
	_, err := a.device.RunShellCommand(cmd)
	if err != nil {
		log.Fatalln("fail to tap on building")
	}
	return nil
}
