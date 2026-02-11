package subtitles_test

import (
	"context"
	"testing"

	"github.com/firamisu/louis/internal/subtitles"
	"github.com/go-playground/assert/v2"
)

func TestList(t *testing.T) {
	testCases := []struct {
		name     string
		errRes   error
		sliceRes []subtitles.Subtitle
	}{
		{
			name:     "empty list",
			errRes:   nil,
			sliceRes: []subtitles.Subtitle{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			svc := subtitles.NewService()

			s, err := svc.List(context.Background())

			assert.Equal(t, testCase.errRes, err)
			assert.Equal(t, testCase.sliceRes, s)
		})
	}
}
