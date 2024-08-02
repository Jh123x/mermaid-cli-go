package common

import "fmt"

const (
	TmplDir  = "template"
	TmplFile = "template.html"
	Selector = "body > pre > svg"

	FORMAT_SVG  = "svg"
	FORMAT_MD   = "md"
	FORMAT_PDF  = "pdf"
	FORMAT_PNG  = "png"
	FORMAT_WEBP = "webp"
	FORMAT_JPEG = "jpeg"

	THEME_DEFAULT = "default"
	THEME_FOREST  = "forest"
	THEME_DARK    = "dark"
	THEME_NEUTRAL = "neutral"
	THEME_NULL    = "null"
)

var (
	ErrInvalidInputPath    = fmt.Errorf("input path cannot be empty")
	ErrFileDoesNotExists   = fmt.Errorf("input file does not exists")
	ErrInvalidOutputFormat = fmt.Errorf("output file must end with md/svg/png or pdf")
	ErrNotSupported        = fmt.Errorf("this is not supported yet")
	ErrConfigNotFound      = fmt.Errorf("config is not found")
	ErrInvalidTheme        = fmt.Errorf("invalid theme")
	ErrInvalidCSSFilePath  = fmt.Errorf("invalid CSS path")

	ValidThemes = []string{THEME_DEFAULT, THEME_FOREST, THEME_DARK, THEME_NEUTRAL, THEME_NULL}

	ValidOutputFormats = []string{FORMAT_SVG, FORMAT_MD, FORMAT_PDF, FORMAT_PNG, FORMAT_JPEG, FORMAT_WEBP}
)
