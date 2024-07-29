package common

import "fmt"

const (
	TmplDir  = "template"
	TmplFile = "template.html"
	Selector = "body > pre > svg"
)

var (
	ErrInvalidInputPath    = fmt.Errorf("input path cannot be empty")
	ErrFileDoesNotExists   = fmt.Errorf("input file does not exists")
	ErrInvalidOutputFormat = fmt.Errorf("output file must end with md/svg/png or pdf")
	ErrConfigNotFound      = fmt.Errorf("config is not found")
)
