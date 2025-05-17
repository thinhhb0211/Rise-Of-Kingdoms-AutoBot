package main

import (
	"embed"
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"github.com/electricbubble/gadb"
	"gocv.io/x/gocv"
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
	checkErr(err)

	if len(devices) == 0 {
		log.Fatalln("list of devices is empty")
	}

	dev := devices[0]

	// config, err := config.LoadConfig()

	// take screenshot
	screenImg, err := dev.RunShellCommand("screencap -p")
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

	checkErr(err, "fail to load config")

	// color for the rect
	red := color.RGBA{255, 0, 0, 1}

	img := gocv.IMRead("screenshot.png", gocv.IMReadAnyColor)
	template := gocv.IMRead("resource/upgrade_button.png", gocv.IMReadAnyColor)

	matResult := gocv.NewMat()
	mask := gocv.NewMat()

	grayImg := gocv.NewMat()
	grayTemplate := gocv.NewMat()

	gocv.CvtColor(img, &grayImg, gocv.ColorBGRToGray)
	gocv.CvtColor(template, &grayTemplate, gocv.ColorBGRToGray)
	// TmCcoeffNormed - works
	// TmCcorrNormed - works
	// TmSqdiffNormed - doesn't work
	gocv.MatchTemplate(grayImg, grayTemplate, &matResult, gocv.TmCcorrNormed, mask)
	defer mask.Close()

	minConfidence, maxConfidence, minLoc, maxLoc := gocv.MinMaxLoc(matResult)
	defer matResult.Close()
	fmt.Println("minConfidence:", minConfidence)
	fmt.Println("maxConfidence:", maxConfidence)
	fmt.Println("minLoc:", minLoc)
	fmt.Println("maxLoc:", maxLoc)

	dims := template.Size()
	// dims[1] = width
	// dims[0] = height

	// can also use Rows & Cols for (x,y)
	// template.Cols() is width
	// template.Rows() is height

	//fmt.Println("Template Size: ", dims[1], dims[0]) // width & height
	//fmt.Println("Template Col/Rows: ", template.Cols(), template.Rows())

	point := image.Point{maxLoc.X + dims[1], maxLoc.Y + dims[0]}
	// if maxConfidence > 0.8 {
	gocv.Rectangle(&img, image.Rectangle{Min: maxLoc, Max: point}, red, 1)
	// }
	gocv.IMWrite("OutputImage.jpg", img)

	// cmd := fmt.Sprintf("input tap %d %d", maxLoc.X+dims[1], maxLoc.Y+dims[0])

	// Create an instance of the app structure
	// app := NewApp(config, dev)

	// // Create application with options
	// err = wails.Run(&options.App{
	// 	Title:  "Rise Of Kingdoms Auto Bot",
	// 	Width:  config.ScreenSize[0],
	// 	Height: config.ScreenSize[1],
	// 	AssetServer: &assetserver.Options{
	// 		Assets: assets,
	// 	},
	// 	BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
	// 	OnStartup:        app.startup,
	// 	Bind: []interface{}{
	// 		app,
	// 	},
	// })

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
