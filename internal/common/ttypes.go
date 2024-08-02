package common

import (
	"errors"
	"fmt"
	"os"
	"path"
	"slices"
)

type Template struct {
	Mermaid    string
	BgColor    string
	Theme      string
	IsDarkMode bool
	FontFamily string
	CSSPath    string
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
	DarkMode     bool
	FontFamily   string
}

func (c *Config) ToTemplate() *Template {
	template := &Template{
		FontFamily: "undefined",
		IsDarkMode: c.DarkMode,
	}

	if len(c.FontFamily) > 0 {
		template.FontFamily = EscapeJS(c.FontFamily)
	}

	if len(c.BgColor) > 0 {
		template.BgColor = EscapeJS(c.BgColor)
	}

	if len(c.Theme) > 0 {
		template.Theme = EscapeJS(c.Theme)
	}

	if len(c.CssFile) > 0 {
		template.CSSPath = EscapeJS(c.CssFile)
	}

	return template
}

func (c *Config) Clone() *Config {
	if c == nil {
		return nil
	}

	data := *c
	return &data
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
	darkMode bool,
	fontFamily string,
) (*Config, error) {
	if len(inputPath) == 0 {
		return nil, ErrInvalidInputPath
	}

	if _, err := os.Stat(inputPath); errors.Is(err, os.ErrNotExist) {
		return nil, ErrFileDoesNotExists
	}

	if len(outputPath) == 0 {
		outputPath = path.Join(path.Dir(inputPath), fmt.Sprintf("%s.%s", path.Base(inputPath), outputFormat))
	}

	if !slices.Contains(ValidOutputFormats, outputFormat) {
		return nil, ErrInvalidOutputFormat
	}

	if _, err := os.Stat(cssFile); len(cssFile) > 0 && errors.Is(err, os.ErrNotExist) {
		return nil, ErrFileDoesNotExists
	}

	if !slices.Contains(ValidThemes, theme) {
		return nil, ErrInvalidTheme
	}

	if _, err := os.Stat(cssFile); len(cssFile) > 0 && err != nil {
		return nil, ErrInvalidCSSFilePath
	}

	return &Config{
		theme,
		width,
		height,
		inputPath,
		outputPath,
		outputFormat,
		bgColor,
		configFile,
		cssFile,
		svgID,
		scale,
		pdfFit,
		quietMode,
		darkMode,
		fontFamily,
	}, nil
}
