package common

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testFilePath = "./test/data_file"
)

func TestNewConfig(t *testing.T) {
	tests := map[string]struct {
		theme        string
		width        int
		height       int
		inputPath    string
		outputPath   string
		outputFormat string
		bgColor      string
		configFile   string
		cssFile      string
		scale        int
		pdfFit       bool
		quietMode    bool
		darkMode     bool
		fontFamily   string

		expectedErr    error
		expectedConfig *Config
	}{
		"empty input path": {
			inputPath:    "",
			outputFormat: FORMAT_JPEG,
			theme:        THEME_DEFAULT,
			expectedErr:  ErrInvalidInputPath,
		},
		"input path not found should error": {
			inputPath:   "not found path",
			expectedErr: ErrFileDoesNotExists,
		},
		"output path should return correct default path": {
			inputPath:    testFilePath,
			outputFormat: FORMAT_JPEG,
			outputPath:   "",
			theme:        THEME_DEFAULT,
			expectedConfig: &Config{
				InputPath:    testFilePath,
				OutputPath:   path.Join(path.Dir(testFilePath), path.Base(testFilePath)+"."+FORMAT_JPEG),
				OutputFormat: FORMAT_JPEG,
				Theme:        THEME_DEFAULT,
			},
		},
		"css path not found": {
			inputPath:    testFilePath,
			cssFile:      "not found path",
			outputFormat: FORMAT_JPEG,
			outputPath:   "test",
			theme:        THEME_DEFAULT,
			expectedErr:  ErrFileDoesNotExists,
		},
		"css path found": {
			inputPath:    testFilePath,
			cssFile:      testFilePath,
			outputFormat: FORMAT_JPEG,
			outputPath:   "test",
			theme:        THEME_DEFAULT,
			expectedConfig: &Config{
				InputPath:    testFilePath,
				CssFile:      testFilePath,
				OutputPath:   "test",
				OutputFormat: FORMAT_JPEG,
				Theme:        THEME_DEFAULT,
			},
		},
		"success": {
			inputPath:    testFilePath,
			outputFormat: FORMAT_JPEG,
			outputPath:   "test",
			theme:        THEME_DEFAULT,
			expectedConfig: &Config{
				InputPath:    testFilePath,
				OutputPath:   "test",
				OutputFormat: FORMAT_JPEG,
				Theme:        THEME_DEFAULT,
			},
		},
		"invalid theme should error": {
			inputPath:    testFilePath,
			outputFormat: FORMAT_JPEG,
			outputPath:   "test",
			theme:        "invalid theme",
			expectedErr:  ErrInvalidTheme,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := NewConfig(tc.theme, tc.width, tc.height, tc.inputPath, tc.outputPath, tc.outputFormat, tc.bgColor, tc.configFile, tc.cssFile, tc.scale, tc.pdfFit, tc.quietMode, tc.darkMode, tc.fontFamily)
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expectedConfig, res)
		})
	}
}
