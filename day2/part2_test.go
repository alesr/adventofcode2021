package day2

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2(t *testing.T) {
	cases := []struct {
		name           string
		given          io.Reader
		expectedResult int
		expectedErr    error
	}{
		{
			name:           "with site example input",
			given:          bytes.NewBufferString("forward 5\ndown 5\nforward 8\nup 3\ndown 8\nforward 2"),
			expectedResult: 900,
			expectedErr:    nil,
		},
		{
			name:           "with nil input",
			given:          nil,
			expectedResult: 0,
			expectedErr:    errInputNil,
		},
		{
			name:           "with empty input",
			given:          bytes.NewBufferString(""),
			expectedResult: 0,
			expectedErr:    nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			observedResult, observedErr := Part2(tc.given)

			assert.Equal(t, tc.expectedResult, observedResult)
			assert.True(t, errors.Is(observedErr, tc.expectedErr))
		})
	}
}
