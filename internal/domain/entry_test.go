package domain_test

import (
	"testing"

	"github.com/firamisu/louis/internal/dictclient"
	"github.com/firamisu/louis/internal/domain"
	"github.com/go-playground/assert/v2"
)

func TestFromDictResponse(t *testing.T) {
	testCases := []struct {
		name     string
		input    dictclient.DictResponse
		expected domain.Entry
	}{
		{
			name: "single meaning and definition",
			input: dictclient.DictResponse{
				Word:     "hello",
				Phonetic: "həˈləʊ",
				Origin:   "Old English",
				Meanings: []dictclient.Meaning{
					{
						PartOfSpeech: "interjection",
						Definitions: []dictclient.Definition{
							{
								Definition: "used as a greeting",
								Example:    "Hello!",
							},
						},
					},
				},
			},
			expected: domain.Entry{
				Word:     "hello",
				Phonetic: "həˈləʊ",
				Origin:   "Old English",
				Meanings: []domain.Meaning{
					{
						PartOfSpeech: "interjection",
						Definitions: []domain.Definition{
							{
								Definition: "used as a greeting",
								Example:    "Hello!",
							},
						},
					},
				},
			},
		},
		{
			name: "multiple meanings and definitions",
			input: dictclient.DictResponse{
				Word: "run",
				Meanings: []dictclient.Meaning{
					{
						PartOfSpeech: "verb",
						Definitions: []dictclient.Definition{
							{Definition: "move fast by using one’s feet", Example: "I run every morning"},
							{Definition: "manage or operate", Example: "run a company"},
						},
					},
					{
						PartOfSpeech: "noun",
						Definitions: []dictclient.Definition{
							{Definition: "an act or spell of running", Example: "a five-mile run"},
						},
					},
				},
			},
			expected: domain.Entry{
				Word: "run",
				Meanings: []domain.Meaning{
					{
						PartOfSpeech: "verb",
						Definitions: []domain.Definition{
							{Definition: "move fast by using one’s feet", Example: "I run every morning"},
							{Definition: "manage or operate", Example: "run a company"},
						},
					},
					{
						PartOfSpeech: "noun",
						Definitions: []domain.Definition{
							{Definition: "an act or spell of running", Example: "a five-mile run"},
						},
					},
				},
			},
		},
		{
			name: "empty meanings",
			input: dictclient.DictResponse{
				Word: "empty",
			},
			expected: domain.Entry{
				Word:     "empty",
				Phonetic: "",
				Origin:   "",
			},
		},
		{
			name: "meaning with no definitions",
			input: dictclient.DictResponse{
				Word: "nodef",
				Meanings: []dictclient.Meaning{
					{
						PartOfSpeech: "adjective",
						Definitions:  []dictclient.Definition{},
					},
				},
			},
			expected: domain.Entry{
				Word: "nodef",
				Meanings: []domain.Meaning{
					{
						PartOfSpeech: "adjective",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := domain.FromDictResponse(tc.input)
			assert.Equal(t, tc.expected, out)
		})
	}
}
