package common

import (
	"errors"
	"fmt"
	"os"
	"slices"
)

type Template struct {
	Mermaid string
	BgColor string
}

type Config struct {
	Theme        string
	Width        int
	Height       int
	InputPath    string
	OutputPath   string
	OutputFormat string
	BgColor      string
	ConfigFile   string
	CssFile      string
	SvgID        string
	Scale        int
	PDFFit       bool
	QuietMode    bool
}

func NewConfig(
	theme string,
	width int,
	height int,
	inputPath string,
	outputPath string,
	outputFormat string,
	bgColor string,
	configFile string,
	cssFile string,
	svgID string,
	scale int,
	pdfFit bool,
	quietMode bool,
) (*Config, error) {
	if len(inputPath) == 0 {
		return nil, ErrInvalidInputPath
	}

	if _, err := os.Stat(inputPath); errors.Is(err, os.ErrNotExist) {
		return nil, ErrFileDoesNotExists
	}

	if len(outputPath) == 0 {
		outputPath = fmt.Sprintf("%s.%s", inputPath, outputFormat)
	}

	if !slices.Contains([]string{"svg", "md", "png", "pdf"}, outputFormat) {
		return nil, ErrInvalidOutputFormat
	}

	if _, err := os.Stat(cssFile); len(cssFile) > 0 && errors.Is(err, os.ErrNotExist) {
		return nil, ErrFileDoesNotExists
	}

	// TODO Config file

	return &Config{theme, width, height, inputPath, outputPath, outputFormat, bgColor, configFile, cssFile, svgID, scale, pdfFit, quietMode}, nil
}
