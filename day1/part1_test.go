package day1

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	cases := []struct {
		name           string
		given          io.Reader
		expectedResult int
		expectedErr    error
	}{
		{
			name:           "with site example input",
			given:          bytes.NewBufferString("199\n200\n208\n210\n200\n207\n240\n269\n260\n263"),
			expectedResult: 7,
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
		{
			name:           "with non-numeric input",
			given:          bytes.NewBufferString("199\n200\n208\n210\n200\n207\n240\n269\n260\n263\nabc"),
			expectedResult: 0,
			expectedErr:    errInputInvalid,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			observedResult, observedErr := Part1(tc.given)

			assert.Equal(t, tc.expectedResult, observedResult)
			assert.True(t, errors.Is(observedErr, tc.expectedErr))
		})
	}
}
