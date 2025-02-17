package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fogleman/primitive/primitive"
)

var (
	Input      string
	Output     string
	Number     int
	Alpha      int
	Scale      int
	Mode       int
	V, VV, VVV bool
)

func init() {
	flag.StringVar(&Input, "i", "", "input image path")
	flag.StringVar(&Output, "o", "", "output image path")
	flag.IntVar(&Number, "n", 0, "number of primitives")
	flag.IntVar(&Alpha, "a", 128, "alpha value")
	flag.IntVar(&Scale, "s", 1, "output image scale")
	flag.IntVar(&Mode, "m", 1, "mode: 0=combo, 1=triangle, 2=rectangle, 3=ellipse, 4=circle")
	flag.BoolVar(&V, "v", false, "verbose")
	// flag.BoolVar(&VV, "vv", false, "very verbose")
	// flag.BoolVar(&VVV, "vvv", false, "very very verbose")
}

func errorMessage(message string) bool {
	fmt.Fprintln(os.Stderr, message)
	return false
}

func main() {
	flag.Parse()
	ok := true
	if Input == "" {
		ok = errorMessage("input argument required")
	}
	if Output == "" {
		ok = errorMessage("output argument required")
	}
	if Number == 0 {
		ok = errorMessage("number argument required")
	}
	if !ok {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if V {
		primitive.LogLevel = 1
	}
	if VV {
		primitive.LogLevel = 2
	}
	if VVV {
		primitive.LogLevel = 3
	}

	rand.Seed(time.Now().UTC().UnixNano())
	primitive.Log(1, "reading %s\n", Input)
	input, err := primitive.LoadImage(Input)
	if err != nil {
		panic(err)
	}
	mode := primitive.Mode(Mode)
	model := primitive.NewModel(input, Alpha, Scale, mode)
	output := model.Run(Number)
	primitive.Log(1, "writing %s\n", Output)
	if strings.HasSuffix(strings.ToLower(Output), ".gif") {
		frames := model.Frames(0.001)
		primitive.SaveGIFImageMagick(Output, frames, 50, 250)
	} else {
		primitive.SavePNG(Output, output)
	}
}
