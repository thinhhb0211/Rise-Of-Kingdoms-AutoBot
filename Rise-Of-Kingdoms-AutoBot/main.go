package main

import (
	"Rise-Of-Kingdoms-AutoBot/pkg/config"
	"embed"
	"fmt"
	"log"

	"github.com/electricbubble/gadb"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	adbClient, err := gadb.NewClientWith(
		"127.0.0.1",
		5037,
	)
	checkErr(err, "fail to connect adb server")

	devices, err := adbClient.DeviceList()
	fmt.Println("devices: ", devices)
	checkErr(err)

	if len(devices) == 0 {
		log.Fatalln("list of devices is empty")
	}

	dev := devices[0]
	fmt.Println("device: ", dev)

	config, err := config.LoadConfig()
	fmt.Println("config: ", config)
	checkErr(err, "fail to load config")
	// Create an instance of the app structure
	app := NewApp(config)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "Rise Of Kingdoms Auto Bot",
		Width:  config.ScreenSize[0],
		Height: config.ScreenSize[1],
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func checkErr(err error, msg ...string) {
	if err == nil {
		return
	}

	var output string
	if len(msg) != 0 {
		output = msg[0] + " "
	}
	output += err.Error()
	log.Fatalln(output)
}
