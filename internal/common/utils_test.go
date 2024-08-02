package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEscapeJS(t *testing.T) {
	tests := map[string]struct {
		val         string
		expectedVal string
	}{
		"nothing to escape": {
			val:         "test value",
			expectedVal: "test value",
		},
		"single quotes should be escaped": {
			val:         `console.log('Hello, Snippets!');`,
			expectedVal: `console.log(\'Hello, Snippets!\');`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedVal, EscapeJS(tc.val))
		})
	}
}

func TestGetPtrOf(t *testing.T) {
	tests := map[string]string{
		"test value 1": "test_1",
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res := GetPtrOf(tc)
			assert.Equal(t, *res, tc)
		})
	}
}
