package day1

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
			given:          bytes.NewBufferString("607\n618\n618\n617\n647\n716\n769\n792"),
			expectedResult: 5,
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
			expectedErr:    errInputInvalid,
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
			observedResult, observedErr := Part2(tc.given)

			assert.Equal(t, tc.expectedResult, observedResult)
			assert.True(t, errors.Is(observedErr, tc.expectedErr))
		})
	}
}
