package handler

import (
	"io/fs"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getInputData(t *testing.T) {
	tests := map[string]struct {
		inputFile   string
		expectedRes string
		expectedErr error
	}{
		"success": {
			inputFile:   "../common/test/data_file",
			expectedRes: "test_data",
		},
		"file does not exists": {
			inputFile:   "does not exists",
			expectedErr: &fs.PathError{Op: "open", Path: "does not exists", Err: syscall.Errno(2)},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			data, err := getInputData(tc.inputFile)
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expectedRes, data)
		})
	}
}
