package day3

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
			given:          bytes.NewBufferString("00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010"),
			expectedResult: 198,
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
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			observedResult, observedErr := Part1(tc.given)

			assert.Equal(t, tc.expectedResult, observedResult)
			assert.True(t, errors.Is(observedErr, tc.expectedErr), observedErr)
		})
	}
}
