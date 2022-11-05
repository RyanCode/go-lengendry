package main

import (
	"log"
	"os"

	"github.com/go-vgo/robotgo"
	"github.com/sirupsen/logrus"
	"github.com/vcaesar/bitmap"
)

//window contains some basic attributes
type window struct {
	screen robotgo.CBitmap
	width  int
	height int
}

var _window *window = &window{}

var logger *logrus.Logger = logrus.New()

func init() {
	//Get the first screen
	_window.width, _window.height = robotgo.GetScreenSize()
	screen := robotgo.CaptureScreen(0, 0, _window.width, _window.height)
	defer robotgo.FreeBitmap(screen)
	_window.screen = screen

	//Set logger settings
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Panic("cannot open log file,ensure we get log file permission")
	}
	logger.Out = file
	logger.SetLevel(logrus.DebugLevel)
}

func main() {
	fpid, err := robotgo.FindIds("Lengend")
	if err != nil || fpid == nil || len(fpid) < 1 {
		logrus.Panicf("Lengend is not running,please open it first")
	}
	bit := robotgo.CaptureScreen(20, 20, 500, 500)
	defer robotgo.FreeBitmap(bit)

	logger.Println(bitmap.Find(bit, _window.screen))
	logger.Info(fpid)
}
