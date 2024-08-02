package main

import (
	"flag"
	"fmt"

	"github.com/jh123x/mermaid-cli-go/internal/common"
	"github.com/jh123x/mermaid-cli-go/internal/handler"
)

func main() {
	// Currently Supported flags
	input := flag.String("i", "", "Input mermaid file (Only raw mermaid supported currently)")
	output := flag.String("o", "", "Output file. It should be svg, default: input + \".svg\"")
	bgColor := flag.String("b", "white", "Background Color")
	quietMode := flag.Bool("q", false, "Suppress log output")
	theme := flag.String("t", "default", "Theme of the chart")
	cssFile := flag.String("C", "", "CSS for the page")
	outputFormat := flag.String("e", "svg", "Output format for the generated image")

	// Other flags not in mermaid cli
	darkMode := flag.Bool("d", false, "Enable dark mode")
	fontFamily := flag.String("ff", "", "Font Family")

	// TODO Not supported yet
	width := flag.Int("w", 800, "Width of the page")
	height := flag.Int("H", 600, "Height of the page")
	configFile := flag.String("c", "", "Config File")
	svgID := flag.String("I", "", "The ID attribute for the svg element to be rendered")
	scale := flag.Int("s", 1, "Scale factor")
	pdfFit := flag.Bool("f", false, "Scale PDF to fit chart")

	flag.Parse()

	config, err := common.NewConfig(
		*theme,
		*width,
		*height,
		*input,
		*output,
		*outputFormat,
		*bgColor,
		*configFile,
		*cssFile,
		*svgID,
		*scale,
		*pdfFit,
		*quietMode,
		*darkMode,
		*fontFamily,
	)
	if err != nil {
		panic(err)
	}

	path, err := handler.GetSVG(config)
	if err != nil {
		panic(err)
	}

	if !config.QuietMode {
		fmt.Printf("Successfully created %s\n", path)
	}
}
